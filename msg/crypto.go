package msg

import (
	"crypto/rand"
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

	symmetricKey := key.GenerateSymmetricKey(shared, ma.BLAKE3_SUM_SIZE, []byte(ma.BLAKE3_LABEL))
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

// Encrypts a message fields from an message
// func encrypt(data []byte, symmetricKey []byte, recipientPublicKeyBytes []byte) ([]byte, error) {
func encrypt(data []byte, symmetricKey []byte) ([]byte, error) {

	// Encrypt the actual message with ChaCha20-Poly1305
	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("msg_encrypt: error creating AEAD: %w", err)
	}
	// log.Debugf("msg_encrypt: aead: %x", aead)

	// Create a random nonce to make the encryption probabilistically unique
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("msg_encrypt: error generating nonce: %w", err)
	}
	// log.Debugf("msg_encrypt: nonce: %x", nonce)
	// log.Debugf("msg_encrypt: nonce size: %x", len(nonce))

	// Seal the generated cipher text
	cipherText := aead.Seal(nil, nonce, []byte(data), nil)
	// log.Debugf("msg_encrypt: cipherText: %x", cipherText)
	cipherTextWithNonce := append(nonce, cipherText...)
	// log.Debugf("msg_encrypt: cipherTextWithNonce: %x", cipherTextWithNonce)

	return cipherTextWithNonce, nil
}

func generateEphemeralKeys(recipientPublicKeyBytes []byte) ([]byte, []byte, error) {

	// The private key is not stored, only used twice, both for the headers and the content encryption.
	// This should be OK, but we could use a different key for the content encryption in the future, if deemed necessary.

	// NB! This function should only ever be called once per message, otherwise the same ephemeral key will be used for both the headers and the content encryption.
	var ephemeralPrivate [curve25519.ScalarSize]byte
	_, err := rand.Read(ephemeralPrivate[:])
	if err != nil {
		return nil, nil, fmt.Errorf("msg_encrypt: error generating ephemeral private key: %w", err)
	}

	ephemeralPublic, err := curve25519.X25519(ephemeralPrivate[:], curve25519.Basepoint)
	if err != nil {
		return nil, nil, fmt.Errorf("msg_encrypt: error deriving ephemeral public key: %w", err)
	}
	log.Debugf("msg_enclose: ephemeralPublic: %x", ephemeralPublic)

	// Derive shared secret
	shared, err := curve25519.X25519(ephemeralPrivate[:], recipientPublicKeyBytes)
	if err != nil {
		return nil, nil, fmt.Errorf("msg_encrypt: error deriving shared secret: %w", err)
	}
	// log.Debugf("msg_encrypt: shared: %x", shared)

	// Generate a symmetric key from the shared secret using blake3
	symmetricKey := key.GenerateSymmetricKey(shared, ma.BLAKE3_SUM_SIZE, []byte(ma.BLAKE3_LABEL))
	// log.Debugf("msg_encrypt: symmetricKey: %x", symmetricKey)

	return ephemeralPublic, symmetricKey, nil

}
