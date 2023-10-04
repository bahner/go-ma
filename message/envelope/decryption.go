package envelope

import (
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
)

func (e *Envelope) Open(privKey []byte) (*message.Message, error) {
	return decryptX25519Message(e, privKey)
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
	cipherTextWithNonce, err := internal.MultibaseDecode(envelope.EncryptedMsg)
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
