package internal

import (
	"fmt"

	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
	log "github.com/sirupsen/logrus"
)

func MulticodecEncode(codecName string, payload []byte) ([]byte, error) {

	var officialCodec multicodec.Code
	if err := officialCodec.Set(codecName); err != nil {
		return nil, fmt.Errorf("error setting codec: %w", err)
	}
	codec := uint64(officialCodec)

	codecBytes := varint.ToUvarint(codec)
	encoded := append(codecBytes, payload...)
	return encoded, nil
}

// Returns the codec name, payload and error of a multicodec encoded byte array
func MulticodecDecode(encoded []byte) (string, []byte, error) {
	if len(encoded) < 1 {
		return "", nil, fmt.Errorf("error decoding: insufficient data")
	}

	log.Debugf("mutlticodecdecode: encoded: %x", encoded)

	code, n, err := varint.FromUvarint(encoded)
	if err != nil {
		return "", nil, fmt.Errorf("error decoding varint: %w", err)
	}
	if n < 1 || n >= len(encoded) {
		return "", nil, fmt.Errorf("error decoding: invalid varint size")
	}
	log.Debugf("mutlticodecdecode: code %d", code)

	codecName := multicodec.Code(code).String()
	if codecName == "" {
		return "", nil, fmt.Errorf("error obtaining codec name: unknown codec")
	}

	log.Debugf("mutlticodecdecode: codecName: %s", codecName)
	return codecName, encoded[n:], nil
}
