package key

import (
	"fmt"

	"crypto/ed25519"
	"crypto/rand"

	"github.com/bahner/go-ma/did"
	mf "github.com/bahner/go-ma/utils"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/multiformats/go-multicodec"
	log "github.com/sirupsen/logrus"
)

const (
	ASSERTION_METHOD_KEY_TYPE   = "MultiKey"
	ASSERTION_METHOD_MULTICODEC = multicodec.Ed25519Pub
)

type SigningKey struct {
	DID                did.DID
	Type               string
	PrivKey            ed25519.PrivateKey
	PubKey             ed25519.PublicKey
	PublicKeyMultibase string
}

func (k *SigningKey) Sign(data []byte) ([]byte, error) {
	err := validateSigningPrivateKey(k.PrivKey)
	if err != nil {
		return nil, err
	}

	return ed25519.Sign(k.PrivKey, data), nil
}

// Generates a signing key for the given identifier, ie. IPNS name.
// The IPNS key itself could be used, but lets stick to a more common DID structure.
func NewSigningKey(d did.DID) (SigningKey, error) {

	name, err := nanoid.New()
	if err != nil {
		return SigningKey{}, fmt.Errorf("NewSigningKey: %w", err)
	}

	// Create a unique identifier for the key
	d = did.New(d.IPNSName(), name)
	log.Info("Created new DID: ", d)

	publicKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return SigningKey{}, fmt.Errorf("NewSigningKey: %w", err)
	}

	publicKeyMultibase, err := mf.PublicKeyMultibaseEncode(ASSERTION_METHOD_MULTICODEC, publicKey)
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

	err := s.DID.Validate()
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

	encodedKey, err := mf.MultibaseDecode(s.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf("SigningKeyVerify: %w", err)
	}

	codec, _, err := mf.MulticodecDecode(encodedKey)
	if err != nil {
		return fmt.Errorf("SigningKeyVerify: %w", err)
	}

	if codec != ASSERTION_METHOD_MULTICODEC {
		return ErrInvalidSigningKeyMulticodec
	}

	return nil

}

func (s SigningKey) IsValid() bool {
	return s.Verify() == nil
}

func validateSigningPrivateKey(privKey ed25519.PrivateKey) error {
	if privKey == nil {
		return ErrPrivateKeyIsNil
	}

	if len(privKey) != ed25519.PrivateKeySize {
		return fmt.Errorf("invalid key length %d, should be %d: %w", len(privKey), ed25519.PrivateKeySize, ErrInvalidKeySize)
	}
	return nil
}
