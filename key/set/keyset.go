package set

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
	log "github.com/sirupsen/logrus"
)

// KeySet struct the encryption and signing keys are actual keys,
// but the IPFSKey is a reference to the IPFS key and holds names and paths.
// The key itself resides in IPFS.
type Keyset struct {
	DID           did.DID
	EncryptionKey key.EncryptionKey
	SigningKey    key.SigningKey
}

// Creates new keyset from a name (typically fragment of a DID)
// This requires that the key is already in IPFS and that IPFS is running.
func GetOrCreate(name string) (Keyset, error) {

	var err error

	d, err := did.GetOrCreate(name)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetGetOrCreate: %w", err)
	}
	log.Infof("Created new key in IPFS: %v", d)

	ks, err := newFromDID(d)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetGetOrCreate: %w", err)
	}
	return ks, nil
}

func (ks Keyset) Verify() error {

	err := ks.EncryptionKey.Verify()
	if err != nil {
		return fmt.Errorf("KeysetVerify: %w", err)
	}

	err = ks.SigningKey.Verify()
	if err != nil {
		return fmt.Errorf("KeysetVerify: %w", err)
	}

	err = ks.DID.Verify()
	if err != nil {
		return fmt.Errorf("KeysetVerify: %w", err)
	}

	return nil
}

func (ks Keyset) IsValid() bool {
	return ks.Verify() == nil
}

// This creates a new keyset from an existing DID
func newFromDID(d did.DID) (Keyset, error) {

	encryptionKey, err := key.NewEncryptionKey(d.Identifier)
	if err != nil {
		return Keyset{}, fmt.Errorf("newFromIPFSKey: %w", err)
	}

	signatureKey, err := key.NewSigningKey(d.Identifier)
	if err != nil {
		return Keyset{}, fmt.Errorf("newFromIPFSKey: %w", err)
	}

	return Keyset{
		DID:           d,
		EncryptionKey: encryptionKey,
		SigningKey:    signatureKey,
	}, nil
}
