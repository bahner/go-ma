package key

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"maze.io/x/crypto/x25519"
)

type EncryptionKey struct {
	DID                string
	Name               string
	PrivKey            *x25519.PrivateKey
	PubKey             *x25519.PublicKey
	PublicKeyMultibase string
}

func GenerateEncryptionKey(name string) (EncryptionKey, error) {
	privKey, err := x25519.GenerateKey(rand.Reader)
	if err != nil {
		return EncryptionKey{}, err
	}
	pubKey := privKey.Public().(*x25519.PublicKey)

	publicKeyMultibase, err := internal.EncodePublicKeyMultibase(pubKey.Bytes(), ma.KEY_AGREEMENT_MULTICODEC_STRING)
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
