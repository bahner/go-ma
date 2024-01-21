package doc

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	cbor "github.com/fxamacker/cbor/v2"
	log "github.com/sirupsen/logrus"
)

type Document struct {
	Context            []string             `cbor:"@context,toarray" json:"@context"`
	Version            string               `cbor:"version" json:"version"`
	ID                 string               `cbor:"id" json:"id"`
	Controller         []string             `cbor:"controller,toarray" json:"controller"`
	VerificationMethod []VerificationMethod `cbor:"verificationMethod,toarray" json:"verificationMethod"`
	AssertionMethod    string               `cbor:"assertionMethod" json:"assertionMethod"`
	KeyAgreement       string               `cbor:"keyAgreement" json:"keyAgreement"`
	Proof              Proof                `cbor:"proof" json:"proof"`
}

func New(identifier string, controller string) (*Document, error) {

	log.Debugf("doc/new: identifier: %s", identifier)
	log.Debugf("doc/new: controller: %s", controller)
	_, err := did.New(identifier)
	if err != nil {
		return nil, fmt.Errorf("doc/new: failed to parse DID: %w", err)
	}

	ctrller := []string{controller}

	doc := Document{
		Context:    DID_CONTEXT,
		Version:    ma.VERSION,
		ID:         identifier,
		Controller: ctrller,
	}
	log.Debugf("doc/new: doc: %v", doc)
	return &doc, nil
}

func (d *Document) MarshalToCBOR() ([]byte, error) {
	bytes, err := cbor.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("doc/string: failed to marshal document to CBOR: %w", err)
	}

	return bytes, nil
}

// Simple string representation of the document
// JSON or empty string on error
func (d *Document) String() string {

	bytes, err := json.Marshal(d)
	if err != nil {
		return ""
	}

	return string(bytes)
}

// GetOrCreate document from cache or fetch from IPFS
// identifier is a did string, eg. "did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#foo"
func GetOrCreate(identifier string) (*Document, error) {

	if exists(identifier) {
		log.Debugf("doc/GetOrCreate: document %s found in cache", identifier)
		return get(identifier)
	}

	doc, err := FetchFromDID(identifier)
	if err == nil {
		log.Debugf("doc/GetOrCreate: found previously published document for: %s", identifier)
		return doc, nil
	}

	doc, err = New(identifier, identifier)
	if err != nil {
		return nil, fmt.Errorf("doc/GetOrCreate: failed to create new document: %w", err)
	}

	// Add document to cache
	cache(doc)

	return doc, nil
}
