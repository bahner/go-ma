package internal

import "fmt"

func EncodePublicKeyMultibase(publicKey []byte, codecName string) (string, error) {

	multicodecedKey, err := MulticodecEncode(codecName, publicKey)
	if err != nil {
		return "", fmt.Errorf("key/codec: error multicodec encoding public key: %s", err)
	}

	publicKeyMultibase, err := MultibaseEncode(multicodecedKey)
	if err != nil {
		return "", fmt.Errorf("key/codec: error multibase encoding public key: %s", err)
	}

	return publicKeyMultibase, nil

}

func DecodePublicKeyMultibase(publicKey string) (string, []byte, error) {

	decodedPublicKeyMultibase, err := MultibaseDecode(publicKey)
	if err != nil {
		return "", nil, fmt.Errorf("key/codec: error multibase decoding public key: %s", err)
	}

	codecName, decodedPublicKey, err := MulticodecDecode(decodedPublicKeyMultibase)
	if err != nil {
		return "", nil, fmt.Errorf("key/codec: error multicodec decoding public key: %s", err)
	}

	return codecName, decodedPublicKey, nil

}
