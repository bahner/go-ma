package msg

import (
	"crypto/ed25519"
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/doc"
	log "github.com/sirupsen/logrus"
)

func (m *Message) Sign(privKey *ed25519.PrivateKey) error {

	// Sign requires key to be of correct size
	if len(*privKey) != ed25519.PrivateKeySize {
		return fmt.Errorf("message/sign: invalid key size %d. Expected %d", len(*privKey), ed25519.PrivateKeySize)
	}

	bytes_to_sign, err := m.marshalUnsignedHeadersToCBOR()
	if err != nil {
		return err
	}

	sig := ed25519.Sign(*privKey, bytes_to_sign)

	log.Debugf("Signed payload with signature: %s", sig)

	// This is the one place where we actually mutate the Message signature
	m.Signature = sig

	return nil
}

// Verify verifies the Message's signature
// Returns nil if the signature is valid
func (m *Message) Verify() error {

	if m == nil {
		return fmt.Errorf("message/verify: nil message")
	}

	if m.From == "" {
		return fmt.Errorf("message/verify: missing From")
	}

	if m.Signature == nil {
		return fmt.Errorf("message/verify: missing signature")
	}

	// Sender document
	did, err := did.New(m.From)
	if err != nil {
		return fmt.Errorf("message/verify: failed to create did from From: %w", err)
	}

	senderDoc, err := doc.Fetch(did.String(), true) // Accept cached document
	if err != nil {
		return fmt.Errorf("message/verify: failed to fetch sender document")
	}

	// Signing key
	signingKey, err := senderDoc.AssertionMethodPublicKey()
	if err != nil {
		return fmt.Errorf("message/verify: failed to get signing key: %w", err)
	}

	if len(signingKey) != ed25519.PublicKeySize {
		return fmt.Errorf("message/verify: invalid key size %d. Expected %d", len(signingKey), ed25519.PublicKeySize)
	}

	// Payload
	payload, err := m.marshalUnsignedHeadersToCBOR()
	if err != nil {
		return fmt.Errorf("message/verify: failed to pack payload: %w", err)
	}

	// Verification
	verification := ed25519.Verify(signingKey, payload, m.Signature)
	if !verification {
		return fmt.Errorf("message/verify: failed to verify signature")
	}

	return nil
}
