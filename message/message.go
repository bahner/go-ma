package message

import (
	"time"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	semver "github.com/blang/semver/v4"
	"github.com/libp2p/go-libp2p/core/crypto"
	nanoid "github.com/matoous/go-nanoid/v2"
)

// This struct mimicks the Message format, but it's *not* Message.
// It should enable using Message later, if that's a good idea.
type Message struct {
	_            struct{} `cbor:",toarray"`
	ID           string   `cbor:"id"`
	MimeType     string   `cbor:"type"`
	From         string   `cbor:"from"`
	To           string   `cbor:"to"`
	Created      int64    `cbor:"created_time,keyasint64"`
	Expires      int64    `cbor:"expires_time,keyasint64"`
	BodyMimeType string   `cbor:"body_mime_type"`
	Body         string   `cbor:"body"`
	Version      string   `cbor:"version"`
	Signature    string   `cbor:"signature"`
}

// New creates a new Message instance
// Message is a string for now, but it should be JSON.
func New(
	from string,
	to string,
	body string,
	mime_type string) (*Message, error) {

	id, err := nanoid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	created := now.Unix()
	expires := now.Add(MESSAGE_TTL).Unix()

	return &Message{
		ID:           id,
		MimeType:     ma.MESSAGE_MIME_TYPE,
		Version:      ma.VERSION,
		From:         from,
		To:           to,
		Created:      created,
		Expires:      expires,
		BodyMimeType: mime_type,
		Body:         body,
		Signature:    "",
	}, nil
}

func Signed(
	from string,
	to string,
	body string,
	mime_type string,
	priv_key crypto.PrivKey) (*Message, error) {

	msg, err := New(from, to, body, mime_type)
	if err != nil {
		return nil, err
	}

	msg.Sign(priv_key)

	return msg, nil

}

func (m *Message) CreatedTime() (time.Time, error) {
	return time.Unix(m.Created, 0), nil
}

func (m *Message) ExpiresTime() (time.Time, error) {
	return time.Unix(m.Expires, 0), nil
}

func (m *Message) Sender() (*did.DID, error) {
	return did.Parse(m.From)
}

func (m *Message) Recipient() (*did.DID, error) {
	return did.Parse(m.To)
}

func (m *Message) SemVersion() (semver.Version, error) {
	return semver.Make(m.Version)
}
