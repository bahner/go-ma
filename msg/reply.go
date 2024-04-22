package msg

import (
	"context"
	"crypto/ed25519"

	"github.com/bahner/go-ma"
	"github.com/fxamacker/cbor/v2"
	nanoid "github.com/matoous/go-nanoid/v2"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const (
	REPLY              = PREFIX + "reply/" + ma.VERSION
	REPLY_CONTENT_TYPE = "application/cbor"
)

type ReplyContent struct {
	// Id of the messagew to reply to
	RequestID string `cbor:"requestID"`
	// Type of the message to reply to
	RequestType string `cbor:"requestType"`
	// Reply content
	Reply []byte `cbor:"reply"`
}

func NewReply(m *Message, reply []byte) ([]byte, error) {
	return cbor.Marshal(
		&ReplyContent{
			RequestID:   m.Id,
			RequestType: m.Type,
			Reply:       reply,
		})
}

func (m *Message) Reply(ctx context.Context, replyBytes []byte, privKey ed25519.PrivateKey, topic *pubsub.Topic) error {
	id, err := nanoid.New()
	if err != nil {
		return err
	}

	replyContent, err := NewReply(m, replyBytes)
	if err != nil {
		return err
	}

	reply := &Message{
		// Message meta data
		Id:   id,
		Type: REPLY,
		// Recipient
		From: m.To,
		To:   m.From,
		// Body
		ContentType: REPLY_CONTENT_TYPE,
		Content:     replyContent,
	}

	err = reply.Sign(privKey)
	if err != nil {
		return err
	}

	envelope, err := reply.Enclose()
	if err != nil {
		return err
	}

	return envelope.Send(ctx, topic)
}
