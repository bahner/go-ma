package key

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/cloudflare/circl/dh/x448"

	"github.com/multiformats/go-multibase"
)

type x448PrivateKey struct {
	privKey            *x448.Key
	pubKey             *x448.Key
	publicKeyMultibase string
	name               string
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

func (k *x448PrivateKey) PublicKeyMultibase() string {

	return k.publicKeyMultibase
}

func (k *x448PrivateKey) DID() string {

	return prefix + k.publicKeyMultibase + "#" + k.name
}

func (k *x448PrivateKey) Name() string {

	return k.name
}

func GenerateX448PrivateKey(name string) (EncryptionKey, error) {
	var secretKey x448.Key
	_, err := rand.Read(secretKey[:])
	if err != nil {
		return nil, fmt.Errorf("keyset/x448: error generating private key: %s", err)
	}

	var publicKey x448.Key
	x448.KeyGen(&publicKey, &secretKey)

	publicKeyMultibase, err := multibase.Encode(
		ma.MULTIBASE_ENCODING,
		publicKey[:])
	if err != nil {
		return nil, fmt.Errorf("keyset/x448: error multibase encoding public key: %s", err)
	}

	return &x448PrivateKey{
		privKey:            &secretKey,
		pubKey:             &publicKey,
		publicKeyMultibase: publicKeyMultibase,
		name:               name,
	}, nil
}
