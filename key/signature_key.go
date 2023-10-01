package key

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
)

type SignatureKey interface {
	PublicKeyMultibase() string
	Sign(data []byte) ([]byte, error)
	DID() string
	Name() string
}

var signatureKeyTypes = map[string]func(name string) (SignatureKey, error){
	"ed25519": GenerateEd25519PrivateKey,
}

func GenerateSignatureKey(signatureKeyType string, name string) (SignatureKey, error) {

	if !internal.IsValidNanoID(name) {
		return nil, fmt.Errorf("keyset/generate_signature_key: invalid name : %s", name)
	}
	if generateFunc, exists := signatureKeyTypes[signatureKeyType]; exists {
		return generateFunc(name)
	}
	return nil, fmt.Errorf("unsupported encryption key type: %s", signatureKeyType)
}
