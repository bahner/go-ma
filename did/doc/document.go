package doc

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/bahner/go-ma/did"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/ipfs/boxo/path"
	log "github.com/sirupsen/logrus"
)

// How many fields to add to the node. Used for sanity checking.
const NUM_NODE_FIELDS = 9

type Document struct {
	Context    []string `cbor:"@context,toarray" json:"@context"`
	ID         string   `cbor:"id" json:"id"`
	Controller []string `cbor:"controller,toarray" json:"controller"`
	// Identity is a cid for the secret keyset of the entity.
	// It's contents should be encrypted and multibase encoded, but the structure is not defined here.
	// Each language should have it's own way of parsing the contents of the keyset.
	Identity string `cbor:"identity" json:"identity"`
	// The node to dial for communication with the entity.
	Node               Node                 `cbor:"node" json:"node"`
	VerificationMethod []VerificationMethod `cbor:"verificationMethod,toarray" json:"verificationMethod"`
	AssertionMethod    string               `cbor:"assertionMethod" json:"assertionMethod"`
	KeyAgreement       string               `cbor:"keyAgreement" json:"keyAgreement"`
	Proof              Proof                `cbor:"proof" json:"proof"`
	immutablePath      path.ImmutablePath   // This isn't published
	did                did.DID              // This isn't published
}

// Takes an identity DID and a controller DID. They can be the same.
func New(identity did.DID, controller did.DID) (*Document, error) {

	log.Debugf("doc/new: identifier: %s", identity.Id)
	log.Debugf("doc/new: controller: %s", controller.Id)

	ctrller := []string{controller.Id}

	doc := Document{
		Context:    DID_CONTEXT,
		ID:         identity.Id,
		Controller: ctrller,
		did:        identity,
	}
	doc.AddController(controller.Id)
	log.Infof("doc/new: created new document for %s", identity)
	return &doc, nil
}

func (d *Document) MarshalToCBOR() ([]byte, error) {

	marshalD := Document{
		Context:            d.Context,
		ID:                 d.ID,
		Controller:         d.Controller,
		Node:               d.Node,
		Identity:           d.Identity,
		VerificationMethod: d.VerificationMethod,
		AssertionMethod:    d.AssertionMethod,
		KeyAgreement:       d.KeyAgreement,
		Proof:              d.Proof,
	}

	bytes, err := cbor.Marshal(marshalD)
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

func (d *Document) Validate() error {

	err := d.validateContexts()
	if err != nil {
		return fmt.Errorf("doc/Validate: %w", err)
	}

	err = d.validateControllers()
	if err != nil {
		return fmt.Errorf("doc/Validate: %w", err)
	}

	err = validateNode(d.Node)
	if err != nil {
		return fmt.Errorf("doc/Validate: %w", err)
	}

	for _, vm := range d.VerificationMethod {
		err = vm.validateVerificationMethod()
		if err != nil {
			return fmt.Errorf("doc/Validate: %w", err)
		}
	}

	err = validateIdentity(d.Identity)
	if err != nil {
		return fmt.Errorf("doc/Validate: %w", err)
	}

	err = d.Verify()
	if err != nil {
		return fmt.Errorf("doc/Validate: %w", err)
	}

	return nil
}
