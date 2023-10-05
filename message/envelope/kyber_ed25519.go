package envelope

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/suites"
	"golang.org/x/crypto/chacha20poly1305"
)

func kyberEd25519Encrypt(m *message.Message) (*Envelope, error) {
	msg, err := m.Pack()
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("kyber_ed25519: error packing message: %s\n", err))
	}

	to, err := doc.Fetch(m.To)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("kyber_ed25519: error fetching recipient document: %s\n", err))
	}

	recipientPublicKey, err := to.EncryptionKey()
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error getting recipient public key: %w", err)
	}

	suite := suites.MustFind(ma.KYBER_SUITE)
	ephemeralPrivate := suite.Scalar().Pick(suite.RandomStream())
	ephemeralPublic := suite.Point().Mul(ephemeralPrivate, nil)

	// Convert recipientPublicKey (which is a []byte) to a Kyber Point
	recipientPublicKeyPoint := suite.Point()
	err = recipientPublicKeyPoint.UnmarshalBinary(recipientPublicKey.([]byte))
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error unmarshalling public key: %w", err)
	}

	sharedPoint := suite.Point().Mul(ephemeralPrivate, recipientPublicKeyPoint)

	// Convert sharedPoint to a byte slice
	sharedBytes, err := sharedPoint.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error marshalling shared point: %w", err)
	}

	symmetricKey := generateSymmetricKey(sharedBytes, 32)

	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error creating AEAD: %w", err)
	}

	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error generating nonce: %w", err)
	}

	cipherText := aead.Seal(nil, nonce, []byte(msg), nil)
	cipherTextWithNonce := append(nonce, cipherText...)
	multicodecedCipherTextWithNonce, err := internal.MulticodecEncode("ECDHKyberEd25519ChaCha20Poly1305BLAKE3", cipherTextWithNonce)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("kyber_ed25519: error multicodec encoding cipher text: %s", err))
	}

	encodedMulticodecedCipherTextWithNonce, err := internal.MultibaseEncode(multicodecedCipherTextWithNonce)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("kyber_ed25519: error encoding cipher text: %s", err))
	}

	ephemeralPublicBytes, err := ephemeralPublic.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error marshalling ephemeral public key: %w", err)
	}
	encodedEphemeralPubkey, err := internal.MultibaseEncode(ephemeralPublicBytes)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error encoding ephemeral public key: %w", err)
	}

	return New(encodedMulticodecedCipherTextWithNonce, encodedEphemeralPubkey)
}

func kyberEd25519Decrypt(envelope *Envelope, recipientPrivateKey interface{}) (*message.Message, error) {
	suite := suites.MustFind(ma.KYBER_SUITE)

	// Decode the ephemeral public key
	ephemeralPublicBytes, err := internal.MultibaseDecode(envelope.EphemeralKey)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error decoding ephemeral public key: %w", err)
	}

	ephemeralPublicKey := suite.Point()
	err = ephemeralPublicKey.UnmarshalBinary(ephemeralPublicBytes)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error unmarshalling ephemeral public key: %w", err)
	}

	// Compute the shared point using the recipient's private key and the ephemeral public key
	sharedPoint := suite.Point().Mul(recipientPrivateKey.(kyber.Scalar), ephemeralPublicKey)

	// Convert sharedPoint to a byte slice
	sharedBytes, err := sharedPoint.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error marshalling shared point: %w", err)
	}

	symmetricKey := generateSymmetricKey(sharedBytes, 32)

	// Decode the ciphertext
	cipherTextWithNonce, err := internal.MultibaseDecode(envelope.Message)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error decoding cipher text: %w", err)
	}

	// Separate nonce and cipher text
	nonce := cipherTextWithNonce[:chacha20poly1305.NonceSizeX]
	cipherText := cipherTextWithNonce[chacha20poly1305.NonceSizeX:]

	aead, err := chacha20poly1305.NewX(symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error creating AEAD: %w", err)
	}

	plainText, err := aead.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: decryption failed: %w", err)
	}

	m, err := message.Unpack(string(plainText))
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error unpacking message: %w", err)
	}

	return m, nil
}
