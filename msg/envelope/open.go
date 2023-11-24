package envelope

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/msg"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

// Open takes an envelope and a private key and returns the decrypted message.
func (e *Envelope) Open(privKey [32]byte) (*msg.Message, error) {

	// Convert the private key to a byte slice
	privKeyBytes := privKey[:]

	// Derive the shared secret using recipient's private key and ephemeral public key
	shared, err := curve25519.X25519(privKeyBytes, e.EphemeralKey)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error deriving shared secret: %w", err)
	}
	log.Debugf("message_decrypt: shared: %x", shared)

	symmetricKey := key.GenerateSymmetricKey(shared, ma.BLAKE3_SUM_SIZE)
	log.Debugf("message_decrypt: symmetricKey: %x", symmetricKey)

	// Split the nonce from the ciphertext
	if len(e.Message) < chacha20poly1305.NonceSizeX {
		return nil, fmt.Errorf("message_decrypt: ciphertext too short")
	}
	nonce, cipherText := e.Message[:chacha20poly1305.NonceSizeX], e.Message[chacha20poly1305.NonceSizeX:]
	log.Debugf("message_decrypt: nonce: %x", nonce)
	log.Debugf("message_decrypt: cipherText: %x", cipherText)

	// Decrypt the message with ChaCha20-Poly1305
	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error creating AEAD: %w", err)
	}
	log.Debugf("message_decrypt: aead: %x", aead)

	m, err := aead.Open(nil, nonce, cipherText, nil)
	log.Debugf("message_decrypt: msg: %x", m)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error decrypting message: %w", err)
	}

	// Unpack the decrypted message
	unpackedMsg, err := msg.Parse(string(m))
	if err != nil {
		return nil, err
	}

	return unpackedMsg, nil
}
