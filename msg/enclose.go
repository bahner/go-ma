package msg

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/key"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/curve25519"
)

// Use a pointer here, this might be arbitrarily big.
func (m *Message) Enclose() (*Envelope, error) {

	// AT this point we *need* to fetch the recipient's document, otherwise we can't encrypt the message.
	// But this fetch should probably have a timeout, so we don't get stuck here - or a caching function.
	to, err := doc.FetchFromDID(m.Headers.To)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: error fetching recipient document: %s", err)
	}

	recipientPublicKeyBytes, err := to.KeyAgreementPublicKeyBytes()
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: error getting recipient public key: %w", err)
	}

	// Generate ephemeral keys to be used for his message
	ephemeralPublic, symmetricKey, err := generateEphemeralKeys(recipientPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: failed to generate ephemeral keys: %w", err)
	}

	msgHeaders, err := m.Headers.MarshalToCBOR()
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: error marshalling headers: %w", err)
	}

	encryptedMsgHeaders, err := encrypt(msgHeaders, symmetricKey, recipientPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: error encrypting headers: %w", err)
	}

	encryptedContent, err := encrypt(m.Body.Content, symmetricKey, recipientPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: error encrypting content: %w", err)
	}

	return &Envelope{
		EphemeralKey: ephemeralPublic,
		Headers:      encryptedMsgHeaders,
		Content:      encryptedContent,
	}, nil
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
	log.Debugf("msg_encrypt: shared: %x", shared)

	// Generate a symmetric key from the shared secret using blake3
	symmetricKey := key.GenerateSymmetricKey(shared, ma.BLAKE3_SUM_SIZE)
	log.Debugf("msg_encrypt: symmetricKey: %x", symmetricKey)

	return ephemeralPublic, symmetricKey, nil

}
