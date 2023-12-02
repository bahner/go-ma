package internal

import (
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-merkledag"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
	log "github.com/sirupsen/logrus"
)

// Takes a CBOR encoded byte array and adds it to IPFS
func IPFSDagAddCBOR(data []byte) (string, error) {
	// Create a new ProtoNode with the provided data
	nd := merkledag.NodeWithData(data)

	// Set CID builder for the node
	pref := cid.Prefix{
		Version:  1,
		Codec:    uint64(mc.DagCbor),
		MhType:   uint64(mh.SHA2_256),
		MhLength: -1, // Default length
	}
	nd.SetCidBuilder(pref)

	// Add the node to IPFS
	err := GetIPSAPI().Dag().Add(GetContext(), nd)
	if err != nil {
		log.Printf("ipld: failed to add node to IPFS: %v", err)
		return "", err
	}

	// Return the CID string of the added node
	return nd.Cid().String(), nil
}
