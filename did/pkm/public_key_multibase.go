package pkm

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/multiformats/go-multibase"
	"github.com/multiformats/go-multicodec"
	"github.com/multiformats/go-varint"
)

// PublicKeyMultibase structure
type PublicKeyMultibase struct {
	MulticodecCodeString string
	PublicKey            crypto.PublicKey
	PublicKeyMultibase   string
}

// New constructs a PublicKeyMultibase instance given a key (either public or private).
func New(key interface{}) (*PublicKeyMultibase, error) {
	var pubBytes []byte
	var err error
	var multicodecCode multicodec.Code
	var pubKey crypto.PublicKey

	switch k := key.(type) {
	case *rsa.PrivateKey:
		multicodecCode = multicodec.RsaPub
		pubBytes, err = x509.MarshalPKIXPublicKey(&k.PublicKey)
		pubKey = &k.PublicKey
	case *rsa.PublicKey:
		multicodecCode = multicodec.RsaPub
		pubBytes, err = x509.MarshalPKIXPublicKey(k)
		pubKey = k
	case ed25519.PrivateKey:
		multicodecCode = multicodec.Ed25519Pub
		pubBytes = k.Public().(ed25519.PublicKey)
		pubKey = k.Public().(ed25519.PublicKey)
	case ed25519.PublicKey:
		multicodecCode = multicodec.Ed25519Pub
		pubBytes = k
		pubKey = k
	default:
		return nil, fmt.Errorf("unsupported key type %T", k)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %v", err)
	}

	multicodecBytes := varint.ToUvarint(uint64(multicodecCode))
	prefixedPubBytes := append(multicodecBytes, pubBytes...)

	multibaseStr, err := multibase.Encode(multibase.Base58BTC, prefixedPubBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to multibase encode the prefixed public key bytes: %v", err)
	}

	return &PublicKeyMultibase{
		MulticodecCodeString: multicodecCode.String(),
		PublicKeyMultibase:   multibaseStr,
		PublicKey:            pubKey, // Set PublicKey correctly here
	}, nil
}

// Parse decodes a PublicKeyMultibase instance
func Parse(pkmb string) (*PublicKeyMultibase, error) {

	// Decode the multibase-encoded public key string
	_, decoded, err := multibase.Decode(pkmb)
	if err != nil {
		return nil, fmt.Errorf("error decoding multibase string: %v", err)
	}

	// Read and remove the multicodec prefix
	codec, n, err := varint.FromUvarint(decoded)
	if err != nil {
		return nil, fmt.Errorf("error reading multicodec varint: %v", err)
	}

	// Process the key bytes based on multicodec
	pubKeyBytes := decoded[n:]
	var pub crypto.PublicKey
	switch multicodec.Code(codec) {
	case multicodec.Ed25519Pub:
		if len(pubKeyBytes) != ed25519.PublicKeySize {
			return nil, fmt.Errorf("invalid Ed25519 public key length")
		}
		pub = ed25519.PublicKey(pubKeyBytes)
	case multicodec.RsaPub:
		pub, err = x509.ParsePKIXPublicKey(pubKeyBytes)
		if err != nil {
			return nil, fmt.Errorf("error parsing RSA public key: %v", err)
		}
	default:
		return nil, fmt.Errorf("unsupported public key type: %s", multicodec.Code(codec).String())
	}

	// Return the decoded public key structure
	return &PublicKeyMultibase{
		MulticodecCodeString: multicodec.Code(codec).String(),
		PublicKey:            pub,
		PublicKeyMultibase:   pkmb,
	}, nil
}
