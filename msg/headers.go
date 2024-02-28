package msg

import (
	cbor "github.com/fxamacker/cbor/v2"
)

// This struct mimicks the Headers format, but it's *not* Headers.
// It should enable using Headers later, if that's a good idea.
// NB! Content is *not* a part of the headers
type Headers struct {
	// Version of the message format
	Id string `cbor:"id"`
	// MIME type of the message
	Type string `cbor:"type"`
	// Sender of the message
	From string `cbor:"from"`
	// Recipient of the message
	To string `cbor:"to"`
	// MIME type of the message body
	ContentType string `cbor:"contentType"`
	// Hexadecimal string representation of the SHA-256 hash of the message body
	Signature []byte `cbor:"signature"`
}

func (m *Message) unsignedHeaders() Headers {

	return Headers{
		// Message Headers
		Id:          m.Id,
		Type:        m.Type,
		From:        m.From,
		To:          m.To,
		ContentType: m.ContentType,
	}
}

func (m *Message) marshalUnsignedHeadersToCBOR() ([]byte, error) {
	return cbor.Marshal(m.unsignedHeaders())
}

// Returns the all the imprimatur headers
func (m *Message) Headers() Headers {

	hdrs := m.unsignedHeaders()
	hdrs.Signature = m.Signature

	return hdrs
}

func (m *Message) marshalHeadersToCBOR() ([]byte, error) {
	return cbor.Marshal(m.Headers())
}
