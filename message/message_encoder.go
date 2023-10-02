package message

import (
	"github.com/bahner/go-ma"
	"github.com/multiformats/go-multibase"
)

// This function is a little sugar to use the proper multibase encoding for
// message attributes and the message itself.
func MessageEncoder(attr []byte) (string, error) {

	encoded_attr, err := multibase.Encode(ma.MULTIBASE_ENCODING, attr)
	if err != nil {
		return "", err
	}

	return encoded_attr, nil
}
