package doc

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	"golang.org/x/crypto/curve25519"
)

// func (d *Document) KeyAgreementPublicKey() (crypto.PublicKey, error) {

// 	pubKeyBytes, err := d.KeyAgreementPublicKeyBytes()
// 	if err != nil {
// 		return nil, fmt.Errorf("doc/key_agreement_public_key: Error getting public key bytes: %s", err)
// 	}

// 	// Convert the extracted bytes to a public key
// 	var pubKey crypto.PublicKey
// 	copy(pubKey.(*[curve25519.PointSize]byte)[:], pubKeyBytes)

// 	return pubKey, nil

// }

func (d *Document) KeyAgreementPublicKeyBytes() ([]byte, error) {

	vm, err := d.GetVerificationMethodbyID(d.KeyAgreement)
	if err != nil {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Error getting verification method by ID: %w", err)
	}
	// Decode the multibase-encoded public key
	codec, pubKeyBytes, err := internal.DecodePublicKeyMultibase(vm.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Error decoding publicKeyMultibase: %w", err)
	}

	if codec != key.KEY_AGREEMENT_MULTICODEC_STRING {
		return nil, fmt.Errorf("doc/key_agreement_public_key: codec != %s", key.KEY_AGREEMENT_MULTICODEC_STRING)
	}

	// Check if the number of bytes is correct for a curve25519 public key
	if len(pubKeyBytes) != curve25519.PointSize {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Invalid number of bytes. Expected %d, got %d",
			curve25519.PointSize, len(pubKeyBytes))
	}

	return pubKeyBytes, nil

}
