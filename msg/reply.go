package msg

import (
	"context"
	"crypto/ed25519"

	"github.com/bahner/go-ma"
	"github.com/fxamacker/cbor/v2"
	nanoid "github.com/matoous/go-nanoid/v2"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const REPLY = PREFIX + "reply/" + ma.VERSION

type Reply struct {
	// Id of the messagew to reply to
	ReplyTo string `cbor:"replyTo"`
	// Reply content
	Reply []byte `cbor:"reply"`
}

func NewReply(replyTo string, reply []byte) ([]byte, error) {
	return cbor.Marshal(
		&Reply{
			ReplyTo: replyTo,
			Reply:   reply,
		})
}

func (m *Message) Reply(replyBytes []byte, privKey ed25519.PrivateKey, topic *pubsub.Topic) error {
	id, err := nanoid.New()
	if err != nil {
		return err
	}

	replyContent, err := NewReply(m.Id, replyBytes)
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
		ContentType: DEFAULT_CONTENT_TYPE,
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

	return envelope.Send(context.Background(), topic)
}
