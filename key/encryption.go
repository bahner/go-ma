package key

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/multi"
	nanoid "github.com/matoous/go-nanoid/v2"
	mc "github.com/multiformats/go-multicodec"
	"golang.org/x/crypto/curve25519"
)

const (
	// KEY_AGREEMENT_MULTICODEC_STRING = "x25519-pub"
	KEY_AGREEMENT_MULTICODEC = mc.X25519Pub
	KEY_AGREEMENT_KEY_TYPE   = "MultiKey"
)

type EncryptionKey struct {
	DID                did.DID
	Type               string
	PrivKey            [32]byte // Private key
	PubKey             [32]byte // Public key
	PublicKeyMultibase string
}

func NewEncryptionKey(d did.DID) (EncryptionKey, error) {

	name, err := nanoid.New()
	if err != nil {
		return EncryptionKey{}, fmt.Errorf("NewEncryptionKey: %w", err)
	}

	// We just mangle the base DID as this is not a pointer to a DID.
	d.Fragment = name

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
	publicKeyMultibase, err := multi.PublicKeyMultibaseEncode(KEY_AGREEMENT_MULTICODEC, pubKey[:])
	if err != nil {
		return EncryptionKey{}, fmt.Errorf("NewEncryptionKey: %w", err)
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

	err := k.DID.Validate()
	if err != nil {
		return fmt.Errorf("encryptionKey: %w", err)
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

	if !multi.IsValidMultibase(k.PublicKeyMultibase) {
		return ErrInvalidPublicKeyMultibase
	}

	if k.PublicKeyMultibase == "" {
		return ErrNoPublicKeyMultibase
	}

	key, err := multi.MultibaseDecode(k.PublicKeyMultibase)
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
