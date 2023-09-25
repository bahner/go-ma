package key

import (
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multibase"
	log "github.com/sirupsen/logrus"
)

func packPrivKey(privKey crypto.PrivKey) (string, error) {

	marshalledPrivKey, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		log.Errorf("Failed to marshal private key: %v", err)
		return "", err
	}

	// Multibase encode the marshalled private key. Again honouring the constants.
	encodedPrivKey, err := multibase.Encode(PRIVATE_KEY_ENCODING, marshalledPrivKey)
	if err != nil {
		return "", err
	}

	return encodedPrivKey, nil
}

func unpackPrivateKey(packedPrivKey string) (crypto.PrivKey, error) {

	_, decoded_private_key, err := multibase.Decode(packedPrivKey)
	if err != nil {
		log.Errorf("Failed to marshal private key: %v", err)
		return nil, err
	}

	unmarshalledPrivKey, err := crypto.UnmarshalPrivateKey(decoded_private_key)
	if err != nil {
		return nil, err
	}

	return unmarshalledPrivKey, nil
}
