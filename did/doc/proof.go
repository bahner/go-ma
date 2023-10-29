package doc

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	"github.com/multiformats/go-multibase"
)

const (
	proofType    = "Ed25519Signature2020"
	proofPurpose = "assertionMethod"
)

type Proof struct {
	Created            string `json:"created"`
	Type               string `json:"type"`
	VerificationMethod string `json:"verificationMethod"`
	ProofPurpose       string `json:"proofPurpose"`
	ProofValue         string `json:"proofValue"`
}

// "type": "Ed25519Signature2020",
// "created": "2023-10-26T12:00:00Z",
// "verificationMethod": "did:key:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#signature",
// "signatureValue": "Base64-encoded-signature-value"

func (d *Document) Sign(signKey key.SignatureKey, vm VerificationMethod) error {
	p, err := d.MarshalPayloadToJSON()
	if err != nil {
		return err
	}

	// Sign the payload with an ed25519 key
	signature, err := signKey.Sign(p)
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error signing payload: %s\n", err))
	}

	proofValue, err := multibase.Encode(ma.MULTIBASE_ENCODING, signature)
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error encoding signature: %s\n", err))
	}

	return nil
}

func NewProof(signKey key.SignatureKey, vm VerificationMethod) (Proof, error) {

	proof := Proof{
		Type:               proofType,
		ProofPurpose:       proofPurpose,
		VerificationMethod: signKey.VerificationMethod,
	}

	return proof, nil
}
