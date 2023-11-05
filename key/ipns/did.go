package ipns

import (
	"github.com/libp2p/go-libp2p/core/crypto"

	"github.com/bahner/go-ma"
)

func CreateDIDFromPublicKeyAndName(pubKey crypto.PubKey, name string) (string, error) {

	ipnsName, err := IdentifierFromPublicKey(pubKey)
	if err != nil {
		return "", err
	}

	return ma.DID_PREFIX + ipnsName.String() + "#" + name, nil

}
