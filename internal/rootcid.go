package internal

import (
	"fmt"

	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
)

// Takes an IPFS path name and returns the root CID.
func RootCID(name string) (cid.Cid, error) {

	api := GetIPFSAPI()

	p, err := api.Name().Resolve(GetContext(), name)
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
