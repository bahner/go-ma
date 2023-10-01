package keyset

import (
	"crypto"
	"fmt"
)

type EncryptionKey interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
	PublicKey() crypto.PublicKey
	PublicKeyMultibase() (string, error)
}

var encryptionKeyTypes = map[string]func() (EncryptionKey, error){
	"x25519": GenerateX25519PrivateKey,
	"x448":   GenerateX448PrivateKey,
}

func GenerateEncryptionKey(encryptionKeyType string) (EncryptionKey, error) {
	if generateFunc, exists := encryptionKeyTypes[encryptionKeyType]; exists {
		return generateFunc()
	}
	return nil, fmt.Errorf("unsupported encryption key type: %s", encryptionKeyType)
}
