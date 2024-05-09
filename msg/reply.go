package msg

import (
	"context"
	"crypto/ed25519"
	"mime"

	"github.com/fxamacker/cbor/v2"
	nanoid "github.com/matoous/go-nanoid/v2"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const REPLY_SERIALIZATION = "cbor"

var REPLY_CONTENT_TYPE_PARAMS = map[string]string{
	"type": "reply",
}

type ReplyContent struct {
	// Id of the messagew to reply to
	RequestID string `cbor:"requestID"`
	// Reply content
	Reply []byte `cbor:"reply"`
}

func NewReply(m *Message, reply []byte) ([]byte, error) {
	return cbor.Marshal(
		&ReplyContent{
			RequestID: m.Id,
			Reply:     reply,
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

	mimeType := MESSAGE_CONTENT_TYPE + "+" + REPLY_SERIALIZATION
	contentType := mime.FormatMediaType(mimeType, REPLY_CONTENT_TYPE_PARAMS)

	reply := &Message{
		// Message meta data
		Id:   id,
		Type: MESSAGE,
		// Recipient
		From: m.To,
		To:   m.From,
		// Body
		ContentType: contentType,
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
