package key

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	nanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/curve25519"
)

type EncryptionKey struct {
	DID                string
	Type               string
	PrivKey            [32]byte // Private key
	PubKey             [32]byte // Public key
	PublicKeyMultibase string
}

func GenerateEncryptionKey(identifier string) (EncryptionKey, error) {

	if !internal.IsValidIPNSName(identifier) {
		return EncryptionKey{}, fmt.Errorf("key/encryption: invalid identifier: %s", identifier)
	}

	// Generate a random private key
	var privKey [curve25519.ScalarSize]byte
	_, err := rand.Read(privKey[:])
	if err != nil {
		return EncryptionKey{}, err
	}

	// Calculate the corresponding public key
	var pubKey [curve25519.PointSize]byte
	curve25519.ScalarBaseMult(&pubKey, &privKey)

	// Encode the public key to multibase
	publicKeyMultibase, err := internal.EncodePublicKeyMultibase(pubKey[:], ma.KEY_AGREEMENT_MULTICODEC_STRING)
	if err != nil {
		return EncryptionKey{}, fmt.Errorf("key_generate: error encoding public key multibase: %w", err)
	}

	name, err := nanoid.New()
	if err != nil {
		return EncryptionKey{}, fmt.Errorf("key_generate: error generating nanoid: %w", err)
	}

	return EncryptionKey{
		DID:                ma.DID_PREFIX + identifier + "#" + name,
		Type:               ma.KEY_AGREEMENT_KEY_TYPE,
		PrivKey:            privKey,
		PubKey:             pubKey,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}
