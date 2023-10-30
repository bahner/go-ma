package key

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	shell "github.com/ipfs/go-ipfs-api"
)

type Keyset struct {
	IPNSKey       *shell.Key
	EncryptionKey EncryptionKey
	SigningKey    SigningKey
}

func NewKeyset(id *did.DID) (Keyset, error) {
	encryptionKey, err := GenerateEncryptionKey(id.Identifier)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate encryption key: %w", err)
	}

	signatureKey, err := GenerateSigningKey(id.Identifier)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate signature key: %w", err)
	}

	ipfsKey, err := internal.GetOrCreateIPNSKey(id.Fragment)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to get or create key in IPFS: %w", err)
	}

	return Keyset{
		IPNSKey:       ipfsKey,
		EncryptionKey: encryptionKey,
		SigningKey:    signatureKey,
	}, nil
}

func NewKeysetFromDID(DID *did.DID) (Keyset, error) {

	return NewKeyset(DID)

}
