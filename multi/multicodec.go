package multi

import (
	"fmt"

	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
)

func MulticodecEncode(codec multicodec.Code, payload []byte) ([]byte, error) {

	c := uint64(codec)

	codecBytes := varint.ToUvarint(c)
	encoded := append(codecBytes, payload...)
	return encoded, nil
}

// Returns the codec, payload and error of a multicodec encoded byte array
func MulticodecDecode(encoded []byte) (multicodec.Code, []byte, error) {

	var codec multicodec.Code

	if len(encoded) < 1 {
		return codec, nil, ErrNoInput
	}

	// log.Debugf("mutlticodecdecode: encoded: %x", encoded)

	code, n, err := varint.FromUvarint(encoded)
	if err != nil {
		return codec, nil, fmt.Errorf("error decoding varint: %w", err)
	}
	if n < 1 || n >= len(encoded) {
		return codec, nil, ErrInvalidSize
	}

	codec = multicodec.Code(code)
	if codec == 0 {
		return codec, nil, ErrUnknownCodec
	}

	return codec, encoded[n:], nil
}
