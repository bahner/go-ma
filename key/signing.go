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
	DID                did.DID
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

	if !did.IsValidIdentifier(identifier) {
		return SigningKey{}, fmt.Errorf("key/signing: identifier %s: %w", identifier, did.ErrInvalidIdentifier)

	}

	name, err := nanoid.New()
	if err != nil {
		return SigningKey{}, fmt.Errorf("key/signing: error generating nanoid: %w", err)
	}

	d, err := did.New(ma.DID_PREFIX + identifier + "#" + name)
	if err != nil {
		return SigningKey{}, fmt.Errorf("key/ed25519: error creating DID: %w", err)
	}

	publicKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return SigningKey{}, fmt.Errorf("key/signing: error generating key pair: %w", err)
	}

	publicKeyMultibase, err := internal.EncodePublicKeyMultibase(publicKey, ASSERTION_METHOD_KEY_MULTICODEC_STRING)
	if err != nil {
		return SigningKey{}, fmt.Errorf("key/ed25519: error encoding public key multibase: %w", err)
	}

	return SigningKey{
		DID:                d,
		Type:               ASSERTION_METHOD_KEY_TYPE,
		PrivKey:            privKey,
		PubKey:             publicKey,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}

func (s SigningKey) Verify() error {

	err := s.DID.Verify()
	if err != nil {
		return err
	}

	if s.Type == "" {
		return ErrNoType
	}

	if s.Type != ASSERTION_METHOD_KEY_TYPE {
		return ErrInvalidAssertionMethodType
	}

	if len(s.PubKey) == 0 {
		return ErrNoPublicKey
	}

	if len(s.PrivKey) == 0 {
		return ErrNoPrivateKey
	}

	if s.PublicKeyMultibase == "" {
		return ErrNoPublicKeyMultibase
	}

	key, err := internal.MultibaseDecode(s.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf("key/encryption: error decoding multibase: %w", err)
	}

	if key[0] != byte(mc.Ed25519Pub) {
		return ErrInvalidMulticodec
	}

	return nil

}

func (s SigningKey) IsValid() bool {
	return s.Verify() == nil
}
