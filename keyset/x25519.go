package keyset

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/multiformats/go-multibase"
	"maze.io/x/crypto/x25519"
)

type X25519PrivateKey struct {
	*x25519.PrivateKey
}

func (k *X25519PrivateKey) Encrypt(data []byte) ([]byte, error) {
	// Implementation here
	return nil, nil
}

func (k *X25519PrivateKey) Decrypt(data []byte) ([]byte, error) {
	// Implementation here
	return nil, nil
}

func (k *X25519PrivateKey) PublicKey() crypto.PublicKey {
	return k.PrivateKey.Public().(*x25519.PublicKey)
}

func (k *X25519PrivateKey) PublicKeyMultibase() (string, error) {
	multibase, err := multibase.Encode(
		multibase.Base58BTC,
		k.PublicKey().(*x25519.PublicKey).Bytes())
	if err != nil {
		return "", fmt.Errorf("keyset/x25519: error multibase encoding public key: %s", err)
	}
	return multibase, nil
}

func GenerateX25519PrivateKey() (EncryptionKey, error) {
	privKey, err := x25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &X25519PrivateKey{privKey}, nil
}
