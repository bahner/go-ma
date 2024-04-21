package doc

import (
	"fmt"

	"github.com/bahner/go-ma/key/set"
)

// Takes a keyset and generates a DID Document. Also takes a controller string
// which is the DID of the controller of the keyset. This is used to set the
// controller of the DID Document.
// It's OK to set it as the DID of the keyset.IPNSKey.DID, but it's not required.
func NewFromKeyset(k set.Keyset) (*Document, error) {

	return NewFromKeysetWithController(k, k.DID.Id)

}

func NewFromKeysetWithController(k set.Keyset, controller string) (*Document, error) {

	id := k.DID.Name.String()

	d, err := New(id, controller)
	if err != nil {
		return nil, err
	}

	encVm, err := NewVerificationMethod(
		id,
		id,
		"MultiKey",
		k.EncryptionKey.DID.Fragment,
		k.EncryptionKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("new_actor: Failed to create encryption verification method: %w", err)
	}
	d.AddVerificationMethod(encVm)
	d.KeyAgreement = encVm.ID

	assertVm, err := NewVerificationMethod(
		id,
		id,
		"MultiKey",
		k.SigningKey.DID.Fragment,
		k.SigningKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("new_actor: Failed to create assertion verification method: %w", err)
	}
	d.AddVerificationMethod(assertVm)
	d.AssertionMethod = assertVm.ID

	d.Sign(k.SigningKey, assertVm)

	return d, nil
}
