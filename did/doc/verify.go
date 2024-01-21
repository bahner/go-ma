package doc

import (
	"crypto/ed25519"
	"fmt"

	"github.com/multiformats/go-multibase"
)

func (d *Document) Verify() error {

	hashed, err := d.PayloadHash()
	if err != nil {
		return fmt.Errorf("doc/verify: Error hashing payload: %s", err)
	}

	// Decode the multibase-encoded signature
	_, signature, err := multibase.Decode(d.Proof.ProofValue)
	if err != nil {
		return fmt.Errorf("doc/verify: Error decoding signature: %s", err)
	}

	pubKey, err := d.AssertionMethodPublicKey()
	if err != nil {
		return fmt.Errorf("doc/verify: Error getting signing key: %s", err)
	}

	// Verify the signature
	if !ed25519.Verify(pubKey, hashed[:], signature) {
		return fmt.Errorf("verification failed")
	}

	return nil
}

func (d *Document) isValid() bool {

	return d.Verify() == nil
}
