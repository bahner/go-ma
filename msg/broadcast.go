package msg

import (
	"context"
	"crypto/ed25519"
	"fmt"

	"github.com/bahner/go-ma"
	cbor "github.com/fxamacker/cbor/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	nanoid "github.com/matoous/go-nanoid/v2"
)

// A Broadcast  message is just a message with no recipient.
// This packages tries to enforce this, so as to avoid mistakes.
func NewBroadcast(
	from string,
	content []byte,
	contentType string,
	priv_key ed25519.PrivateKey) (*Message, error) {

	id, err := nanoid.New()
	if err != nil {
		return nil, err
	}

	m := &Message{
		// Message meta data
		ID:       id,
		MimeType: ma.BROADCAST_MIME_TYPE,
		Version:  ma.VERSION,
		// Recipients
		From: from,
		// Body
		ContentType: contentType,
		// The content is not signed as such, but the hash is.
		Content: content,
	}

	err = m.Sign(priv_key)
	if err != nil {
		return nil, fmt.Errorf("msg_new: failed to sign message: %w", err)
	}

	return m, nil

}

// Send message to a pubsub topic
func (m *Message) Broadcast(ctx context.Context, t *pubsub.Topic) error {

	// Verify that the message is valid before sending it
	err := m.Verify()
	if err != nil {
		return err
	}

	err = m.verifyBroadcast(t)
	if err != nil {
		return err
	}

	msgBytes, err := cbor.Marshal(m)
	if err != nil {
		return fmt.Errorf("send: message serialization error: %w", err)
	}

	// Post the *unencrypted* message to the topic
	t.Publish(ctx, msgBytes)

	return nil
}

func (m *Message) verifyBroadcast(t *pubsub.Topic) error {
	if m.MimeType != ma.BROADCAST_MIME_TYPE {
		return ErrMessageInvalidType
	}

	if m.To != "" {
		return ErrBroadcastHasRecipient
	}

	if t.String() != ma.BROADCAST_TOPIC {
		return ErrBroadcastInvalidTopic
	}

	return nil
}
