package keyset

import (
	"fmt"

	"github.com/bahner/go-ma/did"
)

type Keyset interface {
	DID() *did.DID
	EncryptionKey() (*EncryptionKey, error)
	SignatureKey() (*SignatureKey, error)
}

type keysetImpl struct {
	did           *did.DID
	encryptionKey *EncryptionKey
	signatureKey  *SignatureKey
}

func (k *keysetImpl) DID() *did.DID {
	return k.did
}

func (k *keysetImpl) EncryptionKey() (*EncryptionKey, error) {
	return k.encryptionKey, nil
}

func (k *keysetImpl) SignatureKey() (*SignatureKey, error) {
	return k.signatureKey, nil
}

func New(DID *did.DID, encType string, sigType string) (Keyset, error) {
	encryptionKey, err := GenerateEncryptionKey(encType)
	if err != nil {
		return nil, fmt.Errorf("failed to generate encryption key: %w", err)
	}

	return &keysetImpl{
		did:           DID,
		encryptionKey: &encryptionKey,
	}, nil
}
