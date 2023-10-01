package doc

import (
	"crypto"
	"crypto/ed25519"
	"fmt"

	"github.com/multiformats/go-multibase"
	"lukechampine.com/blake3"
)

func (doc *Document) Verify() error {
	p, err := doc.MarshalPayloadToJSON()
	if err != nil {
		return fmt.Errorf("doc verify: Error marshalling payload to JSON: %s", err)
	}

	hashed := blake3.Sum256(p)

	// Decode the multibase-encoded signature
	_, signature, err := multibase.Decode(doc.Signature)
	if err != nil {
		return fmt.Errorf("doc/verify: Error decoding signature: %s", err)
	}

	pubKey, err := doc.SigningKey()
	if err != nil {
		return fmt.Errorf("doc/verify: Error getting signing key: %s", err)
	}

	// Verify the signature
	err = verifyData(pubKey, hashed[:], signature)
	if err == nil {
		return nil
	}

	return fmt.Errorf("doc verify: Verification failed for all keys")
}

// VerifyData verifies a signature against a message and public key.
func verifyData(publicKey crypto.PublicKey, message, signature []byte) error {
	// Type assertion
	switch pk := publicKey.(type) {
	case ed25519.PublicKey:
		if !ed25519.Verify(pk, message, signature) {
			return fmt.Errorf("verification failed")
		}
	default:
		return fmt.Errorf("unsupported public key type %T", pk)
	}

	return nil
}
