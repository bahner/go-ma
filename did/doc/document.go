package doc

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
	log "github.com/sirupsen/logrus"
)

type Document struct {
	_                  struct{}             `cbor:",toarray"`
	Context            []string             `cbor:"@context,toarray" json:"@context"`
	Version            string               `cbor:"versionId" json:"versionId"`
	ID                 string               `cbor:"id" json:"id"`
	Controller         []string             `cbor:"controller,omitempty,toarray" json:"controller,omitempty"`
	VerificationMethod []VerificationMethod `cbor:"verificationMethod,omitempty,toarray" json:"verificationMethod,omitempty"`
	AssertionMethod    string               `cbor:"assertionMethod,omitempty" json:"assertionMethod,omitempty"`
	KeyAgreement       string               `cbor:"keyAgreement,omitempty" json:"keyAgreement,omitempty"`
	Proof              Proof                `cbor:"proof,omitempty" json:"proof,omitempty"`
}

func New(identifier string, controller string) (*Document, error) {

	log.Debugf("doc/new: identifier: %s", identifier)
	log.Debugf("doc/new: controller: %s", controller)
	_, err := did.NewFromDID(identifier)
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

func (d *Document) JSON() ([]byte, error) {
	bytes, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("doc/string: failed to marshal document to JSON: %w", err)
	}

	return bytes, nil
}

func (d *Document) CBOR() ([]byte, error) {
	bytes, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("doc/string: failed to marshal document to JSON: %w", err)
	}

	return bytes, nil
}
