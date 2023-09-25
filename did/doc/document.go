package doc

import (
	"encoding/json"
	"fmt"

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

func (d *Document) AddVerificationMethod(vm vm.VerificationMethod) error {
	for _, existingVM := range d.VerificationMethod {
		if existingVM.ID == vm.ID {
			return fmt.Errorf("verification method with ID %s already exists", vm.ID)
		}
	}
	d.VerificationMethod = append(d.VerificationMethod, vm)
	return nil
}

func (d *Document) RemoveVerificationMethod(id string) {
	filtered := d.VerificationMethod[:0]
	for _, v := range d.VerificationMethod {
		if v.ID != id {
			filtered = append(filtered, v)
		}
	}
	d.VerificationMethod = filtered
}

func (d *Document) String() (string, error) {
	bytes, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	doc := string(bytes)

	return doc, nil
}
