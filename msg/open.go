package msg

import (
	"fmt"
)

// While this is a very empty file, it's intended to be a starting point for
// validating an doing decision making on the message envelope.
// Eg. validate headers before processing the content

func Open(data []byte, privateKey []byte) (*Message, error) {
	e, err := unmarshalEnvelopeFromCBOR(data)
	if err != nil {
		return nil, fmt.Errorf("open: error unmarshalling envelope: %w", err)
	}

	// Extract headers before decrypting the content
	hdrs, err := e.getHeaders(privateKey)
	if err != nil {
		return nil, fmt.Errorf("open: error decrypting headers: %w", err)
	}

	err = hdrs.validate()
	if err != nil {
		return nil, fmt.Errorf("open: error validating headers: %w", err)
	}

	// Create message from the extracted headers sans content
	m, err := NewFromHeaders(hdrs)
	if err != nil {
		return nil, fmt.Errorf("open: error creating message from headers: %w", err)
	}

	// Verify the signature before proceeding
	m.Verify()

	m.Content, err = e.getContent(privateKey)
	if err != nil {
		return nil, fmt.Errorf("open: error decrypting content: %w", err)
	}

	return m, nil

}
