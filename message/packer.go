package message

import (
	"encoding/json"

	"github.com/bahner/go-ma/internal"
	"github.com/multiformats/go-multibase"
)

func (m *Message) MarshalToJSON() ([]byte, error) {

	return json.Marshal(m)
}

func (m *Message) Pack() (string, error) {

	data, err := m.MarshalToJSON()
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

	err = json.Unmarshal(decoded_data, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
