package headers

import (
	"crypto/ed25519"
	"time"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/msg/body"
	mime "github.com/bahner/go-ma/msg/mime"
	semver "github.com/blang/semver/v4"
	cbor "github.com/fxamacker/cbor/v2"
	nanoid "github.com/matoous/go-nanoid/v2"
)

const (

	// Messages which are older than a day should be ignored
	MESSAGE_TTL = time.Hour * 24

	// How we identify the messages we support
	MESSAGE_ENCRYPTION_LABEL = mime.MESSAGE_MIME_TYPE
)

// This struct mimicks the Headers format, but it's *not* Headers.
// It should enable using Headers later, if that's a good idea.
type Headers struct {
	_ struct{} `cbor:",toarray"`
	// Version of the message format
	Version string `cbor:"versionId" json:"versionId"`
	// Unique identifier of the message
	ID string `cbor:"id" json:"id"`
	// MIME type of the message
	MimeType string `cbor:"type" json:"type"`
	// Creation time of the message in seconds since Unix epoch
	Created int64 `cbor:"created_time,keyasint64" json:"created_time"`
	// Expiration time of the message in seconds since Unix epoch
	Expires int64 `cbor:"expires_time,keyasint64" json:"expires_time"`
	// Sender of the message
	From string `cbor:"from" json:"from"`
	// Recipient of the message
	To string `cbor:"to" json:"to"`
	// MIME type of the message body
	ContentType string `cbor:"content_type" json:"content_type"`
	// Hexadecimal string representation of the SHA-256 hash of the message body
	ContentHash string `cbor:"content_hash" json:"content_hash"`
	// Signature of the message headers. NB! This includes the ContentHash field,
	// which can be used to verify the integrity of the message body.
	Signature string `cbor:"signature" json:"signature"`
}

// New creates a new Message instance
// Message is a string for now, but it should be JSON.
func New(
	from string,
	to string,
	body body.Body) (*Headers, error) {

	id, err := nanoid.New()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	created := now.Unix()
	expires := now.Add(MESSAGE_TTL).Unix()

	return &Headers{
		// Message
		ID:       id,
		MimeType: mime.MESSAGE_MIME_TYPE,
		Version:  ma.VERSION,
		// Headers
		From:    from,
		To:      to,
		Created: created,
		Expires: expires,
		// Body
		ContentType: body.ContentType,
		ContentHash: body.Hash(),
		// Signature
		Signature: "",
	}, nil
}

func Signed(
	from string,
	to string,
	body body.Body,
	priv_key *ed25519.PrivateKey) (*Headers, error) {

	m, err := New(from, to, body)
	if err != nil {
		return nil, err
	}

	m.Sign(priv_key)

	return m, nil

}

func (m *Headers) CreatedTime() (time.Time, error) {
	return time.Unix(m.Created, 0), nil
}

func (m *Headers) ExpiresTime() (time.Time, error) {
	return time.Unix(m.Expires, 0), nil
}

func (m *Headers) Sender() (*did.DID, error) {
	return did.New(m.From)
}

func (m *Headers) Recipient() (*did.DID, error) {
	return did.New(m.To)
}

func (m *Headers) SemVersion() (semver.Version, error) {
	return semver.Make(m.Version)
}

func (m *Headers) MarshalToCBOR() ([]byte, error) {
	return cbor.Marshal(m)
}

func UnmarshalFromCBOR(data []byte) (*Headers, error) {
	var m Headers
	err := cbor.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
