package msg

import (
	"crypto/ed25519"
	"fmt"

	"github.com/bahner/go-ma"
	nanoid "github.com/matoous/go-nanoid/v2"
)

const (
	REQUEST = PREFIX + "request/" + ma.VERSION

	REQUEST_CONTENT_TYPE = "application/cbor"
)

func Request(
	from string,
	to string,
	content []byte,
	priv_key ed25519.PrivateKey) (*Message, error) {

	id, err := nanoid.New()
	if err != nil {
		return nil, err
	}

	m := &Message{
		// Message meta data
		Id:   id,
		Type: REQUEST,
		// Recipient
		From: from,
		To:   to,
		// Body
		ContentType: REQUEST_CONTENT_TYPE,
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
