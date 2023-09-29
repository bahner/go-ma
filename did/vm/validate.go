package vm

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidVerificationMethod = errors.New("invalid verification method")
)

// VerificationPubkeyMethodMapping maps multicodec codes to DID Verification Methods
// Hence the key string comes from the go-multicodec package
// and is not something we can change.
// The value string is the DID Verification Method
// and is something we can change, if we meet the requirements.
var VerificationPubkeyMethodMapping = map[string]string{
	"rsa-pub":     "RsaVerificationKey2020",
	"ed25519-pub": "Ed25519VerificationKey2020",
}

func ValidVerificationMethods() []string {
	var methods []string
	for _, value := range VerificationPubkeyMethodMapping {
		methods = append(methods, value)
	}
	return methods
}

func IsValidVerificationMethod(method string) bool {
	for _, value := range VerificationPubkeyMethodMapping {
		if method == value {
			return true
		}
	}
	return false
}

func getVerificationMethodForKeyType(keyType string) (string, error) {
	for vmKeyType, value := range VerificationPubkeyMethodMapping {
		if keyType == vmKeyType {
			return value, nil
		}
	}
	return "", fmt.Errorf("no verification matches key type: %s", keyType)
}
