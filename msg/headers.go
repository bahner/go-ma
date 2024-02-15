package msg

import (
	"time"

	"github.com/bahner/go-ma/did"
	semver "github.com/blang/semver/v4"
	cbor "github.com/fxamacker/cbor/v2"
)

// This struct mimicks the Headers format, but it's *not* Headers.
// It should enable using Headers later, if that's a good idea.
// NB! Content is *not* a part of the headers
type Headers struct {
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
	Signature []byte `cbor:"signature"`
}

func (m *Message) baseHeaders() Headers {

	return Headers{
		// Message Headers
		ID:          m.ID,
		MimeType:    m.MimeType,
		Version:     m.Version,
		From:        m.From,
		To:          m.To,
		ContentType: m.ContentType,
	}
}

func (m *Message) unsignedHeaders() Headers {

	return m.baseHeaders()
}

func (m *Message) marshalUnsignedHeadersToCBOR() ([]byte, error) {
	return cbor.Marshal(m.unsignedHeaders())
}

// Returns the all the imprimatur headers
func (m *Message) Headers() Headers {

	hdrs := m.baseHeaders()
	hdrs.Signature = m.Signature

	return hdrs
}

func (m *Message) marshalHeadersToCBOR() ([]byte, error) {
	return cbor.Marshal(m.Headers())
}

func (h *Headers) CreatedTime() time.Time {
	return time.Unix(h.Created, 0)
}

func (h *Headers) ExpiresTime() time.Time {
	return time.Unix(h.Expires, 0)
}

func (h *Headers) Sender() (did.DID, error) {
	return did.New(h.From)
}

func (h *Headers) Recipient() (did.DID, error) {
	return did.New(h.To)
}

func (h *Headers) SemVersion() (semver.Version, error) {
	return semver.Make(h.Version)
}
