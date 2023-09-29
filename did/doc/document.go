package doc

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/coll"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/internal"
)

type Document struct {
	Context            coll.Collection         `json:"@context"`
	ID                 string                  `json:"id"`
	Controller         coll.Collection         `json:"controller"`
	VerificationMethod []vm.VerificationMethod `json:"verificationMethod"`
	Signature          string                  `json:"signature,omitempty"`
}

func New(identifier string) (*Document, error) {

	_, err := did.Parse(identifier)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("doc/new: failed to parse DID: %v", err))
	}

	// if !did.IsValidDID(identifier) {
	// 	return nil, fmt.Errorf("invalid DID: %s", identifier)
	// }

	context, err := coll.New(CONTEXT)
	if err != nil {
		return nil, err
	}

	// Set identify as default controller
	ctrller, err := coll.New(identifier)
	if err != nil {
		return nil, fmt.Errorf("failed to create controller: %v", err)
	}

	doc := Document{
		Context:    context,
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
