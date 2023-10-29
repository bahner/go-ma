package doc

import (
	"fmt"

	"crypto/ed25519"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
)

func (d *Document) AssertionMethodPublicKey() (ed25519.PublicKey, error) {

	// Decode the multibase-encoded public key
	codec, pubKeyBytes, err := internal.DecodePublicKeyMultibase(d.KeyAgreement.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("doc/key_agreement_public_key: Error decoding publicKeyMultibase: %s", err)
	}

	if codec != ma.ASSERTION_METHOD_MULTICODEC_STRING {
		return nil, fmt.Errorf("doc/key_agreement_public_key: codec != %s", ma.ASSERTION_METHOD_MULTICODEC_STRING)
	}

	// Convert the extracted bytes to a public key
	var pubKey ed25519.PublicKey
	copy(pubKey[:], pubKeyBytes)
	return pubKey, nil

}
