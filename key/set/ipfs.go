package set

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/key/ipfs"
)

// This creates a new keyset from an existing IPFS key.
func newFromIPFSKey(k ipfs.Key) (Keyset, error) {

	encryptionKey, err := key.NewEncryptionKey(k.Identifier)
	if err != nil {
		return Keyset{}, fmt.Errorf("newFromIPFSKey: %w", err)
	}

	signatureKey, err := key.NewSigningKey(k.Identifier)
	if err != nil {
		return Keyset{}, fmt.Errorf("newFromIPFSKey: %w", err)
	}

	d, err := did.New(k.Id)
	if err != nil {
		return Keyset{}, fmt.Errorf("newFromIPFSKey: %w", err)
	}

	return Keyset{
		DID:           d,
		IPFSKey:       k,
		EncryptionKey: encryptionKey,
		SigningKey:    signatureKey,
	}, nil
}
