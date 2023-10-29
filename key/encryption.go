package key

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"golang.org/x/crypto/curve25519"
)

type EncryptionKey struct {
	DID                string
	Name               string
	PrivKey            [32]byte // Private key
	PubKey             [32]byte // Public key
	PublicKeyMultibase string
}

func GenerateEncryptionKey(name string) (EncryptionKey, error) {
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

	return EncryptionKey{
		DID:                DID_KEY_PREFIX + publicKeyMultibase + "#" + name,
		Name:               name,
		PrivKey:            privKey,
		PubKey:             pubKey,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}
