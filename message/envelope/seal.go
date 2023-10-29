package envelope

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

func Seal(m *message.Message) (*Envelope, error) {
	msg, err := m.Pack()
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("message_encrypt: error packing message: %s\n", err))
	}

	to, err := doc.Fetch(m.To)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("message_encrypt: error fetching recipient document: %s\n", err))
	}

	recipientPublicKey, err := to.EncryptionKey()
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error getting recipient public key: %w", err)
	}

	// 1. Generate ephemeral x25519 key pair
	var ephemeralPrivate [32]byte
	_, err = rand.Read(ephemeralPrivate[:])
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error generating ephemeral private key: %w", err)
	}
	ephemeralPublic, err := curve25519.X25519(ephemeralPrivate[:], curve25519.Basepoint)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error deriving ephemeral public key: %w", err)
	}

	recipientPublicKeyBytes := recipientPublicKey.([]byte)
	// 2. Derive shared secret
	shared, err := curve25519.X25519(ephemeralPrivate[:], recipientPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error deriving shared secret: %w", err)
	}

	symmetricKey := internal.GenerateSymmetricKey(shared, ma.BLAKE3_SUM_SIZE)

	// Encrypt the actual message with ChaCha20-Poly1305
	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error creating AEAD: %w", err)
	}

	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error generating nonce: %w", err)
	}

	cipherText := aead.Seal(nil, nonce, []byte(msg), nil)
	cipherTextWithNonce := append(nonce, cipherText...)

	encodedCipherTextWithNonce, err := internal.MultibaseEncode(cipherTextWithNonce)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("message_encrypt: error encoding cipher text: %s", err))
	}

	// Send the ephemeral public key
	encodedEphemeralPubkey, err := internal.MultibaseEncode(ephemeralPublic)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error encoding ephemeral public key: %w", err)
	}

	return New(encodedCipherTextWithNonce, encodedEphemeralPubkey)
}
