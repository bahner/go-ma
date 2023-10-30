package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
	log "github.com/sirupsen/logrus"
)

type Document struct {
	_                  struct{}             `cbor:",toarray"`
	Context            []string             `cbor:"@context,toarray"`
	ID                 string               `cbor:"id"`
	Controller         []string             `cbor:"controller,omitempty,toarray"`
	VerificationMethod []VerificationMethod `cbor:"verificationMethod,omitempty,toarray"`
	AssertionMethod    string               `cbor:"assertionMethod,omitempty"`
	KeyAgreement       string               `cbor:"keyAgreement,omitempty"`
	Proof              Proof                `cbor:"proof,omitempty"`
}

func New(identifier string, controller string) (*Document, error) {

	log.Debugf("doc/new: identifier: %s", identifier)
	log.Debugf("doc/new: controller: %s", controller)
	_, err := did.Parse(identifier)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("doc/new: failed to parse DID: %v\n", err))
	}

	ctrller := []string{controller}

	doc := Document{
		Context:    DID_CONTEXT,
		ID:         identifier,
		Controller: ctrller,
	}
	log.Debugf("doc/new: doc: %v", doc)
	return &doc, nil
}

func (d *Document) String() (string, error) {
	bytes, err := cbor.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("doc/string: failed to marshal document to JSON: %w", err)
	}

	doc, err := internal.MultibaseEncode(bytes)
	if err != nil {
		return "", fmt.Errorf("doc/string: failed to multibase encode document: %w", err)
	}

	return doc, nil
}
