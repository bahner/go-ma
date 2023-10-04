package envelope

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
	"lukechampine.com/blake3"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
)

func Enclose(m *message.Message) (*Envelope, error) {
	return x25519Encrypt(m)
}

func x25519Encrypt(m *message.Message) (*Envelope, error) {
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

	to.EncryptionKey()

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

	symmetricKey := generateSymmetricKey(shared, 32)

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
	multicodecedCipherTextWithNonce, err := internal.MulticodecEncode("ECDHX25519ChaCha20Poly1305BLAKE3", cipherTextWithNonce)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("message_encrypt: error multicodec encoding cipher text: %s", err))
	}

	encodedMulticodecedCipherTextWithNonce, err := internal.MultibaseEncode(multicodecedCipherTextWithNonce)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("message_encrypt: error encoding cipher text: %s", err))
	}

	// Send the ephemeral public key
	encodedEphemeralPubkey, err := internal.MultibaseEncode(ephemeralPublic)
	if err != nil {
		return nil, fmt.Errorf("message_encrypt: error encoding ephemeral public key: %w", err)
	}

	return New(encodedMulticodecedCipherTextWithNonce, encodedEphemeralPubkey)
}

func generateSymmetricKey(shared []byte, size int) []byte {

	// Hash the shared secret with Blake3 in a uniform way.

	// The label is the MIME Type, just so we have our own namespace.
	label := []byte(ma.MESSAGE_MIME_TYPE)
	hasher := blake3.New(size, nil)
	hasher.Write(label)
	hasher.Write(shared)
	return hasher.Sum(nil)[:size]
}
