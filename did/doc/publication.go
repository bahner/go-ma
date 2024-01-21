package doc

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/ipfs/boxo/ipns"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	caopts "github.com/ipfs/kubo/core/coreiface/options"
)

// Publish publishes the document to IPFS and returns the CID
func (d *Document) Publish(pin bool) (ipns.Name, error) {

	ipfsAPI := internal.GetIPFSAPI()
	ctx := internal.GetContext()

	data, err := d.MarshalToCBOR()
	if err != nil {
		return ipns.Name{}, fmt.Errorf("doc/publish: failed to marshal document to CBOR: %w", err)
	}

	// Actually add the document to IPFS and possibly pins it.
	var c cid.Cid
	if pin {
		c, err = internal.IPFSDagPutCBORAndPin(data)
	} else {
		c, err = internal.IPFSDagPutCBOR(data)
	}
	if err != nil {
		return ipns.Name{}, fmt.Errorf("doc/publish: failed to add document to IPFS: %w", err)
	}

	// Creates an immutable path from the CID
	p := path.FromCid(c)

	return ipfsAPI.Name().Publish(ctx, p, caopts.Name.Key(internal.GetDIDFragment(d.ID)))

}
