package key

import (
	"crypto/rand"
	"fmt"

	"crypto/ed25519"

	"github.com/bahner/go-ma/internal"
)

type SignatureKey struct {
	DID                string
	Name               string
	PrivKey            *ed25519.PrivateKey
	PubKey             *ed25519.PublicKey
	PublicKeyMultibase string
}

func (k *SignatureKey) Sign(data []byte) ([]byte, error) {
	if !isValidEd25519PrivateKey(k.PrivKey) {
		return nil, fmt.Errorf("keyset/ed25519: invalid private key")
	}

	return ed25519.Sign(*k.PrivKey, data), nil
}

func GenerateSignatureKey(name string) (SignatureKey, error) {
	publicKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return SignatureKey{}, err
	}

	publicKeyMultibase, err := internal.EncodePublicKeyMultibase(publicKey, "ed25519-pub")
	if err != nil {
		return SignatureKey{}, fmt.Errorf("key/ed25519: error encoding public key multibase: %w", err)
	}

	return SignatureKey{
		DID:                DID_KEY_PREFIX + publicKeyMultibase + "#" + name,
		Name:               name,
		PrivKey:            &privKey,
		PubKey:             &publicKey,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}

func isValidEd25519PrivateKey(privKey *ed25519.PrivateKey) bool {
	if privKey == nil || len(*privKey) != ed25519.PrivateKeySize {
		return false
	}
	return true
}
