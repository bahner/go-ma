package doc

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/bahner/go-ma/api"
	"github.com/ipfs/boxo/ipns"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/core/coreiface/options"
	log "github.com/sirupsen/logrus"
)

var ErrAlreadyPublished = errors.New("Document is already published")

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
		Force:         true,
		AllowBigBlock: false,
	}
}

// Publishes document to a key known by the IPFS node. This maybe a peer ID or a name.
// Both provided as a simple string.
// The only option we honour is the force option. If set to true we will update the existing pin regardless.
func (d *Document) Publish() (ipns.Name, error) {

	ctx := context.Background()
	alreadyPublishedString := "'to' cid was already recursively pinned"

	node, newCID, err := d.Node()
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}
	newImmutablePath := path.FromCid(newCID)

	// Get the IPFS API
	a := api.GetIPFSAPI()
	a.Dag().Add(ctx, node)

	// If an existing document is already published and Pin is set we need to update the existing pin f asked to force.
	err = a.Pin().Update(ctx, d.immutablePath, newImmutablePath)
	if err != nil {
		if err.Error() == alreadyPublishedString {
			return ipns.Name{}, fmt.Errorf("DocPublish: %w", ErrAlreadyPublished)
		}
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}

	// Now that the document is pinned we can update the path to the new one.
	d.immutablePath = newImmutablePath

	log.Debugf("DocPublish: Announcing publication of document %s to IPNS. Please wait ...", newCID.String())
	n, err := a.Name().Publish(ctx, newImmutablePath, options.Name.Key(d.did.Name.String()))
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}
	log.Debugf("DocPublish: Successfully announced publication of document %s to %s", newCID.String(), n.AsPath().String())
	return n, nil

}

// The Goroutine version of Publish must get a cancel function as argument.
// This is to force the caller to use a context with a cancel function.
// Obviously this should probably be a timeout context.
// Other than that it is the same as Publish.
func (d *Document) PublishGoroutine(wg *sync.WaitGroup, cancel context.CancelFunc) {

	defer cancel()
	defer wg.Done()

	d.Publish()

}

// Takes an IPFS path name and returns the root CID.
// The cached field tells the function whether to use the cached value or not.
func resolveRootCID(name string) (cid.Cid, error) {

	api := api.GetIPFSAPI()

	// Set cached to false, we need to find the latest version of the document
	opts := func(settings *options.NameResolveSettings) error {
		settings.Cache = false
		return nil
	}

	p, err := api.Name().Resolve(context.Background(), name, opts)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("doc/fetch: failed to decode cid: %w", err)
	}

	// Create an immutable path from the resolved path
	// NB! The resolved path must be immutable.
	ip, err := path.NewImmutablePath(p)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("failed to create RootCID: %w", err)
	}

	return ip.RootCid(), nil

}
