package message

import (
	"time"

	"github.com/Masterminds/semver"
	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	"github.com/libp2p/go-libp2p/core/crypto"
	nanoid "github.com/matoous/go-nanoid/v2"
)

// This struct mimicks the Message format, but it's *not* Message.
// It should enable using Message later, if that's a good idea.
type Message struct {
	ID           string `json:"id"`
	MimeType     string `json:"mime_type"`
	From         string `json:"from"`
	To           string `json:"to"`
	Created      string `json:"created"`
	Expires      string `json:"expires"`
	BodyMimeType string `json:"body_mime_type"`
	Body         string `json:"body"`
	Version      string `json:"version"`
	Signature    string `json:"signature"`
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
	created_time := now.Format(time.RFC3339)
	expires_time := now.Add(MESSAGE_TTL).Format(time.RFC3339)

	return &Message{
		ID:           id,
		MimeType:     ma.MESSAGE_MIME_TYPE,
		Version:      ma.VERSION,
		From:         from,
		To:           to,
		Created:      created_time,
		Expires:      expires_time,
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
	return internal.CreateTimeFromIsoString(m.Created)
}

func (m *Message) ExpiresTime() (time.Time, error) {
	return internal.CreateTimeFromIsoString(m.Expires)
}

func (m *Message) Sender() (*did.DID, error) {
	return did.Parse(m.From)
}

func (m *Message) Recipient() (*did.DID, error) {
	return did.Parse(m.To)
}

func (m *Message) SemVersion() (*semver.Version, error) {
	return semver.NewVersion(m.Version)
}
