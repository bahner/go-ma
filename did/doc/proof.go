package doc

import (
	"fmt"

	"github.com/bahner/go-ma/key"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/multiformats/go-multibase"
)

const (
	proofType      = "MultiformatSignature2023"
	proofPurpose   = "assertionMethod"
	proofNumFields = 4
)

type Proof struct {
	Type               string `cbor:"type" json:"type"`
	VerificationMethod string `cbor:"verificationMethod" json:"verificationMethod"`
	ProofPurpose       string `cbor:"proofPurpose" json:"proofPurpose"`
	ProofValue         string `cbor:"proofValue" json:"proofValue"`
}

func (d *Document) Sign(signKey key.SigningKey, vm VerificationMethod) error {

	if signKey.PublicKeyMultibase != vm.PublicKeyMultibase {
		return fmt.Errorf("doc sign: signKey.PublicKeyMultibase != vm.PublicKeyMultibase. %w", ErrPublicKeyMultibaseMismatch)
	}

	multicodecHashed, err := d.PayloadHash()
	if err != nil {
		return fmt.Errorf("doc sign: Error hashing payload: %s", err)
	}

	// Sign the hashed payload with an ed25519 key
	signature, err := signKey.Sign(multicodecHashed)
	if err != nil {
		return fmt.Errorf("doc sign: %w", err)
	}

	// Multibase encode the signed data for public consumption
	proofValue, err := multibase.Encode(multibase.Base58BTC, signature)
	if err != nil {
		return fmt.Errorf("doc sign: Error encoding signature: %w", err)
	}

	d.Proof = NewProof(proofValue, vm.ID)

	return nil
}

func NewProof(proofValue string, vm string) Proof {

	return Proof{
		Type:               proofType,
		ProofPurpose:       proofPurpose,
		ProofValue:         proofValue,
		VerificationMethod: vm,
	}
}

func buildProofNode(proof Proof) (ipld.Node, error) {
	nb := basicnode.Prototype.Map.NewBuilder()
	ma, err := nb.BeginMap(proofNumFields)
	if err != nil {
		return nil, err
	}

	ma.AssembleKey().AssignString("type")
	ma.AssembleValue().AssignString(proof.Type)

	ma.AssembleKey().AssignString("verificationMethod")
	ma.AssembleValue().AssignString(proof.VerificationMethod)

	ma.AssembleKey().AssignString("proofPurpose")
	ma.AssembleValue().AssignString(proof.ProofPurpose)

	ma.AssembleKey().AssignString("proofValue")
	ma.AssembleValue().AssignString(proof.ProofValue)

	ma.Finish()

	return nb.Build(), nil
}
