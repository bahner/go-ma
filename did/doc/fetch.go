package doc

import (
	"context"
	"fmt"

	"github.com/bahner/go-ma/api"
	"github.com/bahner/go-ma/did"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
)

// Takes a DID and fetches the document from IPFS.
// Eg. Fetch("did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner")
// The cached flag determines whether the document should be fetched from IPNS cache or not.
func Fetch(didStr string) (*Document, cid.Cid, error) {

	d, err := did.NewFromString(didStr)
	if err != nil {
		return nil, cid.Cid{}, err
	}

	return FetchFromDID(d)

}

// Fetched the document from IPFS using the DID object.
// The cached flag determines whether the document should be fetched from IPNS cache or not.
func FetchFromDID(d did.DID) (*Document, cid.Cid, error) {

	var document = &Document{}

	// Resolve the root CID, ie. the actual IPLD CID of the document
	c, err := resolveRootCID(d.IPNS)
	if err != nil {
		return nil, cid.Cid{}, fmt.Errorf("fetchFromDID: %w", err)
	}

	log.Debugf("Fetching CID: %s", c)
	ipfsAPI := api.GetIPFSAPI()
	node, err := ipfsAPI.Dag().Get(context.Background(), c)
	if err != nil {
		return nil, cid.Cid{}, fmt.Errorf("fetchFromDID: %w", err)
	}

	err = cbor.Unmarshal(node.RawData(), document)
	if err != nil {
		return nil, cid.Cid{}, fmt.Errorf("fetchFromDID: %w", err)
	}

	err = document.Verify()
	if err != nil {
		return nil, cid.Cid{}, fmt.Errorf("fetchFromDID: %w", err)
	}

	log.Debugf("Fetched and cached document for : %s", d.Id)
	return document, c, nil

}
