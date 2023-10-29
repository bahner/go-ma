package doc

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
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

	// Decode the multibase-encoded public key
	codec, pubKeyBytes, err := internal.DecodePublicKeyMultibase(d.KeyAgreement.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Error decoding publicKeyMultibase: %s", err)
	}

	if codec != ma.KEY_AGREEMENT_MULTICODEC_STRING {
		return nil, fmt.Errorf("doc/key_agreement_public_key: codec != %s", ma.KEY_AGREEMENT_MULTICODEC_STRING)
	}

	// Check if the number of bytes is correct for a curve25519 public key
	if len(pubKeyBytes) != curve25519.PointSize {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Invalid number of bytes. Expected %d, got %d",
			curve25519.PointSize, len(pubKeyBytes))
	}

	return pubKeyBytes, nil

}
