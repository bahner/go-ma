package internal

import (
	"github.com/bahner/go-ma"
	"github.com/multiformats/go-multibase"
)

// This function is a little sugar to use the proper multibase encoding for
// message attributes and the message itself.
func MultibaseEncode(data []byte) (string, error) {

	encoded_attr, err := multibase.Encode(ma.MULTIBASE_ENCODING, data)
	if err != nil {
		return "", err
	}

	return encoded_attr, nil
}

func MultibaseDecode(attr string) ([]byte, error) {

	_, decoded_attr, err := multibase.Decode(attr)
	if err != nil {
		return nil, err
	}

	return decoded_attr, nil
}
