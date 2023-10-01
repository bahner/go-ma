package keyset

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"crypto/ed25519"

	"github.com/multiformats/go-multibase"
)

type Ed25519PrivateKey struct {
	*ed25519.PrivateKey
}

func (k *Ed25519PrivateKey) PublicKey() crypto.PublicKey {
	return k.PrivateKey.Public().(ed25519.PublicKey)
}

func (k *Ed25519PrivateKey) PublicKeyMultibase() (string, error) {
	multibase, err := multibase.Encode(
		multibase.Base58BTC,
		k.PublicKey().(ed25519.PublicKey))
	if err != nil {
		return "", fmt.Errorf("keyset/ed25519: error multibase encoding public key: %s", err)
	}
	return multibase, nil
}

func (k *Ed25519PrivateKey) Sign(data []byte) ([]byte, error) {
	return ed25519.Sign(*k.PrivateKey, data), nil
}

func GenerateEd25519PrivateKey() (SignatureKey, error) {
	_, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &Ed25519PrivateKey{&privKey}, nil
}
