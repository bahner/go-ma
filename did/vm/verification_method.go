package vm

import (
	"fmt"

	"github.com/bahner/go-ma/did/pkm"
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
	publicKeyMultibase *pkm.PublicKeyMultibase) (VerificationMethod, error) {

	if !internal.IsAlnum(id) {
		return VerificationMethod{}, internal.ErrInvalidID
	}

	if !internal.IsValidFragment(fragment) {
		return VerificationMethod{}, internal.ErrInvalidFragment
	}

	// Check if we have a valid method for the given multicodec code
	// For given key
	vmType, err := getVerificationMethodForKeyType(publicKeyMultibase.MulticodecCodeString)
	if err != nil {
		return VerificationMethod{}, fmt.Errorf("no verification matches key type: %s", publicKeyMultibase.MulticodecCodeString)
	}

	return VerificationMethod{
		ID:                 id + fragment, // A DID.String() and a "#fragment(Of Some Sort)"
		Type:               vmType,
		PublicKeyMultibase: publicKeyMultibase.PublicKeyMultibase,
	}, nil
}

func VerificationMethodTypeFromPKM(pkmb *pkm.PublicKeyMultibase) string {
	// Convert string to multicodec.Code
	return VerificationPubkeyMethodMapping[pkmb.MulticodecCodeString]
}
