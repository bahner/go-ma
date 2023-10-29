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
	SignatureKey  SignatureKey
}

func NewKeyset(name string) (Keyset, error) {
	encryptionKey, err := GenerateEncryptionKey(name)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate encryption key: %w", err)
	}

	signatureKey, err := GenerateSignatureKey(name)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate signature key: %w", err)
	}

	ipfsKey, err := internal.IPNSGetOrCreateKey(name)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to get or create key in IPFS: %w", err)
	}

	return Keyset{
		IPNSKey:       ipfsKey,
		EncryptionKey: encryptionKey,
		SignatureKey:  signatureKey,
	}, nil
}

func NewKeysetFromDID(DID *did.DID) (Keyset, error) {

	return NewKeyset(DID.Fragment)

}

func NewKeysetFromIPFSKey(ipfsKey *shell.Key) (Keyset, error) {

	return NewKeyset(ipfsKey.Name)
}
