package vm

import (
	"github.com/bahner/go-ma/internal"
)

// We only use multikeys, even for RSA keys.
// This is because we want to support multiple key types.
// But this is a moving target and should be ready for change.
// Ref. https://w3c-ccg.github.io/did-method-key/#rsa-repr
const VerificationMethodType = "Multikey"

// VerificationMethod defines the structure of a Verification Method
type VerificationMethod struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	PublicKeyMultibase string `json:"publicKeyMultibase"`
}

func New(
	id string,
	name string,
	publicKeyMultibase string) (VerificationMethod, error) {

	if !internal.IsAlnum(id) {
		return VerificationMethod{}, internal.ErrInvalidID
	}

	if !internal.IsValidNanoID(name) {
		return VerificationMethod{}, internal.ErrInvalidFragment
	}

	return VerificationMethod{
		ID:                 id + "#" + name,
		Type:               VerificationMethodType,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}

func (vm VerificationMethod) GetID() string {
	return vm.ID
}

func (vm VerificationMethod) Fragment() string {
	return internal.GetFragmentFromDID(vm.ID)
}
