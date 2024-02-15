package key

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	nanoid "github.com/matoous/go-nanoid/v2"
	mc "github.com/multiformats/go-multicodec"
	"golang.org/x/crypto/curve25519"
)

const (
	KEY_AGREEMENT_MULTICODEC_STRING = "x25519-pub"
	KEY_AGREEMENT_KEY_TYPE          = "MultiKey"
)

type EncryptionKey struct {
	DID                did.DID
	Type               string
	PrivKey            [32]byte // Private key
	PubKey             [32]byte // Public key
	PublicKeyMultibase string
}

func NewEncryptionKey(identifier string) (EncryptionKey, error) {

	if !internal.IsValidIPNSName(identifier) {
		return EncryptionKey{}, fmt.Errorf("key/encryption: invalid identifier: %s", identifier)
	}

	name, err := nanoid.New()
	if err != nil {
		return EncryptionKey{}, fmt.Errorf("key_generate: error generating nanoid: %w", err)
	}

	d, err := did.New(ma.DID_PREFIX + identifier + "#" + name)
	if err != nil {
		return EncryptionKey{}, fmt.Errorf("key_generate: error creating DID: %w", err)
	}

	// Generate a random private key
	var privKey [curve25519.ScalarSize]byte
	_, err = rand.Read(privKey[:])
	if err != nil {
		return EncryptionKey{}, err
	}

	// Calculate the corresponding public key
	var pubKey [curve25519.PointSize]byte
	curve25519.ScalarBaseMult(&pubKey, &privKey)

	// Encode the public key to multibase
	publicKeyMultibase, err := internal.EncodePublicKeyMultibase(pubKey[:], KEY_AGREEMENT_MULTICODEC_STRING)
	if err != nil {
		return EncryptionKey{}, fmt.Errorf("key_generate: error encoding public key multibase: %w", err)
	}

	return EncryptionKey{
		DID:                d,
		Type:               KEY_AGREEMENT_KEY_TYPE,
		PrivKey:            privKey,
		PubKey:             pubKey,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}

func (k EncryptionKey) Verify() error {

	err := k.DID.Verify()
	if err != nil {
		return fmt.Errorf("key/encryption: %w", err)
	}

	if k.Type == "" {
		return ErrNoType
	}

	if k.Type != KEY_AGREEMENT_KEY_TYPE {
		return ErrInvalidKeyAgreementType
	}

	if k.PubKey == [curve25519.ScalarSize]byte{} {
		return ErrNoPublicKey
	}

	if k.PrivKey == [curve25519.PointSize]byte{} {
		return ErrNoPrivateKey
	}

	if k.PublicKeyMultibase == "" {
		return ErrNoPublicKeyMultibase
	}

	if !internal.IsValidMultibase(k.PublicKeyMultibase) {
		return ErrInvalidPublicKeyMultibase
	}

	if k.PublicKeyMultibase == "" {
		return ErrNoPublicKeyMultibase
	}

	key, err := internal.MultibaseDecode(k.PublicKeyMultibase)
	if err != nil {
		return ErrInvalidPublicKeyMultibase
	}

	if key[0] != byte(mc.X25519Pub) {
		return ErrInvalidMulticodec
	}

	return nil

}

func (k EncryptionKey) IsValid() bool {
	return k.Verify() == nil
}
