package doc

import (
	"crypto/ed25519"
	"fmt"

	"github.com/multiformats/go-multibase"
)

func (d *Document) Verify() error {

	hashed, err := d.MulticodecHashedPayload()
	if err != nil {
		return fmt.Errorf("doc/verify: Error hashing payload: %s", err)
	}

	// Decode the multibase-encoded signature
	_, signature, err := multibase.Decode(d.Proof.ProofValue)
	if err != nil {
		return fmt.Errorf("doc/verify: Error decoding signature: %s", err)
	}

	pubKey, err := d.KeyAgreement.PublicKeyMultibase.Decode()
	if err != nil {
		return fmt.Errorf("doc/verify: Error getting signing key: %s", err)
	}

	// Verify the signature
	if !ed25519.Verify(pubKey.(ed25519.PublicKey), hashed[:], signature) {
		return fmt.Errorf("verification failed")
	}

	if err == nil {
		return nil
	}

	return fmt.Errorf("doc verify: Verification failed for all keys")
}
