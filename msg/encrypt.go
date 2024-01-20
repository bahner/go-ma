package msg

import (
	"crypto/rand"
	"fmt"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/chacha20poly1305"
)

// Encrypts a message fields from an message
func encrypt(data []byte, symmetricKey []byte, recipientPublicKeyBytes []byte) ([]byte, error) {

	// Encrypt the actual message with ChaCha20-Poly1305
	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("msg_encrypt: error creating AEAD: %w", err)
	}
	log.Debugf("msg_encrypt: aead: %x", aead)

	// Create a random nonce to make the encryption probabilistically unique
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("msg_encrypt: error generating nonce: %w", err)
	}
	log.Debugf("msg_encrypt: nonce: %x", nonce)
	log.Debugf("msg_encrypt: nonce size: %x", len(nonce))

	// Seal the generated cipher text
	cipherText := aead.Seal(nil, nonce, []byte(data), nil)
	log.Debugf("msg_encrypt: cipherText: %x", cipherText)
	cipherTextWithNonce := append(nonce, cipherText...)
	log.Debugf("msg_encrypt: cipherTextWithNonce: %x", cipherTextWithNonce)

	return cipherTextWithNonce, nil
}
