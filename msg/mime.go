package msg

import (
	"fmt"
	"mime"
	"strings"
)

var (
	ErrInvalidContentType = fmt.Errorf("invalid Message ContentType")
	ErrMissingContentType = fmt.Errorf("empty ContentType")
	ErrMissingSerialiser  = fmt.Errorf("missing serialiser")
)

// Returns the message type as defined in the contenttype, not the message as such
func (m *Message) MessageType() (string, error) {

	_, params, err := m.parseContentType()
	if err != nil {
		return "", ErrInvalidContentType
	}

	for k, v := range params {
		if k == "type" {
			return v, nil
		}
	}

	return "", ErrInvalidContentType
}

// Returns the message type as defined in the contenttype, not the message as such
func (m *Message) MessageSerialiser() (string, error) {

	contentType, _, err := m.parseContentType()
	if err != nil {
		return "", ErrInvalidContentType
	}

	elements := strings.Split(contentType, "+")
	if len(elements) == 1 {
		return contentType, fmt.Errorf("%s: %w", contentType, ErrMissingSerialiser)
	}

	return elements[1], nil
}

// A parser which add a little specific functionality to the standard mime.ParseMediaType
// Use this instead.
func (m *Message) parseContentType() (string, map[string]string, error) {

	contentType, params, err := mime.ParseMediaType(m.ContentType)
	if err != nil {
		return contentType, params, fmt.Errorf("msg_parse_content_type: %w", err)
	}

	if contentType == "" {
		return contentType, params, ErrMissingContentType
	}

	if contentType != CONTENT_TYPE {
		return contentType, params, fmt.Errorf("%s: %w", contentType, ErrInvalidContentType)
	}

	return contentType, params, nil
}
