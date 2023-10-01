package internal

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
)

func MulticodecEncode(codecName string, payload []byte) ([]byte, error) {
	var codec uint64
	var err error
	codec, err = GetPrivateCodecValue(codecName)
	if err != nil {
		// It's not a private multicodec, try to set as an official multicodec
		var officialCodec multicodec.Code
		if err := officialCodec.Set(codecName); err != nil {
			return nil, fmt.Errorf("error setting codec: %w", err)
		}
		codec = uint64(officialCodec)
	}

	codecBytes := varint.ToUvarint(codec)
	encoded := append(codecBytes, payload...)
	return encoded, nil
}
func MulticodecDecode(encoded []byte) (string, []byte, error) {
	if len(encoded) < 1 {
		return "", nil, fmt.Errorf("error decoding: insufficient data")
	}

	code, n, err := varint.FromUvarint(encoded)
	if err != nil {
		return "", nil, fmt.Errorf("error decoding varint: %w", err)
	}
	if n < 1 || n >= len(encoded) {
		return "", nil, fmt.Errorf("error decoding: invalid varint size")
	}

	var codecName string
	codecName, err = GetPrivateCodecName(code)
	if err != nil {
		// It's not a private multicodec, try to get as an official multicodec
		codec := multicodec.Code(code)
		codecName = codec.String()
		if codecName == "" {
			return "", nil, fmt.Errorf("error obtaining codec name: unknown codec")
		}
	}
	return codecName, encoded[n:], nil
}

// GetPrivateCodecName maps a private multicodec value to its corresponding name
func GetPrivateCodecName(code uint64) (string, error) {
	switch code {

	// X25519 encryption codecs
	case ma.ECDHX25519ChaCha20Poly1305BLAKE3:
		return "ECDHX25519ChaCha20Poly1305BLAKE3", nil

	// X448 encryption codecs
	case ma.ECDHX448ChaCha20Poly1305BLAKE3:
		return "ECDHX448ChaCha20Poly1305BLAKE3", nil
	}
	return "", fmt.Errorf("unknown private multicodec value: %d", code)
}

// GetPrivateCodecValue maps a private multicodec name to its corresponding value
func GetPrivateCodecValue(name string) (uint64, error) {
	switch name {

	// X25519 encryption codecs
	case "ECDHX25519ChaCha20Poly1305BLAKE3":
		return ma.ECDHX25519ChaCha20Poly1305BLAKE3, nil

	// X448 encryption codecs
	case "ECDHX448ChaCha20Poly1305BLAKE3":
		return ma.ECDHX448ChaCha20Poly1305BLAKE3, nil

	}
	return 0, fmt.Errorf("unknown private multicodec name: %s", name)
}
