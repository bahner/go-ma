package keyset

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/cloudflare/circl/dh/x448"

	"github.com/multiformats/go-multibase"
)

type x448PrivateKey struct {
	privKey *x448.Key
	pubKey  *x448.Key
}

func (k *x448PrivateKey) Encrypt(data []byte) ([]byte, error) {
	// Implementation here
	return nil, nil
}

func (k *x448PrivateKey) Decrypt(data []byte) ([]byte, error) {
	// Implementation here
	return nil, nil
}

func (k *x448PrivateKey) PublicKey() crypto.PublicKey {

	return k.pubKey

}

func (k *x448PrivateKey) PublicKeyMultibase() (string, error) {

	pubKeyBytes := []byte((*k.pubKey)[:])

	multibase, err := multibase.Encode(
		multibase.Base58BTC,
		pubKeyBytes)
	if err != nil {
		return "", fmt.Errorf("keyset/x448: error multibase encoding public key: %s", err)
	}
	return multibase, nil
}

func GenerateX448PrivateKey() (EncryptionKey, error) {
	var secretKey x448.Key
	_, err := rand.Read(secretKey[:])
	if err != nil {
		return nil, fmt.Errorf("keyset/x448: error generating private key: %s", err)
	}

	var publicKey x448.Key
	x448.KeyGen(&publicKey, &secretKey)

	return &x448PrivateKey{
		privKey: &secretKey,
		pubKey:  &publicKey,
	}, nil
}
