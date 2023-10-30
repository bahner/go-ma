package message

import (
	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/multiformats/go-multibase"
)

func (m *Message) MarshalToCBOR() ([]byte, error) {

	return cbor.Marshal(m)
}

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

	var p Message

	_, decoded_data, err := multibase.Decode(data)
	if err != nil {
		return nil, err
	}

	err = cbor.Unmarshal(decoded_data, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
