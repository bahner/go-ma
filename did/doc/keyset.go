package doc

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key/set"
)

// Takes a keyset and generates a DID Document. Also takes a controller string
// which is the DID of the controller of the keyset. This is used to set the
// controller of the DID Document.
// It's OK to set it as the DID of the keyset.IPNSKey.DID, but it's not required.
func NewFromKeyset(k *set.Keyset) (*Document, error) {

	return NewFromKeysetWithController(k, k.IPFSKey.DID.String())

}

func NewFromKeysetWithController(k *set.Keyset, controller string) (*Document, error) {

	id := k.IPFSKey.DID.String()

	d, err := New(id, controller)
	if err != nil {
		return nil, err
	}

	encVm, err := NewVerificationMethod(
		id,
		id,
		"MultiKey",
		internal.GetDIDFragment(k.EncryptionKey.DID),
		k.EncryptionKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("new_actor: Failed to create encryption verification method: %w", err)
	}
	err = d.AddVerificationMethod(encVm)
	if err != nil {
		return nil, fmt.Errorf("new_actor: Failed to add encryption verification method to DOC: %w", err)
	}
	d.KeyAgreement = encVm.ID

	assertVm, err := NewVerificationMethod(
		id,
		id,
		"MultiKey",
		internal.GetDIDFragment(k.SigningKey.DID),
		k.SigningKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("new_actor: Failed to create assertion verification method: %w", err)
	}
	err = d.AddVerificationMethod(assertVm)
	if err != nil {
		return nil, fmt.Errorf("new_actor: Failed to add assertion verification method to DOC: %w", err)
	}
	d.AssertionMethod = assertVm.ID

	d.Sign(k.SigningKey, assertVm)

	return d, nil
}
