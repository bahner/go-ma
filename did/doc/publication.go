package doc

import (
	"context"
	"fmt"
	"sync"

	"github.com/bahner/go-ma/api"
	"github.com/bahner/go-ma/did"
	ipfsKey "github.com/bahner/go-ma/key/ipfs"
	"github.com/ipfs/boxo/ipns"
	"github.com/ipfs/boxo/path"
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
		Pin:           false,
		Force:         false,
		AllowBigBlock: false,
	}
}

// Publish publishes the document to IPFS and returns the CID
// If the opts is nil, the default options are used.
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

	if d.isPublished() {
		if opts.Force {
			log.Infof("DocPublish: Document %s is already published. Forcing publication.", d.ID)
		} else {
			log.Warnf("DocPublish: Document %s is already published and not forced. Aborting ....", d.ID)
			return ipns.Name{}, ErrDoumentAlreadyPublished
		}
	}

	_did, err := did.New(d.ID)
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}

	// Make sure a key is available for the document
	ik, err := ipfsKey.GetOrCreate(_did.Fragment)
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

	// Creates an immutable path from the CID
	p := path.FromCid(c)

	log.Debugf("DocPublish: Announcing publication of document %s to IPNS. Please wait ...", c.String())
	ipfsAPI := api.GetIPFSAPI()
	n, err := ipfsAPI.Name().Publish(opts.Ctx, p, caopts.Name.Key(ik.Fragment))
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

func (d *Document) isPublished() bool {

	maybeDoc, err := Fetch(d.ID, false) // Don't accept cached document
	if err != nil {
		log.Debugf("DocumentPublishGoroutine: %v", err)
		return false
	}

	if maybeDoc == nil {
		log.Debugf("DocumentPublishGoroutine: Document is Nill")
		return false
	}

	if !d.Equal(maybeDoc) {
		log.Debugf("DocumentPublishGoroutine: %s:", d.ID)
		return true
	}
	return false
}
