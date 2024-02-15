package set

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/key/ipfs"
)

// This creates a new keyset from an existing IPFS key.
func newFromIPFSKey(k ipfs.Key) (Keyset, error) {

	encryptionKey, err := key.NewEncryptionKey(k.IPNSName)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate encryption key: %w", err)
	}

	signatureKey, err := key.NewSigningKey(k.IPNSName)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate signature key: %w", err)
	}

	d, err := did.New(k.Id)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to get or create DID: %w", err)
	}

	return Keyset{
		DID:           d,
		IPFSKey:       k,
		EncryptionKey: encryptionKey,
		SigningKey:    signatureKey,
	}, nil
}
