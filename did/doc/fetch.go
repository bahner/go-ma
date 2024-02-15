package doc

import (
	"context"
	"fmt"

	"github.com/bahner/go-ma/api"
	"github.com/bahner/go-ma/did"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/ipfs/boxo/path"
	log "github.com/sirupsen/logrus"
)

// Takes a DID and fetches the document from IPFS.
// Eg. Fetch("did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner")
// The cached flag determines whether the document should be fetched from IPNS cache or not.
func Fetch(didStr string, cached bool) (*Document, error) {

	d, err := did.New(didStr)
	if err != nil {
		return nil, err
	}

	return FetchFromDID(d, cached)

}

// Fetched the document from IPFS using the DID object.
// The cached flag determines whether the document should be fetched from IPNS cache or not.
func FetchFromDID(d did.DID, cached bool) (*Document, error) {

	var document = &Document{}

	ipfsAPI := api.GetIPFSAPI()

	_cid, err := api.RootCID(d.Path(path.IPNSNamespace), cached)
	if err != nil {
		return nil, fmt.Errorf("failed to get CID from IPNS name: %w", err)
	}
	log.Debugf("Fetching CID: %s", _cid)

	node, err := ipfsAPI.Dag().Get(context.Background(), _cid)
	if err != nil {
		return nil, fmt.Errorf("failed to get document from IPFS: %w", err)
	}

	err = cbor.Unmarshal(node.RawData(), document)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal document: %w", err)
	}

	if !document.isValid() {
		return nil, ErrDocumentInvalid
	}

	log.Debugf("Fetched and cached document for : %s", d.DID())
	return document, nil

}
