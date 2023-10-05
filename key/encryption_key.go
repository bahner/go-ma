package key

import (
	"crypto"
	"fmt"
)

type EncryptionKey interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
	DID() string
	Name() string
	PublicKey() crypto.PublicKey
	PublicKeyMultibase() string
}

var encryptionKeyTypes = map[string]func(name string) (EncryptionKey, error){
	"x25519":        GenerateX25519PrivateKey,
	"x448":          GenerateX448PrivateKey,
	"kyber-ed25519": GenerateKyberEd25519PrivateKey,
}

func GenerateEncryptionKey(encryptionKeyType string, name string) (EncryptionKey, error) {
	if generateFunc, exists := encryptionKeyTypes[encryptionKeyType]; exists {
		return generateFunc(name)
	}
	return nil, fmt.Errorf("unsupported encryption key type: %s", encryptionKeyType)
}
