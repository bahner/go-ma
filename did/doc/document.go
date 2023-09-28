package doc

import (
	"encoding/json"

	"github.com/bahner/go-ma/did/coll"
	"github.com/bahner/go-ma/did/vm"
)

type Document struct {
	Context            coll.Collection         `json:"@context"`
	ID                 string                  `json:"id"`
	Controller         coll.Collection         `json:"controller"`
	VerificationMethod []vm.VerificationMethod `json:"verificationMethod"`
	Signature          string                  `json:"signature,omitempty"`
}

func New(identifier string) (*Document, error) {
	context, err := coll.New(CONTEXT)
	if err != nil {
		return nil, err
	}
	doc := Document{
		Context: context,
		ID:      identifier,
	}
	return &doc, nil
}

func (d *Document) String() (string, error) {
	bytes, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	doc := string(bytes)

	return doc, nil
}
