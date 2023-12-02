package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/ipfs/go-cid"
)

// Takes a DID and fetches the document from IPFS.
// Eg. Fetch("did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner")
func FetchFromDID(didStr string) (*Document, error) {

	d, err := did.NewFromDID(didStr)
	if err != nil {
		return nil, err
	}

	return Fetch(d.Identifier)

}

func Fetch(id string) (*Document, error) {

	var err error

	var document = &Document{}

	api := internal.GetIPSAPI()

	_cid, err := cid.Decode("/ipns" + id)
	if err != nil {
		return nil, fmt.Errorf("doc/fetch: failed to decode cid: %w", err)
	}

	// err = api.DagGet("/ipns/"+id, document)
	node, err := api.Dag().Get(internal.GetContext(), _cid)
	if err != nil {
		return nil, fmt.Errorf("doc/fetch: failed to get document from IPFS: %w", err)
	}

	err = cbor.Unmarshal(node.RawData(), document)
	if err != nil {
		return nil, fmt.Errorf("doc/fetch: failed to unmarshal document: %w", err)
	}

	return document, nil

}
