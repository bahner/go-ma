package msg

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/key"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

// Decrypts a message fields from an message
func decrypt(data []byte, ephemeralKey []byte, privKey []byte) ([]byte, error) {

	// Derive the shared secret using recipient's private key and ephemeral public key
	shared, err := curve25519.X25519(privKey, ephemeralKey)
	if err != nil {
		return nil, fmt.Errorf("error deriving shared secret: %w", err)
	}
	log.Debugf("shared: %x", shared)

	symmetricKey := key.GenerateSymmetricKey(shared, ma.BLAKE3_SUM_SIZE)
	log.Debugf("symmetricKey: %x", symmetricKey)

	// Split the nonce from the ciphertext
	if len(data) < chacha20poly1305.NonceSizeX {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, cipherText := data[:chacha20poly1305.NonceSizeX], data[chacha20poly1305.NonceSizeX:]
	log.Debugf("nonce: %x", nonce)
	log.Debugf("cipherText: %x", cipherText)

	// Decrypt the message with ChaCha20-Poly1305
	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("error creating AEAD: %w", err)
	}
	log.Debugf("aead: %x", aead)

	m, err := aead.Open(nil, nonce, cipherText, nil)
	log.Debugf("msg: %x", m)
	if err != nil {
		return nil, fmt.Errorf("envelope_data: error decrypting message: %w", err)
	}

	return m, nil
}
