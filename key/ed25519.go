package key

import (
	"crypto/rand"
	"fmt"

	"crypto/ed25519"

	"github.com/bahner/go-ma/internal"
)

type Ed25519PrivateKey struct {
	privKey            *ed25519.PrivateKey
	pubKey             *ed25519.PublicKey
	publicKeyMultibase string
	name               string
}

func (k *Ed25519PrivateKey) PublicKey() *ed25519.PublicKey {
	return k.pubKey
}

func (k *Ed25519PrivateKey) PublicKeyMultibase() string {
	return k.publicKeyMultibase
}

func (k *Ed25519PrivateKey) DID() string {
	return KEY_PREFIX + k.publicKeyMultibase + "#" + k.name
}

func (k *Ed25519PrivateKey) Name() string {
	return k.name
}

func (k *Ed25519PrivateKey) Sign(data []byte) ([]byte, error) {
	if !isValidEd25519PrivateKey(k.privKey) {
		return nil, fmt.Errorf("keyset/ed25519: invalid private key")
	}

	return ed25519.Sign(*k.privKey, data), nil
}
func GenerateEd25519PrivateKey(name string) (SignatureKey, error) {
	publicKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	publicKeyMultibase, err := internal.EncodePublicKeyMultibase(publicKey, "ed25519-pub")
	if err != nil {
		return nil, fmt.Errorf("key/ed25519: error encoding public key multibase: %w", err)
	}

	return &Ed25519PrivateKey{
		privKey:            &privKey,
		pubKey:             &publicKey,
		publicKeyMultibase: publicKeyMultibase,
		name:               name}, nil
}

func isValidEd25519PrivateKey(privKey *ed25519.PrivateKey) bool {
	if privKey == nil || len(*privKey) != ed25519.PrivateKeySize {
		return false
	}
	return true
}
