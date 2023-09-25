package vm

import (
	"github.com/bahner/go-ma/did/pubkey"
	"github.com/bahner/go-ma/internal"
)

// VerificationMethod defines the structure of a Verification Method
type VerificationMethod struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	PublicKeyMultibase string `json:"publicKeyMultibase"`
}

func New(
	id string,
	fragment string,
	publicKeyMultibase pubkey.PublicKeyMultibase) (VerificationMethod, error) {

	if !internal.IsAlnum(id) {
		return VerificationMethod{}, internal.ErrInvalidID
	}

	if !internal.IsValidFragment(fragment) {
		return VerificationMethod{}, internal.ErrInvalidFragment
	}

	return VerificationMethod{
		ID:                 id + fragment, // A DID.String() and a "#fragment(Of Some Sort)"
		Type:               publicKeyMultibase.Type,
		PublicKeyMultibase: publicKeyMultibase.PublicKeyMultibase,
	}, nil
}
