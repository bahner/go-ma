package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/bahner/go-ma/did/pkm"
	"github.com/bahner/go-ma/internal"

	ipfs "github.com/ipfs/go-ipfs-api"
)

type Keys struct {
	Id                 string // IPNS name from kubo
	Name               string // short name from kubo, eg Nanoid you specified.
	RSAPrivateKey      *rsa.PrivateKey
	RSAPublicKey       *rsa.PublicKey
	Ed25519PrivateKey  *ed25519.PrivateKey
	Ed25519PublicKey   *ed25519.PublicKey
	PublicKeyMultibase string
}

func New(ApiKey *ipfs.Key) (*Keys, error) {

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

	ed25519PublicKey, ed25519PrivateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		internal.LogError(fmt.Sprintf("key: failed to generate ed25519 keypair: %v", err))
		return nil, err
	}

	return &Keys{
		Id:                 ApiKey.Id,
		Name:               ApiKey.Name,
		RSAPrivateKey:      rsaPrivateKey,
		RSAPublicKey:       &rsaPrivateKey.PublicKey,
		Ed25519PrivateKey:  &ed25519PrivateKey,
		Ed25519PublicKey:   &ed25519PublicKey,
		PublicKeyMultibase: pkmk.PublicKeyMultibase,
	}, nil
}
