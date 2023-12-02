package msg

import (
	"crypto/ed25519"
	"time"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	mime "github.com/bahner/go-ma/msg/mime"
	semver "github.com/blang/semver/v4"
	nanoid "github.com/matoous/go-nanoid/v2"
)

const (

	// Messages which are older than a day should be ignored
	MESSAGE_TTL = time.Hour * 24

	// How we identify the messages we support
	MESSAGE_ENCRYPTION_LABEL = mime.MESSAGE_MIME_TYPE
)

// This struct mimicks the Message format, but it's *not* Message.
// It should enable using Message later, if that's a good idea.
type Message struct {
	_            struct{} `cbor:",toarray"`
	ID           string   `cbor:"id" json:"id"`
	MimeType     string   `cbor:"type" json:"type"`
	From         string   `cbor:"from" json:"from"`
	To           string   `cbor:"to" json:"to"`
	Created      int64    `cbor:"created_time,keyasint64" json:"created_time"`
	Expires      int64    `cbor:"expires_time,keyasint64" json:"expires_time"`
	BodyMimeType string   `cbor:"body_mime_type" json:"body_mime_type"`
	Body         []byte   `cbor:"body" json:"body"`
	Version      string   `cbor:"versionId" json:"versionId"`
	Signature    string   `cbor:"signature" json:"signature"`
}

// New creates a new Message instance
// Message is a string for now, but it should be JSON.
func New(
	from string,
	to string,
	body []byte,
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
		MimeType:     mime.MESSAGE_MIME_TYPE,
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
	body []byte,
	mime_type string,
	priv_key *ed25519.PrivateKey) (*Message, error) {

	m, err := New(from, to, body, mime_type)
	if err != nil {
		return nil, err
	}

	m.Sign(priv_key)

	return m, nil

}

func (m *Message) CreatedTime() (time.Time, error) {
	return time.Unix(m.Created, 0), nil
}

func (m *Message) ExpiresTime() (time.Time, error) {
	return time.Unix(m.Expires, 0), nil
}

func (m *Message) Sender() (*did.DID, error) {
	return did.New(m.From)
}

func (m *Message) Recipient() (*did.DID, error) {
	return did.New(m.To)
}

func (m *Message) SemVersion() (semver.Version, error) {
	return semver.Make(m.Version)
}
