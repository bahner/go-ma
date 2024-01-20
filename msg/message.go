package msg

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/msg/body"
	"github.com/bahner/go-ma/msg/headers"
	cbor "github.com/fxamacker/cbor/v2"
)

// Bask the encrypted message and the encrypted symmetric key in a JSON envelope.
type Message struct {
	Headers *headers.Headers `cbor:"headers" json:"headers"`
	Body    *body.Body       `cbor:"body" json:"body"`
}

// Use a pointer here, this might be arbitrarily big.
// from, to, content are required, but content_type defaults to text/plain
func New(from string,
	to string,
	content []byte,
	content_type string) (*Message, error) {

	// A little sugar goes a long way
	if content_type == "" {
		content_type = MESSAGE_DEFAULT_CONTENT_TYPE
	}

	msgBody, err := body.New(content, content_type)
	if err != nil {
		return nil, fmt.Errorf("msg_new: error creating body: %w", err)
	}

	msgHeaders, err := headers.New(from, to, msgBody)
	if err != nil {
		return nil, fmt.Errorf("msg_new: error creating headers: %w", err)
	}

	return &Message{
		Headers: msgHeaders,
		Body:    msgBody,
	}, nil
}

func (e *Message) MarshalToCBOR() ([]byte, error) {
	return cbor.Marshal(e)
}

func (e *Message) MarshalToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func UnmarshalFromCBOR(data []byte) (*Message, error) {

	e := &Message{}

	err := cbor.Unmarshal(data, e)
	if err != nil {
		return nil, fmt.Errorf("envelope: error unmarshalling envelope: %s", err)
	}

	return e, nil
}

// Simply the CBOR encoded envelope, but a generic named function.
// Returns nil if an error occurs
func (e *Message) Bytes() []byte {

	bytes, err := e.MarshalToCBOR()
	if err != nil {
		return nil
	}

	return bytes
}

// Returns the multibase-encoded hash of the entire envelope
// Returns empty string if an error occurs
func (e *Message) String() string {

	str, err := internal.MultibaseEncode(e.Bytes())
	if err != nil {
		return ""
	}

	return str
}
