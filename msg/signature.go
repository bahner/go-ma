package msg

import (
	"crypto/ed25519"
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
)

func (m *Message) Sign(privKey *ed25519.PrivateKey) error {

	// Sign requires key to be of correct size
	if len(*privKey) != ed25519.PrivateKeySize {
		return fmt.Errorf("message/sign: invalid key size %d. Expected %d", len(*privKey), ed25519.PrivateKeySize)
	}

	data_to_sign, err := m.PayloadPack()
	if err != nil {
		return err
	}

	bytes_to_sign := []byte(data_to_sign)

	// sig, err := privKey.Sign(rand.Reader, bytes_to_sign, nil)
	sig := ed25519.Sign(*privKey, bytes_to_sign)

	encoded_sig, err := internal.MultibaseEncode(sig)
	if err != nil {
		return fmt.Errorf("failed to encode signature: %w", err)
	}

	m.Signature = encoded_sig

	return nil
}

// Verify verifies the Message's signature
// Returns nil if the signature is valid
func (m *Message) Verify() error {

	did, err := did.NewFromDID(m.From)
	if err != nil {
		return fmt.Errorf("message/verify: failed to create did from From: %w", err)
	}

	senderDoc, err := doc.Fetch(did.Identifier)
	if err != nil {
		return fmt.Errorf("message/verify: failed to fetch sender document")
	}

	signingKey, err := senderDoc.AssertionMethodPublicKey()
	if err != nil {
		return fmt.Errorf("message/verify: failed to get signing key: %w", err)
	}

	payload, err := m.PayloadPack()
	if err != nil {
		return fmt.Errorf("message/verify: failed to pack payload: %w", err)
	}

	verification := ed25519.Verify(signingKey, []byte(payload), []byte(m.Signature))
	if !verification {
		return fmt.Errorf("message/verify: failed to verify signature")
	}

	return nil
}
