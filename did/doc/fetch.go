package doc

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	log "github.com/sirupsen/logrus"
)

func Fetch(identifier string) (*Document, error) {

	var document = &Document{}
	var err error

	// First we need to fetch the document from IPFS.
	// We need to parse the identifier out of the document.
	// in order to find the fragment, which *MUST* be the keyname.

	// Lookup short name of the identifier, ie. the fragment
	// The shortname is given to IPFS to lookup the actual key,
	// but that is transparent to use.
	// This gived us the possibility to change the key without
	// having to change the entity name within a given context.
	docdid, err := did.Parse(identifier)
	if err != nil {
		return nil, err
	}

	shell := internal.GetShell()

	// Fetch the document from IPFS
	// Start by getting an IO.Reader
	data, err := shell.Cat("/ipns/" + docdid.Identifier)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("doc/fetch: failed to fetch document from IPFS: %v\n", err))
	}
	defer data.Close() // Ensure the reader is closed once done
	log.Debugf("doc/fetch: data: %v", data)

	// The DIDDocuments aren't supposed to be that big,
	// so we can read it all into memory.
	content, err := io.ReadAll(data)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("doc/fetch: failed to read contents from: %v\n", err))
	}
	log.Debugf("doc/fetch: content: %s", content)
	// Unmarshal the document
	err = json.Unmarshal(content, document)
	if err != nil {
		return nil, err
	}

	retrieved_document, err := document.String()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Document: %s\n", retrieved_document)

	return document, nil
}
