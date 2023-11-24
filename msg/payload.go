package msg

import (
	"github.com/bahner/go-ma/internal"

	cbor "github.com/fxamacker/cbor/v2"
	"lukechampine.com/blake3"
)

// Returns a copy of the Message payload
func (m *Message) Unsigned() (Message, error) {

	c := *m
	c.Signature = ""

	return c, nil
}

func (m *Message) MarshalUnsignedToCBOR() ([]byte, error) {

	payload, err := m.Unsigned()
	if err != nil {
		return nil, err
	}

	marshalled_payload, err := cbor.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return marshalled_payload, nil
}

// Returns the multibase-encoded hash of the Message payload
// Or empty string if an error occurs
// This is mostly intended for debugging purposes
func (m *Message) MultibaseEncodedPayloadHash() string {

	marshalled_payload, err := m.MarshalUnsignedToCBOR()
	if err != nil {
		return ""
	}

	payloadHash := blake3.Sum256(marshalled_payload)

	multibaseEncodedPayloadHash, err := internal.MultibaseEncode(payloadHash[:])
	if err != nil {
		return ""
	}

	return multibaseEncodedPayloadHash

}
