package msg

import (
	"crypto/ed25519"
	"fmt"

	"github.com/bahner/go-ma"
	cbor "github.com/fxamacker/cbor/v2"
	nanoid "github.com/matoous/go-nanoid/v2"
)

const (
	PREFIX       = "/ma/"
	MESSAGE      = PREFIX + ma.VERSION
	CONTENT_TYPE = "application/x.ma"
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
	contentType string,
	content []byte,
	priv_key ed25519.PrivateKey) (*Message, error) {

	id, err := nanoid.New()
	if err != nil {
		return nil, err
	}

	m := &Message{
		// Message meta data
		Id:   id,
		Type: MESSAGE,
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

func (m *Message) Bytes() ([]byte, error) {
	return cbor.Marshal(m)
}

// UnmarshalMessage unmarshals a Message from a CBOR byte slice
// and verifies the signature
func UnmarshalMessage(b []byte) (*Message, error) {
	var m *Message = new(Message)
	err := cbor.Unmarshal(b, m)
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_message_from_cbor: failed to unmarshal message: %w", err)
	}
	return m, nil
}

func UnmarshalAndVerifyMessage(b []byte) (*Message, error) {
	m, err := UnmarshalMessage(b)
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_message_from_cbor_and_verify_signature: failed to unmarshal message: %w", err)
	}

	err = m.Verify()
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_message_from_cbor_and_verify_signature: failed to verify message: %w", err)
	}

	return m, nil
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
