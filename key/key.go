package key

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/bahner/go-ma/did/pkm"
	"github.com/bahner/go-ma/internal"

	ipfs "github.com/ipfs/go-ipfs-api"
)

type Key struct {
	Id                 string // IPNS name from kubo
	Name               string // short name from kubo, eg Nanoid you specified.
	RSAPrivateKey      *rsa.PrivateKey
	RSAPublicKey       *rsa.PublicKey
	PublicKeyMultibase string
}

func New(ApiKey *ipfs.Key) (*Key, error) {

	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, PRIVATE_KEY_BITS)
	if err != nil {
		internal.LogError(fmt.Sprintf("key: failed to generate private key: %v", err))
		return nil, err
	}

	pkmk, err := pkm.New(rsaPrivateKey)
	if err != nil {
		return nil,
			internal.LogError(fmt.Sprintf("key: failed to generate public key multibase: %v", err))
	}

	return &Key{
		Id:                 ApiKey.Id,
		Name:               ApiKey.Name,
		RSAPrivateKey:      rsaPrivateKey,
		RSAPublicKey:       &rsaPrivateKey.PublicKey,
		PublicKeyMultibase: pkmk.PublicKeyMultibase,
	}, nil
}
