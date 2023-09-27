package pubkey

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
)

type PublicKeyMultibase struct {
	Type               string
	PublicKeyMultibase string
}

// Takes a secret key and returns a multibase encoded public key.
// The return is a struct which also has the correspongding key type.
func New(secretKey interface{}) (*PublicKeyMultibase, error) {
	var pubBytes []byte
	var err error
	var keyType string
	var multicodecCode multicodec.Code

	switch sk := secretKey.(type) {
	case *rsa.PrivateKey:
		keyType = "RsaVerificationKey2018"
		multicodecCode = multicodec.RsaPub

		// Marshal RSA Public Key to ASN.1 DER encoded form
		pubBytes, err = x509.MarshalPKIXPublicKey(&sk.PublicKey)
	case *crypto.Ed25519PrivateKey:
		keyType = "Ed25519VerificationKey2018"
		multicodecCode = multicodec.Ed25519Pub

		// For ed25519, the public key is directly available as part of the private key
		pubBytes, err = sk.Raw()
		if err != nil {
			return nil, fmt.Errorf("failed to get raw public key bytes: %v", err)
		}

	default:
		return nil, fmt.Errorf("unsupported secret key type %T", sk)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %v", err)
	}

	// Convert multicodec code to varint before appending
	multicodecBytes := varint.ToUvarint(uint64(multicodecCode))
	prefixedPubBytes := append(multicodecBytes, pubBytes...)

	multibaseStr, err := multibase.Encode(multibase.Base58BTC, prefixedPubBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to multibase encode the prefixed public key bytes: %v", err)
	}

	return &PublicKeyMultibase{
		Type:               keyType,
		PublicKeyMultibase: multibaseStr,
	}, nil
}

func Decode(publicKeyMultibase string) (*rsa.PublicKey, error) {
	// Decode the multibase-encoded public key string
	_, decoded, err := multibase.Decode(publicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("parse: error decoding multibase string: %s", err)
	}

	// Read and remove the multicodec prefix
	_, n, err := varint.FromUvarint(decoded)
	if err != nil {
		return nil, fmt.Errorf("parse: error reading multicodec varint: %s", err)
	}
	pubKeyBytes := decoded[n:]

	// Parse the decoded bytes (after removing multicodec prefix) to get the *rsa.PublicKey
	pub, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("parse: error parsing public key: %s", err)
	}

	rsaPubKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("parse: not an RSA public key")
	}

	return rsaPubKey, nil
}
