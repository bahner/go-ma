package doc

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	"github.com/multiformats/go-multibase"
)

const (
	proofType    = "MultiformatSignature2023"
	proofPurpose = "assertionMethod"
)

type Proof struct {
	_                  struct{} `cbor:",toarray"`
	Created            string   `cbor:"created"`
	Type               string   `cbor:"type"`
	VerificationMethod string   `cbor:"verificationMethod"`
	ProofPurpose       string   `cbor:"proofPurpose"`
	ProofValue         string   `cbor:"proofValue"`
}

func (d *Document) Sign(signKey key.SigningKey, vm VerificationMethod) error {

	if signKey.PublicKeyMultibase != vm.PublicKeyMultibase {
		return fmt.Errorf("doc sign: signKey.PublicKeyMultibase != vm.PublicKeyMultibase")
	}

	multicodecHashed, err := d.MulticodecHashedPayload()
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error hashing payload: %s\n", err))
	}

	// Sign the hashed payload with an ed25519 key
	signature, err := signKey.Sign(multicodecHashed)
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error signing payload: %s\n", err))
	}

	// Multibase encode the signed data for public consumption
	proofValue, err := multibase.Encode(ma.MULTIBASE_ENCODING, signature)
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error encoding signature: %s\n", err))
	}

	d.Proof = NewProof(proofValue, vm.ID)

	return nil
}

func NewProof(proofValue string, vm string) Proof {

	created := internal.NowIsoString()

	return Proof{
		Created:            created,
		Type:               proofType,
		ProofPurpose:       proofPurpose,
		ProofValue:         proofValue,
		VerificationMethod: vm,
	}
}
