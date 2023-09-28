package message

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message/envelope"
	"github.com/multiformats/go-multibase"
)

var MESSAGE_HASHER = crypto.BLAKE2b_256.New()

// Takes a Pointer to a RSA public key and a label string.
// The label is a hint to the reciever that they should look for
// a key with this fragment in the key name.
func (m *Message) Encrypt(to_rsa_pubkey *rsa.PublicKey, label string) (*envelope.Envelope, error) {
	msg, err := m.Pack()
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("message_encrypt: error packing message: %s", err))
	}

	// Generate a random symmetric (AES) key
	symmetricKey := make([]byte, 32) // assuming AES-256
	_, err = rand.Read(symmetricKey)
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("message_encrypt: error generating symmetric key: %s", err))
	}

	// Encrypt the actual message with the symmetric key
	block, err := aes.NewCipher(symmetricKey)
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("message_encrypt: error creating cipher: %s", err))
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("message_encrypt: error creating GCM: %s", err))
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("message_encrypt: error generating nonce: %s", err))
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(msg), nil)
	encodedCipherText, err := MessageEncoder(cipherText)
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("message_encrypt: error encoding cipher text: %s", err))
	}

	// Encrypt the symmetric key with the RSA public key
	encryptedSymmetricKey, err := rsa.EncryptOAEP(MESSAGE_HASHER, rand.Reader, to_rsa_pubkey, symmetricKey, []byte(label))
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("message_encrypt: error encrypting symmetric key: %s", err))
	}
	encodedEncryptedSymKey, err := MessageEncoder(encryptedSymmetricKey)
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("message_encrypt: error encoding encrypted symmetric key: %s", err))
	}

	return envelope.New(encodedCipherText, encodedEncryptedSymKey)
}

func Open(envelope *envelope.Envelope, privKey *rsa.PrivateKey) (*Message, error) {
	// Decode and Decrypt the symmetric key using the RSA private key
	_, encryptedSymmetricKey, err := multibase.Decode(envelope.EncryptedKey)
	if err != nil {
		return nil, err
	}
	symmetricKey, err := rsa.DecryptOAEP(
		MESSAGE_HASHER,
		nil,
		privKey,
		encryptedSymmetricKey,
		[]byte(MESSAGE_ENCRYPTION_LABEL))
	if err != nil {
		return nil, err
	}

	// Decode and Decrypt the actual message with the symmetric key
	_, cipherText, err := multibase.Decode(envelope.EncryptedMsg)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(symmetricKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	msg, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	// Unpack the decrypted message
	unpackedMsg, err := Parse(string(msg))
	if err != nil {
		return nil, err
	}

	return unpackedMsg, nil
}
