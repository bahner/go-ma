package msg

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/msg/envelope"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

func (m *Message) Enclose() (*envelope.Envelope, error) {
	// First check the stuff we don't have control over.
	// Fail fast.
	msg, err := m.Pack()
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error packing message: %s", err)
	}

	to, err := doc.FetchFromDID(m.To)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error fetching recipient document: %s", err)
	}

	recipientPublicKeyBytes, err := to.KeyAgreementPublicKeyBytes()
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error getting recipient public key: %w", err)
	}

	// Generate an ephemeral key pair
	var ephemeralPrivate [curve25519.ScalarSize]byte
	_, err = rand.Read(ephemeralPrivate[:])
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error generating ephemeral private key: %w", err)
	}
	ephemeralPublic, err := curve25519.X25519(ephemeralPrivate[:], curve25519.Basepoint)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error deriving ephemeral public key: %w", err)
	}
	log.Debugf("message_encrypt: ephemeralPublic: %x", ephemeralPublic)

	// Derive shared secret
	shared, err := curve25519.X25519(ephemeralPrivate[:], recipientPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error deriving shared secret: %w", err)
	}
	log.Debugf("message_encrypt: shared: %x", shared)

	// Generate a symmetric key from the shared secret using blake3
	symmetricKey := key.GenerateSymmetricKey(shared, ma.BLAKE3_SUM_SIZE)
	log.Debugf("message_encrypt: symmetricKey: %x", symmetricKey)

	// Encrypt the actual message with ChaCha20-Poly1305
	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error creating AEAD: %w", err)
	}
	log.Debugf("message_encrypt: aead: %x", aead)

	// Create a random nonce to make the encryption probabilistically unique
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error generating nonce: %w", err)
	}
	log.Debugf("message_encrypt: nonce: %x", nonce)
	log.Debugf("message_encrypt: nonce size: %x", len(nonce))

	// Seal the generated cipher text
	cipherText := aead.Seal(nil, nonce, []byte(msg), nil)
	log.Debugf("message_encrypt: cipherText: %x", cipherText)
	cipherTextWithNonce := append(nonce, cipherText...)
	log.Debugf("message_encrypt: cipherTextWithNonce: %x", cipherTextWithNonce)

	return envelope.New(cipherTextWithNonce, ephemeralPublic)
}
