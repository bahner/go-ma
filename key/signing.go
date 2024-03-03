package key

import (
	"fmt"

	"crypto/ed25519"
	"crypto/rand"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/multi"
	nanoid "github.com/matoous/go-nanoid/v2"
	mc "github.com/multiformats/go-multicodec"
)

const (
	ASSERTION_METHOD_KEY_TYPE   = "MultiKey"
	ASSERTION_METHOD_MULTICODEC = mc.Ed25519Pub
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

	name, err := nanoid.New()
	if err != nil {
		return SigningKey{}, fmt.Errorf("NewSigningKey: %w", err)
	}

	d, err := did.New(ma.DID_PREFIX + identifier + "#" + name)
	if err != nil {
		return SigningKey{}, fmt.Errorf("NewSigningKey: %w", err)
	}

	publicKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return SigningKey{}, fmt.Errorf("NewSigningKey: %w", err)
	}

	publicKeyMultibase, err := multi.PublicKeyMultibaseEncode(ASSERTION_METHOD_MULTICODEC, publicKey)
	if err != nil {
		return SigningKey{}, fmt.Errorf("NewSigningKey: %w", err)
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

	key, err := multi.MultibaseDecode(s.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf("SigningKeyVerify: %w", err)
	}

	if key[0] != byte(ASSERTION_METHOD_MULTICODEC) {
		return ErrInvalidMulticodec
	}

	return nil

}

func (s SigningKey) IsValid() bool {
	return s.Verify() == nil
}
