package key

import (
	"crypto/rand"
	"fmt"

	"crypto/ed25519"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
)

type SigningKey struct {
	DID                string
	Name               string
	Type               string
	PrivKey            *ed25519.PrivateKey
	PubKey             *ed25519.PublicKey
	PublicKeyMultibase string
}

func (k *SigningKey) Sign(data []byte) ([]byte, error) {
	if !internal.IsValidEd25519PrivateKey(k.PrivKey) {
		return nil, fmt.Errorf("keyset/ed25519: invalid private key")
	}

	return ed25519.Sign(*k.PrivKey, data), nil
}

func GenerateSigningKey(name string) (SigningKey, error) {
	publicKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return SigningKey{}, err
	}

	publicKeyMultibase, err := internal.EncodePublicKeyMultibase(publicKey, ma.ASSERTION_METHOD_MULTICODEC_STRING)
	if err != nil {
		return SigningKey{}, fmt.Errorf("key/ed25519: error encoding public key multibase: %w", err)
	}

	return SigningKey{
		DID:                DID_KEY_PREFIX + publicKeyMultibase + "#" + name,
		Type:               ma.VERIFICATION_METHOD_KEY_TYPE,
		Name:               name,
		PrivKey:            &privKey,
		PubKey:             &publicKey,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}
