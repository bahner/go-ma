package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
)

// Takes a DID and fetches the document from IPFS.
// Eg. Fetch("did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner")
func Fetch(d string) (*Document, error) {

	var document = &Document{}
	var err error

	docdid, err := did.Parse(d)
	if err != nil {
		return nil, err
	}

	shell := internal.GetShell()

	err = shell.DagGet("/ipns/"+docdid.Identifier, document)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("doc/fetch: failed to get document from IPFS: %v\n", err))
	}

	return document, nil
}
