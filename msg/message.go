package msg

import (
	"context"
	"crypto/ed25519"
	"fmt"

	cbor "github.com/fxamacker/cbor/v2"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	nanoid "github.com/matoous/go-nanoid/v2"
)

const (
	PREFIX               = "/ma/message/"
	DEFAULT_CONTENT_TYPE = "text/plain"
)

// This struct mimicks the Message format, but it's *not* Message.
// It should enable using Message later, if that's a good idea.
type Message struct {
	// Unique identifier of the message
	Id string `cbor:"id"`
	// MA message type of the message
	Type string `cbor:"type"`
	// Sender of the message
	From string `cbor:"from"`
	// Recipient of the message
	To string `cbor:"to"`
	// MIME type of the message body
	ContentType string `cbor:"contentType"`
	// Byte slice of the content
	Content []byte `cbor:"content"`
	// Signature of the message headers. NB! This includes the ContentHash field,
	// which can be used to verify the integrity of the message body.
	Signature []byte `cbor:"signature"`
}

// New creates a new Message instance
func New(
	from string,
	to string,
	msg_type string,
	content []byte,
	contentType string,
	priv_key ed25519.PrivateKey) (*Message, error) {

	id, err := nanoid.New()
	if err != nil {
		return nil, err
	}

	m := &Message{
		// Message meta data
		Id:   id,
		Type: msg_type,
		// Recipient
		From: from,
		To:   to,
		// Body
		ContentType: contentType,
		Content:     content,
	}

	err = verifyContent(content)
	if err != nil {
		return nil, err
	}

	err = m.Sign(priv_key)
	if err != nil {
		return nil, fmt.Errorf("msg_new: failed to sign message: %w", err)
	}

	return m, nil
}

// UnmarshalMessageFromCBOR unmarshals a Message from a CBOR byte slice
// and verifies the signature
func UnmarshalMessageFromCBOR(b []byte) (*Message, error) {
	var m *Message = new(Message)
	err := cbor.Unmarshal(b, m)
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_message_from_cbor: failed to unmarshal message: %w", err)
	}
	return m, nil
}

func UnmarshalAndVerifyMessageFromCBOR(b []byte) (*Message, error) {
	m, err := UnmarshalMessageFromCBOR(b)
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_message_from_cbor_and_verify_signature: failed to unmarshal message: %w", err)
	}

	err = m.Verify()
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_message_from_cbor_and_verify_signature: failed to verify message: %w", err)
	}

	return m, nil
}

//	message send is meant to send the message signed, but not encrypted
//
// Create an envelope by calling Enclose() first and then send *that* with Send()
func (m *Message) Send(ctx context.Context, t *pubsub.Topic) error {

	eBytes, err := cbor.Marshal(m)
	if err != nil {
		return fmt.Errorf("send: envelope serialization error: %w", err)
	}

	// t.Publish(ctx, eBytes, nil)
	t.Publish(ctx, eBytes)

	return nil
}

// Create a new Message from the headers
// Validate the headers (sinature) before adding the content. This is to be lazy
// about decrypting the content, in case we don't need it.
func newFromHeaders(h *Headers) (*Message, error) {

	err := h.validate()
	if err != nil {
		return nil, fmt.Errorf("newFromHeaders: %w", err)
	}

	m := &Message{
		// Message meta data
		Id:   h.Id,
		Type: h.Type,
		// Recipient
		From: h.From,
		To:   h.To,
		// Body
		ContentType: h.ContentType,
		// Signature
		Signature: h.Signature,
	}

	return m, nil
}
