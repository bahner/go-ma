package msg

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/bahner/go-ma"
	cbor "github.com/fxamacker/cbor/v2"
	nanoid "github.com/matoous/go-nanoid/v2"
)

const (

	// Messages which are older than a day should be ignored
	MESSAGE_TTL = ma.MESSAGE_DEFAULT_TTL
)

// This struct mimicks the Message format, but it's *not* Message.
// It should enable using Message later, if that's a good idea.
type Message struct {
	// Version of the message format
	Version string
	// Unique identifier of the message
	ID string `cbor:"id"`
	// MIME type of the message
	MimeType string `cbor:"mimeType"`
	// Creation time of the message in seconds since Unix epoch
	Created int64 `cbor:"keyasint64"`
	// Expiration time of the message in seconds since Unix epoch
	Expires int64 `cbor:"keyasint64"`
	// Sender of the message
	From string `cbor:"from"`
	// Recipient of the message
	To string `cbor:"to"`
	// MIME type of the message body
	ContentType string `cbor:"contentType"`
	// Hexadecimal string representation of the SHA-256 hash of the message body
	Content []byte `cbor:"content"`
	// Signature of the message headers. NB! This includes the ContentHash field,
	// which can be used to verify the integrity of the message body.
	Signature []byte `cbor:"signature"`
}

// New creates a new Message instance
func New(
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
		MimeType: ma.MESSAGE_MIME_TYPE,
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

// Create a new Message from the headers
// Validate the headers (sinature) before adding the content. This is to be lazy
// about decrypting the content, in case we don't need it.
func newFromHeaders(h *Headers) (*Message, error) {

	err := h.validate()
	if err != nil {
		return nil, fmt.Errorf("msg_new_from_headers: failed to validate headers: %w", err)
	}

	m := &Message{
		// Message meta data
		ID:       h.ID,
		MimeType: h.MimeType,
		Version:  h.Version,
		// Recipients
		From: h.From,
		To:   h.To,
		// Timestamps
		Created: h.Created,
		Expires: h.Expires,
		// Body
		ContentType: h.ContentType,
		// Signature
		Signature: h.Signature,
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
