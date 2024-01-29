package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/ipfs/boxo/path"
	log "github.com/sirupsen/logrus"
)

// Takes a DID and fetches the document from IPFS.
// Eg. fetch("did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner")
func fetch(didStr string) (*Document, error) {

	d, err := did.New(didStr)
	if err != nil {
		return nil, err
	}

	return fetchFromDID(d)

}

func fetchFromDID(d *did.DID) (*Document, error) {

	var document = &Document{}

	api := internal.GetIPFSAPI()

	_cid, err := internal.RootCID(d.Path(path.IPNSNamespace), false) // NB! Cached = false
	if err != nil {
		return nil, fmt.Errorf("failed to get CID from IPNS name: %w", err)
	}
	log.Debugf("Fetching CID: %s", _cid)

	node, err := api.Dag().Get(internal.GetContext(), _cid)
	if err != nil {
		return nil, fmt.Errorf("failed to get document from IPFS: %w", err)
	}

	err = cbor.Unmarshal(node.RawData(), document)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal document: %w", err)
	}

	if !document.isValid() {
		return nil, fmt.Errorf("document is invalid")
	}

	// Add fetched document to cache
	cache(document)

	log.Debugf("Fetched and cached document for : %s", d.String())
	return document, nil

}

func GetOrFetch(id string) (*Document, error) {

	doc, err := get(id)
	if err == nil {
		log.Debugf("doc/getorfetch: found document for %s in cache", id)
		return doc, err
	}
	log.Debugf("doc/getorfetch: failed to get document for %s from cache", id)

	// Then try to fetch from IPFS
	if doc == nil {
		return fetch(id)
	}
	log.Debugf("doc/getorfetch: failed to fetch document for %s from IPFS", id)

	return nil, fmt.Errorf("doc/getorfetch: failed to get document from cache or IPFS")

}
