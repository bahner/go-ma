package msg

import (
	"context"
	"fmt"

	"github.com/fxamacker/cbor/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// EncloseAndSend message to a pubsub topic
// EncloseAndSend differs from Send in that it encrypts the message
// and encloses it in an envelope before sending it.
func (m *Message) Send(ctx context.Context, t *pubsub.Topic) error {

	// Create a new envelope
	e, err := m.enclose()
	if err != nil {
		return fmt.Errorf("send: envelope creation error: %w", err)
	}

	eBytes, err := cbor.Marshal(e)
	if err != nil {
		return fmt.Errorf("send: envelope serialization error: %w", err)
	}

	// t.Publish(ctx, eBytes, nil)
	t.Publish(ctx, eBytes)

	return nil
}
