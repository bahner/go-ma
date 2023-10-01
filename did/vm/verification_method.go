package vm

import (
	"github.com/bahner/go-ma/did/pkm"
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
	fragment string,
	publicKeyMultibase *pkm.PublicKeyMultibase) (VerificationMethod, error) {

	if !internal.IsAlnum(id) {
		return VerificationMethod{}, internal.ErrInvalidID
	}

	if !internal.IsValidFragment(fragment) {
		return VerificationMethod{}, internal.ErrInvalidFragment
	}

	return VerificationMethod{
		ID:                 id + fragment, // A DID.String() and a "#fragment(Of Some Sort)"
		Type:               VerificationMethodType,
		PublicKeyMultibase: publicKeyMultibase.PublicKeyMultibase,
	}, nil
}
