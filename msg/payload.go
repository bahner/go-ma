package msg

import (
	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
)

// Returns a copy of the Message payload
func Payload(m Message) (Message, error) {

	m.Signature = ""

	return m, nil
}

func (m *Message) MarshalPayloadToCBOR() ([]byte, error) {

	payload, err := Payload(*m)
	if err != nil {
		return nil, err
	}

	marshalled_payload, err := cbor.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return marshalled_payload, nil
}

// Returns the Message as a multibase encoded JSON string
// with the Signature field set to the empty string.
// NB! This is made from a copy of the message.
// The original Message is not changed.
// This is what we will sign!
func (m *Message) PayloadPack() (string, error) {

	marshalled_payload, err := m.MarshalPayloadToCBOR()
	if err != nil {
		return "", err
	}
	encoded_payload, err := internal.MultibaseEncode(marshalled_payload)
	if err != nil {
		return "", err
	}

	return encoded_payload, nil
}
