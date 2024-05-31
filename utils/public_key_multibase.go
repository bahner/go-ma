package utils

import (
	"fmt"

	"github.com/multiformats/go-multicodec"
)

func PublicKeyMultibaseEncode(codec multicodec.Code, publicKey []byte) (string, error) {

	multicodecedKey := MulticodecEncode(codec, publicKey)

	publicKeyMultibase, err := MultibaseEncode(multicodecedKey)
	if err != nil {
		return "", fmt.Errorf("key/codec: error multibase encoding public key: %s", err)
	}

	return publicKeyMultibase, nil

}

func PublicKeyMultibaseDecode(publicKey string) (multicodec.Code, []byte, error) {

	var codec multicodec.Code

	decodedPublicKeyMultibase, err := MultibaseDecode(publicKey)
	if err != nil {
		return codec, nil, err
	}

	codec, decodedPublicKey, err := MulticodecDecode(decodedPublicKeyMultibase)
	if err != nil {
		return codec, nil, err
	}

	return codec, decodedPublicKey, nil

}
