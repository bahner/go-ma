package msg

import (
	"crypto/ed25519"
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
	log "github.com/sirupsen/logrus"
)

func (m *Message) Sign(privKey *ed25519.PrivateKey) error {

	// Sign requires key to be of correct size
	if len(*privKey) != ed25519.PrivateKeySize {
		return fmt.Errorf("message/sign: invalid key size %d. Expected %d", len(*privKey), ed25519.PrivateKeySize)
	}

	bytes_to_sign, err := m.MarshalUnsignedToCBOR()
	if err != nil {
		return err
	}

	log.Debugf("Signed payload with hash: %s", m.MultibaseEncodedPayloadHash())
	// sig, err := privKey.Sign(rand.Reader, bytes_to_sign, nil)
	sig := ed25519.Sign(*privKey, bytes_to_sign)

	encoded_sig, err := internal.MultibaseEncode(sig)
	if err != nil {
		return fmt.Errorf("failed to encode signature: %w", err)
	}

	log.Debugf("Signed payload with signature: %s", encoded_sig)

	// This is the one place where we actually mutate the Message signature
	m.Signature = encoded_sig

	return nil
}

// VerifySignature verifies the Message's signature
// Returns nil if the signature is valid
func (m *Message) VerifySignature() error {

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

	payload, err := m.MarshalUnsignedToCBOR()
	if err != nil {
		return fmt.Errorf("message/verify: failed to pack payload: %w", err)
	}

	if len(signingKey) != ed25519.PublicKeySize {
		return fmt.Errorf("message/verify: invalid key size %d. Expected %d", len(signingKey), ed25519.PublicKeySize)
	}
	verification := ed25519.Verify(signingKey, payload, m.SignatureBytes())
	if !verification {
		return fmt.Errorf("message/verify: failed to verify signature")
	}

	return nil
}

// SignatureBytes returns the signature bytes
// It doesn't return an error, but logs it instead
func (m *Message) SignatureBytes() []byte {

	sigBytes, err := internal.MultibaseDecode(m.Signature)
	if err != nil {
		log.Errorf("failed to decode messageSignature: %s", err)

	}
	return sigBytes

}
