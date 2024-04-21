package doc

import (
	"github.com/ipfs/go-cid"
	ipldcbor "github.com/ipfs/go-ipld-cbor"
	format "github.com/ipfs/go-ipld-format"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
)

func (d *Document) Node() (format.Node, cid.Cid, error) {

	// Hash the CBOR data to create a CID
	data, err := d.MarshalToCBOR()
	if err != nil {
		return nil, cid.Cid{}, err
	}

	hash, err := mh.Sum(data, mh.SHA2_256, -1)
	if err != nil {
		return nil, cid.Cid{}, err
	}

	// Create a CID with the DagCBOR codec
	codecType := uint64(mc.DagCbor)
	c := cid.NewCidV1(codecType, hash)

	// Use go-ipld-cbor to decode the CBOR data into an IPLD node
	node, err := ipldcbor.Decode(data, mh.SHA2_256, -1)
	if err != nil {
		return nil, cid.Cid{}, err
	}

	return node, c, nil
}
