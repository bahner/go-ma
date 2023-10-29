package doc

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/internal"
	log "github.com/sirupsen/logrus"
)

type Document struct {
	Context            []string                `json:"@context"`
	ID                 string                  `json:"id"`
	Controller         []string                `json:"controller"`
	VerificationMethod []vm.VerificationMethod `json:"verificationMethod"`
	Signature          string                  `json:"signature,omitempty"`
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
	bytes, err := json.Marshal(d)
	if err != nil {
		return "", fmt.Errorf("doc/string: failed to marshal document to JSON: %w", err)
	}

	doc := string(bytes)

	return doc, nil
}
