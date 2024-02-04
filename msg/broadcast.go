package msg

import (
	"context"
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/bahner/go-ma"
	cbor "github.com/fxamacker/cbor/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	nanoid "github.com/matoous/go-nanoid/v2"
)

func NewBroadcast(
	from string,
	to string,
	content []byte,
	contentType string,
	priv_key *ed25519.PrivateKey) (*Message, error) {

	id, err := nanoid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	created := now.Unix()
	expires := now.Add(MESSAGE_TTL).Unix()

	m := &Message{
		// Message meta data
		ID:       id,
		MimeType: ma.BROADCAST_MIME_TYPE,
		Version:  ma.VERSION,
		// Recipients
		From: from,
		To:   to,
		// Timestamps
		Created: created,
		Expires: expires,
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
	if m.Verify() != nil {
		return fmt.Errorf("send: message verification failed")
	}

	msgBytes, err := cbor.Marshal(m)
	if err != nil {
		return fmt.Errorf("send: message serialization error: %w", err)
	}

	// Post the *unencrypted* message to the topic
	t.Publish(ctx, msgBytes)

	return nil
}
