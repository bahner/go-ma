package key

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/multiformats/go-multibase"
	"maze.io/x/crypto/x25519"
)

type X25519PrivateKey struct {
	privKey            *x25519.PrivateKey
	pubKey             *x25519.PublicKey
	publicKeyMultibase string
	name               string
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
	return k.pubKey
}

func (k *X25519PrivateKey) PublicKeyMultibase() string {
	return k.publicKeyMultibase
}

func (k *X25519PrivateKey) DID() string {
	return prefix + k.publicKeyMultibase + "#" + k.name
}

func (k *X25519PrivateKey) Name() string {
	return k.name
}

func GenerateX25519PrivateKey(name string) (EncryptionKey, error) {
	privKey, err := x25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	pubKey := privKey.Public().(*x25519.PublicKey)

	publicKeyMultibase, err := multibase.Encode(
		ma.MULTIBASE_ENCODING,
		pubKey.Bytes())
	if err != nil {
		return nil, fmt.Errorf("keyset/x25519: error multibase encoding public key: %s", err)
	}

	return &X25519PrivateKey{
		privKey:            privKey,
		pubKey:             pubKey,
		publicKeyMultibase: publicKeyMultibase,
		name:               name}, nil
}
