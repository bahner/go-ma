package doc

import (
	"context"
	"fmt"
	"sync"

	"github.com/bahner/go-ma/api"
	"github.com/bahner/go-ma/did"
	"github.com/ipfs/boxo/ipns"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	caopts "github.com/ipfs/kubo/core/coreiface/options"
	log "github.com/sirupsen/logrus"
)

type PublishOptions struct {
	Ctx           context.Context
	Pin           bool
	Force         bool
	AllowBigBlock bool
}

func DefaultPublishOptions() *PublishOptions {
	return &PublishOptions{
		Ctx:           context.Background(),
		Pin:           true,
		Force:         false,
		AllowBigBlock: false,
	}
}

// Publish publishes the document to IPFS and returns the CID
// If the opts is nil, the default options are used.
// NB! Publication is more than simply adding the document to IPFS.
// It's about publishing the document to IPNS and possibly pinning it.
func (d *Document) Publish(opts *PublishOptions) (ipns.Name, error) {

	if opts == nil {
		opts = DefaultPublishOptions()
	}

	if opts.Force {
		log.Debugf("DocPublish: force flag is set")
	}

	if opts.Pin {
		log.Debugf("DocPublish: pin flag is set")
	}

	if opts.AllowBigBlock {
		log.Debugf("DocPublish: allow big block flag is set")
	}

	_did, err := did.New(d.ID)
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}

	// Make sure a key is available for the document
	ik, err := did.GetOrCreate(_did.Fragment)
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}

	data, err := d.MarshalToCBOR()
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}

	// Actually add the document to IPFS and possibly pin it and allow bib blocks.
	c, err := api.IPFSDagPutCBOR(data, opts.Pin, opts.AllowBigBlock)
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}
	p := path.FromCid(c)

	// Get the IPFS API
	a := api.GetIPFSAPI()

	// If an existing document is already published and Pin is set we need to update the existing pin f asked to force.
	e, err := d.CID()
	if err == nil && c != e && opts.Pin && opts.Force {
		err := a.Pin().Update(opts.Ctx, path.FromCid(e), p)
		if err != nil {
			// Not sure if this is "fatal" or not. We should probably return the error and let the caller decide.
			return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
		}
	} else {
		log.Debugf("DocPublish: Document %s is not yet published.", d.ID)
	}

	log.Debugf("DocPublish: Announcing publication of document %s to IPNS. Please wait ...", c.String())
	n, err := a.Name().Publish(opts.Ctx, p, caopts.Name.Key(ik.Fragment))
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}
	log.Debugf("DocPublish: Successfully announced publication of document %s to %s", c.String(), n.AsPath().String())
	return n, nil

}

// The Goroutine version of Publish must get a cancel function as argument.
// This is to force the caller to use a context with a cancel function.
// Obviously this should probably be a timeout context.
// Other than that it is the same as Publish.
func (d *Document) PublishGoroutine(wg *sync.WaitGroup, cancel context.CancelFunc, opts *PublishOptions) {

	defer cancel()
	defer wg.Done()

	d.Publish(opts)

}

func (d *Document) CID() (cid.Cid, error) {

	maybeDoc, c, err := Fetch(d.ID, false) // Don't accept cached document
	if err != nil {
		log.Debugf("DocumentPublishGoroutine: %v", err)
		return cid.Cid{}, err
	}

	if maybeDoc == nil {
		log.Debugf("DocumentPublishGoroutine: Document is nil")
		return cid.Cid{}, fmt.Errorf("Document is nil")
	}

	return c, nil
}
