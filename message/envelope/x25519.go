package envelope

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

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

func decryptX25519Message(envelope *Envelope, privKey []byte) (*message.Message, error) {
	// Decode the ephemeral public key from the envelope
	encodedEphemeralPublickeyBytes, err := internal.MultibaseDecode(envelope.EphemeralKey)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error multibase decoding ephemeral public key: %w", err)
	}

	encodedEphemeralPublicKey := string(encodedEphemeralPublickeyBytes)
	ephemeralKey, err := internal.MultibaseDecode(encodedEphemeralPublicKey)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error multibase decoding ephemeral public key: %w", err)
	}

	// Derive the shared secret using recipient's private key and ephemeral public key
	shared, err := curve25519.X25519(privKey, ephemeralKey)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error deriving shared secret: %w", err)
	}

	symmetricKey := generateSymmetricKey(shared, 32)

	// Decode the encrypted message from the envelope
	cipherTextWithNonce, err := internal.MultibaseDecode(envelope.Message)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error multibase decoding cipher text: %w", err)
	}

	// Split the nonce from the ciphertext
	if len(cipherTextWithNonce) < chacha20poly1305.NonceSizeX {
		return nil, fmt.Errorf("message_decrypt: ciphertext too short")
	}
	nonce, cipherText := cipherTextWithNonce[:chacha20poly1305.NonceSizeX], cipherTextWithNonce[chacha20poly1305.NonceSizeX:]

	// Decrypt the message with ChaCha20-Poly1305
	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, err
	}

	msg, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	// Unpack the decrypted message
	unpackedMsg, err := message.Parse(string(msg))
	if err != nil {
		return nil, err
	}

	return unpackedMsg, nil
}
