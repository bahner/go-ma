package doc

import (
	"crypto/rsa"
	"fmt"

	"github.com/bahner/go-ma/did/pubkey"
	"github.com/multiformats/go-multibase"
)

func (doc *Document) Verify() error {
	p, err := doc.MarshalPayloadToJSON()
	if err != nil {
		return fmt.Errorf("doc verify: Error marshalling payload to JSON: %s", err)
	}

	// Compute the hash of the payload
	h := SIGNATURE_HASH.New()
	h.Write(p)
	hashed := h.Sum(nil)

	// Decode the multibase-encoded signature
	_, signature, err := multibase.Decode(doc.Signature)
	if err != nil {
		return fmt.Errorf("doc verify: Error decoding signature: %s", err)
	}

	// Loop through the VerificationMethods and try to verify with each RSA public key found
	for _, method := range doc.VerificationMethod {
		if method.Type == "RsaVerificationKey2018" {
			pubKey, err := pubkey.Decode(method.PublicKeyMultibase) // Directly get the *rsa.PublicKey
			if err != nil {
				return fmt.Errorf("doc verify: Error decoding and parsing public key: %s", err)
			}

			err = rsa.VerifyPKCS1v15(pubKey, SIGNATURE_HASH, hashed, signature)
			if err == nil { // If the key verifies successfully, return nil immediately
				return nil
			}
		}
	}

	return fmt.Errorf("doc verify: Verification failed for all keys")
}
