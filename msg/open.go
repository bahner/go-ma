package msg

import (
	"fmt"
)

// While this is a very empty file, it's intended to be a starting point for
// validating an doing decision making on the message envelope.
// Eg. validate headers before processing the content

func (e *Envelope) Open(privateKey []byte) (*Message, error) {

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
	m, err := newFromHeaders(hdrs)
	if err != nil {
		return nil, fmt.Errorf("open: error creating message from headers: %w", err)
	}

	// Verify the signature before proceeding
	err = m.Verify()
	if err != nil {
		return nil, fmt.Errorf("open: error verifying message: %w", err)
	}

	m.Content, err = e.getContent(privateKey)
	if err != nil {
		return nil, fmt.Errorf("open: error decrypting content: %w", err)
	}

	return m, nil

}
