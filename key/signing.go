package key

import (
	"fmt"

	"crypto/ed25519"
	"crypto/rand"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	nanoid "github.com/matoous/go-nanoid/v2"
	mc "github.com/multiformats/go-multicodec"
)

const (
	ASSERTION_METHOD_KEY_MULTICODEC_STRING = "ed25519-pub"
	ASSERTION_METHOD_KEY_TYPE              = "MultiKey"
)

type SigningKey struct {
	DID                string
	Type               string
	PrivKey            ed25519.PrivateKey
	PubKey             ed25519.PublicKey
	PublicKeyMultibase string
}

func (k *SigningKey) Sign(data []byte) ([]byte, error) {
	if !internal.IsValidEd25519PrivateKey(k.PrivKey) {
		return nil, fmt.Errorf("keyset/ed25519: invalid private key")
	}

	return ed25519.Sign(k.PrivKey, data), nil
}

// Generates a signing key for the given identifier, ie. IPNS name
func NewSigningKey(identifier string) (SigningKey, error) {

	if !internal.IsValidIPNSName(identifier) {
		return SigningKey{}, fmt.Errorf("key/ed25519: invalid identifier: %s", identifier)
	}

	publicKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return SigningKey{}, fmt.Errorf("key/signing: error generating key pair: %w", err)
	}

	publicKeyMultibase, err := internal.EncodePublicKeyMultibase(publicKey, ASSERTION_METHOD_KEY_MULTICODEC_STRING)
	if err != nil {
		return SigningKey{}, fmt.Errorf("key/ed25519: error encoding public key multibase: %w", err)
	}

	name, err := nanoid.New()
	if err != nil {
		return SigningKey{}, fmt.Errorf("key/ed25519: error generating nanoid: %w", err)
	}

	return SigningKey{
		DID:                ma.DID_PREFIX + identifier + "#" + name,
		Type:               ASSERTION_METHOD_KEY_TYPE,
		PrivKey:            privKey,
		PubKey:             publicKey,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}

func (s SigningKey) Verify() error {

	err := did.ValidateDID(s.DID)
	if err != nil {
		return err
	}

	if s.Type == "" {
		return fmt.Errorf("key/encryption: key has no type")
	}

	if s.Type != ASSERTION_METHOD_KEY_TYPE {
		return fmt.Errorf("key/encryption: key type is not %s", ASSERTION_METHOD_KEY_TYPE)
	}

	if len(s.PubKey) == 0 {
		return fmt.Errorf("key/encryption: key has no public key")
	}

	if len(s.PrivKey) == 0 {
		return fmt.Errorf("key/encryption: key has no private key")
	}

	if s.PublicKeyMultibase == "" {
		return fmt.Errorf("key/encryption: key has no public key")
	}

	key, err := internal.MultibaseDecode(s.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf("key/encryption: error decoding multibase: %w", err)
	}

	if key[0] != byte(mc.Ed25519Pub) {
		return fmt.Errorf("key/encryption: invalid multicodec")
	}

	return nil

}

func (s SigningKey) IsValid() bool {
	return s.Verify() == nil
}
