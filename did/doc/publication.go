package doc

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/ipfs/boxo/path"
	caopts "github.com/ipfs/kubo/core/coreiface/options"
)

// Publish publishes the document to IPFS and returns the CID
func (d *Document) Publish() (string, error) {
	data, err := d.MarshalToCBOR()
	if err != nil {
		return "", err
	}

	cid, err := internal.IPFSDagPutWithOptions(data, "dag-cbor", "dag-cbor", true, "sha2-256", false)
	if err != nil {
		return "", fmt.Errorf("doc/publish: failed to put document to IPFS: %w", err)
	}

	ipfsAPI := internal.GetIPFSAPI()
	ctx := internal.GetContext()
	p, err := path.NewPathFromSegments(path.IPFSNamespace, cid)
	if err != nil {
		return "", err
	}

	ipfsAPI.Name().Publish(ctx, p, caopts.Name.Key(internal.GetDIDFragment(d.ID)))

	return cid, nil
}
