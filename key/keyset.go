package key

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	shell "github.com/ipfs/go-ipfs-api"
)

const EncryptionKeyType = "x25519"
const SignatureKeyType = "ed25519"

type Keyset interface {
	IPFSKey() *shell.Key
	EncryptionKey() EncryptionKey
	SignatureKey() SignatureKey
}

type keysetImpl struct {
	ipfsKey       *shell.Key
	encryptionKey *EncryptionKey
	signatureKey  *SignatureKey
}

func (k *keysetImpl) IPFSKey() *shell.Key {
	return k.ipfsKey
}

func (k *keysetImpl) EncryptionKey() EncryptionKey {
	return *k.encryptionKey
}

func (k *keysetImpl) SignatureKey() SignatureKey {
	return *k.signatureKey
}
func NewKeyset(name string, enc_type string, sig_type string) (Keyset, error) {
	encryptionKey, err := GenerateEncryptionKey(enc_type, name)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to generate encryption key: %w", err)
	}

	signatureKey, err := GenerateSignatureKey(sig_type, name)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to generate signature key: %w", err)
	}

	ipfsKey, err := internal.IPNSGetOrCreateKey(name)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to get or create key in IPFS: %w", err)
	}

	return &keysetImpl{
		ipfsKey:       ipfsKey,
		encryptionKey: &encryptionKey,
		signatureKey:  &signatureKey,
	}, nil
}

func NewKeysetFromDID(DID *did.DID) (Keyset, error) {

	return NewKeyset(DID.Fragment, EncryptionKeyType, SignatureKeyType)

}

func NewKeysetFromIPFSKey(ipfsKey *shell.Key) (Keyset, error) {

	return NewKeyset(ipfsKey.Name, EncryptionKeyType, SignatureKeyType)
}
