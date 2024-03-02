package multi

import (
	"github.com/multiformats/go-multibase"
)

// This function is a little sugar to use the proper multibase encoding for
// message attributes and the message itself.
func MultibaseEncode(data []byte) (string, error) {

	encoded_attr, err := multibase.Encode(multibase.Base58BTC, data)
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

func IsValidMultibase(input string) bool {
	_, _, err := multibase.Decode(input)
	return err == nil
}
