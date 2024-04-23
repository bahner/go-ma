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

// Publishes the document to IPFS and IPNS. This is a little noise, as it does a lot of things.
// Notably it can take a long time, which is why there's also a gorutine version.
func (d *Document) Publish() (ipns.Name, error) {

	ctx := context.Background()

	node, err := d.Node()
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}
	newImmutablePath := path.FromCid(node.Cid)

	// Get the IPFS API
	a := api.GetIPFSAPI()

	// Add the Document Node to the IPFS DAG
	err = a.Dag().Add(ctx, node.Node)
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}
	log.Infof("DocPublish: Successfully added document %s to IPLD", node.Cid.String())

	// Pin the document
	log.Infof("Pinning %s in IPFS. Please wait ...", newImmutablePath.String())
	err = pinDocument(ctx, d.immutablePath, newImmutablePath)
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}

	// Now that the document is pinned we can update the path to the new one.
	d.immutablePath = newImmutablePath

	log.Debugf("DocPublish: Announcing publication of document %s to IPNS. Please wait ...", node.Cid.String())
	name, err := a.Name().Publish(ctx, newImmutablePath, options.Name.Key(d.did.IPNS))
	if err != nil {
		return ipns.Name{}, fmt.Errorf("DocPublish: %w", err)
	}
	log.Debugf("DocPublish: Successfully announced publication of document %s to %s", node.Cid.String(), name.AsPath().String())
	return name, nil

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

func pinDocument(ctx context.Context, oldP path.ImmutablePath, newP path.ImmutablePath) error {

	alreadyPublishedString := "'to' cid was already recursively pinned"

	a := api.GetIPFSAPI()

	var err error
	if (oldP == path.ImmutablePath{}) {
		err = a.Pin().Add(ctx, newP)
		if err != nil {
			return fmt.Errorf("DocPublish: %w", err)
		}
	} else {
		// If an existing document is already published and Pin is set we need to update the existing pin f asked to force.
		err = a.Pin().Update(ctx, oldP, newP)
		if err != nil {
			if err.Error() == alreadyPublishedString {
				return fmt.Errorf("DocPublish: %w", ErrAlreadyPublished)
			}
			return fmt.Errorf("DocPublish: %w", err)
		}
	}

	return nil

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
