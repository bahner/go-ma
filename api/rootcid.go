package api

import (
	"context"
	"fmt"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/core/coreiface/options"
)

// Takes an IPFS path name and returns the root CID.
// The cached field tells the function whether to use the cached value or not.
func ResolveRootCID(name string) (cid.Cid, error) {

	api := GetIPFSAPI()

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
