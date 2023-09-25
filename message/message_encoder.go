package message

import (
	"github.com/multiformats/go-multibase"
)

// This function is a little sugar to use the proper multibase encoding for
// message attributes and the message itself.
func MessageEncoder(attr []byte) (string, error) {

	encoded_attr, err := multibase.Encode(MESSAGE_ENCODER_ENCODING, attr)
	if err != nil {
		return "", err
	}

	return encoded_attr, nil
}
