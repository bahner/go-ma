package envelope

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

func Unseal(e *Envelope, privKey []byte) (*message.Message, error) {
	ephemeralKey, err := internal.MultibaseDecode(e.Seal)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error multibase decoding ephemeral public key: %w", err)
	}

	// Derive the shared secret using recipient's private key and ephemeral public key
	shared, err := curve25519.X25519(privKey, ephemeralKey)
	if err != nil {
		return nil, fmt.Errorf("message_decrypt: error deriving shared secret: %w", err)
	}

	symmetricKey := internal.GenerateSymmetricKey(shared, ma.BLAKE3_SUM_SIZE)

	// Decode the encrypted message from the envelope
	cipherTextWithNonce, err := internal.MultibaseDecode(e.Message)
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
