package doc

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	cbor "github.com/fxamacker/cbor/v2"
	log "github.com/sirupsen/logrus"
)

type Document struct {
	Context            []string             `cbor:"@context,toarray"`
	Version            string               `cbor:"version"`
	ID                 string               `cbor:"id"`
	Controller         []string             `cbor:"controller,toarray"`
	VerificationMethod []VerificationMethod `cbor:"verificationMethod,toarray"`
	AssertionMethod    string               `cbor:"assertionMethod"`
	KeyAgreement       string               `cbor:"keyAgreement"`
	Proof              Proof                `cbor:"proof"`
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

func GetOrCreate(identifier string) (*Document, error) {

	if exists(identifier) {
		log.Debugf("doc/new: document %s found in cache", identifier)
		return get(identifier)
	}

	doc, err := New(identifier, identifier)
	if err != nil {
		return nil, fmt.Errorf("doc/new: failed to create new document: %w", err)
	}

	// Add document to cache
	cache(doc)

	return doc, nil
}
