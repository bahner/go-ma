package doc

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	"github.com/multiformats/go-multibase"
)

const (
	proofType    = "MultiformatSignature2023"
	proofPurpose = "assertionMethod"
)

type Proof struct {
	Created            string `cbor:"created" json:"created"`
	Type               string `cbor:"type" json:"type"`
	VerificationMethod string `cbor:"verificationMethod" json:"verificationMethod"`
	ProofPurpose       string `cbor:"proofPurpose" json:"proofPurpose"`
	ProofValue         string `cbor:"proofValue" json:"proofValue"`
}

func (d *Document) Sign(signKey key.SigningKey, vm VerificationMethod) error {

	if signKey.PublicKeyMultibase != vm.PublicKeyMultibase {
		return fmt.Errorf("doc sign: signKey.PublicKeyMultibase != vm.PublicKeyMultibase")
	}

	multicodecHashed, err := d.PayloadHash()
	if err != nil {
		return fmt.Errorf("doc sign: Error hashing payload: %s", err)
	}

	// Sign the hashed payload with an ed25519 key
	signature, err := signKey.Sign(multicodecHashed)
	if err != nil {
		return fmt.Errorf("doc sign: Error signing payload: %s", err)
	}

	// Multibase encode the signed data for public consumption
	proofValue, err := multibase.Encode(multibase.Base58BTC, signature)
	if err != nil {
		return fmt.Errorf("doc sign: Error encoding signature: %s", err)
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
