package set

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
	"github.com/libp2p/go-libp2p/core/crypto"
)

// KeySet struct the encryption and signing keys are actual keys,
// but the IPFSKey is a reference to the IPFS key and holds names and paths.
// The key itself resides in IPFS.
// The PrivKey is a libp2p key.
type Keyset struct {
	Identity      crypto.PrivKey
	DID           did.DID
	EncryptionKey key.EncryptionKey
	SigningKey    key.SigningKey
}

// Creates new keyset from a name (typically fragment of a DID)
// This requires that the key is already in IPFS and that IPFS is running.
func New(d did.DID, identity crypto.PrivKey) (Keyset, error) {

	var err error

	encryptionKey, err := key.NewEncryptionKey(d)
	if err != nil {
		return Keyset{}, fmt.Errorf("newFromIPFSKey: %w", err)
	}

	signatureKey, err := key.NewSigningKey(d)
	if err != nil {
		return Keyset{}, fmt.Errorf("newFromIPFSKey: %w", err)
	}

	return Keyset{
		Identity:      identity,
		DID:           d,
		EncryptionKey: encryptionKey,
		SigningKey:    signatureKey,
	}, nil
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

	err = ks.DID.Validate()
	if err != nil {
		return fmt.Errorf("KeysetVerify: %w", err)
	}

	return nil
}

func (ks Keyset) IsValid() bool {
	return ks.Verify() == nil
}
