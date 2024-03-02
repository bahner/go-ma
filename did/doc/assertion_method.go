package doc

import (
	"fmt"

	"crypto/ed25519"

	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/multi"
)

func (d *Document) AssertionMethodPublicKey() (ed25519.PublicKey, error) {
	// Decode the multibase-encoded public key
	vm, err := d.GetVerificationMethodByID(d.AssertionMethod)
	if err != nil {
		return nil, ErrVerificationMethoddUnkownID
	}
	codec, pubKeyBytes, err := multi.PublicKeyMultibaseDecode(vm.PublicKeyMultibase)
	if err != nil {
		return nil, ErrPublicKeyMultibaseInvalid
	}

	if codec != key.ASSERTION_METHOD_KEY_MULTICODEC_STRING {
		return nil, ErrMultiCodecInvalid
	}

	// Check if the length of pubKeyBytes matches the expected length for an ed25519 public key
	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("invalid keysize %d. %w", len(pubKeyBytes), ErrPublicKeyLengthInvalid)
	}

	// Convert the extracted bytes to a public key
	pubKey := make(ed25519.PublicKey, ed25519.PublicKeySize)
	copy(pubKey, pubKeyBytes)
	return pubKey, nil
}

func (d *Document) GetAssertionMethod() (VerificationMethod, error) {

	return d.GetVerificationMethodByID(d.AssertionMethod)
}

func (d *Document) GetKeyAgreementMethod() (VerificationMethod, error) {

	return d.GetVerificationMethodByID(d.KeyAgreement)
}
