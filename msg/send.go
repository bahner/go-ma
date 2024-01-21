package msg

import (
	"context"
	"fmt"

	"github.com/bahner/go-ma/did/doc"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// Send message to a pubsub topic
// We could have done p2p here, but it's much cleaner to
// take a topic as an argument. The topic is just a medium
// for sending messages, not actually the recipient necessarily.
func (m *Message) Send(ctx context.Context, t *pubsub.Topic) error {

	// Create a new envelope
	e, err := m.enclose()
	if err != nil {
		return fmt.Errorf("send: envelope creation error: %w", err)
	}

	eBytes, err := e.marshalToCBOR()
	if err != nil {
		return fmt.Errorf("send: envelope serialization error: %w", err)
	}

	t.Publish(ctx, eBytes, nil)

	return nil
}

// Use a pointer here, this might be arbitrarily big.
func (m *Message) enclose() (*Envelope, error) {

	// AT this point we *need* to fetch the recipient's document, otherwise we can't encrypt the message.
	// But this fetch should probably have a timeout, so we don't get stuck here - or a caching function.
	to, err := doc.GetOrFetch(m.To)
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

	msgHeaders, err := m.marshalHeadersToCBOR()
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: error marshalling headers: %w", err)
	}

	encryptedMsgHeaders, err := encrypt(msgHeaders, symmetricKey, recipientPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: error encrypting headers: %w", err)
	}

	encryptedContent, err := encrypt(m.Content, symmetricKey, recipientPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: error encrypting content: %w", err)
	}

	return &Envelope{
		EphemeralKey:     ephemeralPublic,
		EncryptedHeaders: encryptedMsgHeaders,
		EncryptedContent: encryptedContent,
	}, nil
}
