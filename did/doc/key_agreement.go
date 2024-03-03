package doc

import (
	"fmt"

	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/multi"
	"golang.org/x/crypto/curve25519"
)

func (d *Document) KeyAgreementPublicKeyBytes() ([]byte, error) {

	vm, err := d.GetVerificationMethodByID(d.KeyAgreement)
	if err != nil {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Error getting verification method by ID: %w", err)
	}
	// Decode the multibase-encoded public key
	codec, pubKeyBytes, err := multi.PublicKeyMultibaseDecode(vm.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Error decoding publicKeyMultibase: %w", err)
	}

	if codec != key.KEY_AGREEMENT_MULTICODEC {
		return nil, ErrMultiCodecInvalid
	}

	// Check if the number of bytes is correct for a curve25519 public key
	if len(pubKeyBytes) != curve25519.PointSize {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Invalid number of bytes. Expected %d, got %d",
			curve25519.PointSize, len(pubKeyBytes))
	}

	return pubKeyBytes, nil

}
