package keyset

import (
	"crypto"
)

type SignatureKey interface {
	PublicKey() crypto.PublicKey
	PublicKeyMultibase() (string, error)
	Sign(data []byte) ([]byte, error) // Maybe string
}

// var SignatureKeyTypes = map[string]func() (SignatureKey, error){
// 	"EC25519": GenerateEC25519PrivateKey,
// 	// Add "X448": GenerateX448Keypair, when you implement X448 key pair generation
// }
