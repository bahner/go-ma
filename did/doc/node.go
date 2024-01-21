package doc

import (
	"fmt"

	cbor "github.com/ipfs/go-ipld-cbor"
	mh "github.com/multiformats/go-multihash"
)

// IPFSDagAddCBOR takes a CBOR encoded byte array and adds it to IPFS.
func (d *Document) Node() (*cbor.Node, error) {

	dCBOR, err := d.MarshalToCBOR()
	if err != nil {
		return nil, fmt.Errorf("doc/Node: failed to marshal document to CBOR: %w", err)
	}

	n, err := cbor.WrapObject(dCBOR, mh.SHA2_256, -1)
	if err != nil {
		return nil, fmt.Errorf("doc/Node: failed to wrap object: %w", err)
	}

	return n, nil
}
