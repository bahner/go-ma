package message

import (
	"encoding/json"

	"github.com/bahner/go-ma/internal"
)

// Returns a copy of the Message payload
func Payload(m Message) (Message, error) {

	m.Signature = ""

	return m, nil
}

func (m *Message) MarshalPayloadToJSON() ([]byte, error) {

	payload, err := Payload(*m)
	if err != nil {
		return nil, err
	}

	marshalled_payload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return marshalled_payload, nil
}

// Returns the Message as a multibase encoded JSON string
// with the Signature field set to the empty string.
// NB! This is made from a copy of the Message.
// The original Message is not changed.
// This is what we will sign!
func (m *Message) PayloadPack() (string, error) {

	marshalled_payload, err := m.MarshalPayloadToJSON()
	if err != nil {
		return "", err
	}
	encoded_payload, err := internal.MultibaseEncode(marshalled_payload)
	if err != nil {
		return "", err
	}

	return encoded_payload, nil
}
