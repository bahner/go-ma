package msg

import (
	"context"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// Send message to a pubsub topic
// We could have done p2p here, but it's much cleaner to
// take a topic as an argument. The topic is just a medium
// for sending messages, not actually the recipient necessarily.
func (m *Message) Send(ctx context.Context, t *pubsub.Topic) error {

	// Create a new envelope
	e, err := m.Enclose()
	if err != nil {
		return fmt.Errorf("send: envelope creation error: %w", err)
	}

	eBytes, err := e.Bytes()
	if err != nil {
		return fmt.Errorf("send: envelope serialization error: %w", err)
	}

	t.Publish(ctx, eBytes, nil)

	return nil
}
