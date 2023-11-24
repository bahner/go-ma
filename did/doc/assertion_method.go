package doc

import (
	"fmt"

	"crypto/ed25519"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
)

func (d *Document) AssertionMethodPublicKey() (ed25519.PublicKey, error) {
	// Decode the multibase-encoded public key
	vm, err := d.GetVerificationMethodbyID(d.AssertionMethod)
	if err != nil {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Error getting verification method by ID: %w", err)
	}
	codec, pubKeyBytes, err := internal.DecodePublicKeyMultibase(vm.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Error decoding publicKeyMultibase: %w", err)
	}

	if codec != key.ASSERTION_METHOD_KEY_MULTICODEC_STRING {
		return nil, fmt.Errorf("doc/key_agreement_public_key: codec != %s", key.ASSERTION_METHOD_KEY_MULTICODEC_STRING)
	}

	// Check if the length of pubKeyBytes matches the expected length for an ed25519 public key
	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("doc/key_agreement_public_key: invalid public key size: expected %d, got %d", ed25519.PublicKeySize, len(pubKeyBytes))
	}

	// Convert the extracted bytes to a public key
	pubKey := make(ed25519.PublicKey, ed25519.PublicKeySize)
	copy(pubKey, pubKeyBytes)
	return pubKey, nil
}

func (d *Document) GetAssertionMethod() (VerificationMethod, error) {

	return d.GetVerificationMethodbyID(d.AssertionMethod)
}

func (d *Document) GetKeyAgreementMethod() (VerificationMethod, error) {

	return d.GetVerificationMethodbyID(d.KeyAgreement)
}
