package internal

import (
	"fmt"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/core/coreiface/options"
)

// Takes an IPFS path name and returns the root CID.
// The cached field tells the function whether to use the cached value or not.
func RootCID(name string, cached bool) (cid.Cid, error) {

	api := GetIPFSAPI()

	opts := func(settings *options.NameResolveSettings) error {
		settings.Cache = cached
		return nil
	}

	p, err := api.Name().Resolve(GetContext(), name, opts)
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
