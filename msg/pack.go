package msg

import (
	"encoding/json"

	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/multiformats/go-multibase"
)

func (m *Message) Pack() (string, error) {

	data, err := m.MarshalToCBOR()
	if err != nil {
		return "", err
	}

	encoded_data, err := internal.MultibaseEncode(data)
	if err != nil {
		return "", err
	}

	return encoded_data, nil
}

func Unpack(data string) (*Message, error) {

	_, decoded_data, err := multibase.Decode(data)
	if err != nil {
		return nil, err
	}

	return UnmarshalFromCBOR(decoded_data)
}

// Returns a copy of the Message payload as a byte slice
// No error messages, just nil
// Try with IsValidPack() first, if you need to make sure.
func (m *Message) Bytes() []byte {

	packedMessage, err := m.Pack()
	if err != nil {
		return nil
	}

	return []byte(packedMessage)
}

func (m *Message) MarshalToCBOR() ([]byte, error) {

	return cbor.Marshal(m)
}

func (m *Message) MarshalToJSON() ([]byte, error) {

	return json.Marshal(m)
}

func IsValidPack(data string) bool {

	_, err := Unpack(data)
	return err != nil
}

func UnmarshalFromCBOR(data []byte) (*Message, error) {

	m := &Message{}

	err := cbor.Unmarshal(data, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
