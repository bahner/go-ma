package doc

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/internal"
)

type Document struct {
	Context            []string                `json:"@context"`
	ID                 string                  `json:"id"`
	Controller         []string                `json:"controller"`
	VerificationMethod []vm.VerificationMethod `json:"verificationMethod"`
	Signature          string                  `json:"signature,omitempty"`
}

func New(identifier string) (*Document, error) {

	_, err := did.Parse(identifier)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("doc/new: failed to parse DID: %v", err))
	}

	ctrller := []string{identifier}

	doc := Document{
		Context:    Context,
		ID:         identifier,
		Controller: ctrller,
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
