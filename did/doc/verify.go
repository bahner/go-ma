package doc

import (
	"crypto/rsa"
	"fmt"

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

	rsaKeys, err := doc.VerificationMethodRSAPublicKeys()
	if err != nil {
		return fmt.Errorf("doc verify: Error getting RSA public keys: %s", err)
	}
	for _, rsaKey := range rsaKeys {

		err = rsa.VerifyPKCS1v15(rsaKey, SIGNATURE_HASH, hashed, signature)
		if err == nil { // If the key verifies successfully, return nil immediately
			return nil
		}
	}

	return fmt.Errorf("doc verify: Verification failed for all keys")
}
