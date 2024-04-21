package doc

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/bahner/go-ma/did"
	cbor "github.com/fxamacker/cbor/v2"
	log "github.com/sirupsen/logrus"
)

type Document struct {
	Context            []string             `cbor:"@context,toarray" json:"@context"`
	ID                 string               `cbor:"id" json:"id"`
	Controller         []string             `cbor:"controller,toarray" json:"controller"`
	VerificationMethod []VerificationMethod `cbor:"verificationMethod,toarray" json:"verificationMethod"`
	AssertionMethod    string               `cbor:"assertionMethod" json:"assertionMethod"`
	KeyAgreement       string               `cbor:"keyAgreement" json:"keyAgreement"`
	Proof              Proof                `cbor:"proof" json:"proof"`
}

// Takes an identity DID and a controller DID. They can be the same.
func New(identity string, controller string) (*Document, error) {

	log.Debugf("doc/new: identifier: %s", identity)
	log.Debugf("doc/new: controller: %s", controller)

	err := did.Validate(identity)
	if err != nil {
		return nil, fmt.Errorf("doc/new: invalid identifier: %w", err)
	}

	ctrller := []string{identity}

	doc := Document{
		Context:    DID_CONTEXT,
		ID:         identity,
		Controller: ctrller,
	}
	doc.AddController(controller)
	log.Infof("doc/new: created new document for %s", identity)
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

func (d *Document) Equal(other *Document) bool {
	if d.ID != other.ID {
		return false
	}

	if d.KeyAgreement != other.KeyAgreement {
		return false
	}

	if d.AssertionMethod != other.AssertionMethod {
		return false
	}

	if !compareSlices(d.Context, other.Context) {
		return false
	}

	if !compareSlices(d.Controller, other.Controller) {
		return false
	}

	if d.Proof != other.Proof {
		return false
	}

	return true

}

func compareSlices(a []string, b []string) bool {

	slices.Sort(a)
	slices.Sort(b)

	return slices.Compare(a, b) == 0
}
