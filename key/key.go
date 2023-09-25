package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/multiformats/go-multibase"

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

	rsaPublicKey := &rsaPrivateKey.PublicKey
	rsaPublicKeyBytes := x509.MarshalPKCS1PublicKey(rsaPublicKey)

	publicKeyMultibase, err := multibase.Encode(PRIVATE_KEY_ENCODING, []byte(rsaPublicKeyBytes))
	if err != nil {
		internal.LogError(fmt.Sprintf("key: failed to multibase encode public key: %v", err))
	}

	return &Key{
		Id:                 ApiKey.Id,
		Name:               ApiKey.Name,
		RSAPrivateKey:      rsaPrivateKey,
		RSAPublicKey:       rsaPublicKey,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}
