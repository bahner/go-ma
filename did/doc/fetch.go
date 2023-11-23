package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
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

	var document = &Document{}

	shell := internal.GetShell()

	err = shell.DagGet("/ipns/"+id, document)
	if err != nil {
		return nil, fmt.Errorf("doc/fetch: failed to get document from IPFS: %w", err)
	}

	return document, nil

}
