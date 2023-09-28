package vm

import "errors"

var (
	ErrInvalidVerificationMethod = errors.New("invalid verification method")
)

var VerificationPubkeyMethodMapping = map[string]string{
	"RsaPub":     "RsaVerificationKey2018",
	"Ed25519Pub": "Ed25519VerificationKey2018",
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
