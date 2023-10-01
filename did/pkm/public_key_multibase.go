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

	log "github.com/sirupsen/logrus"
)

// PublicKeyMultibase structure
type PublicKeyMultibase struct {
	MulticodecCodeString string
	PublicKey            crypto.PublicKey
	PublicKeyMultibase   string
}

// New constructs a PublicKeyMultibase instance given a key (either public or private).
// Multiencode the key and return it as a multibase encoded string.
func New(key interface{}) (*PublicKeyMultibase, error) {
	var pubBytes []byte
	var err error
	var multicodecCode multicodec.Code
	var pubKey crypto.PublicKey

	// So this is compact, but it's just the same over and over again.
	// Set the codec depednfing on key type
	switch k := key.(type) {
	case *rsa.PrivateKey:
		multicodecCode = multicodec.RsaPub
		pubBytes, err = x509.MarshalPKIXPublicKey(&k.PublicKey)
		pubKey = &k.PublicKey
	case *rsa.PublicKey:
		multicodecCode = multicodec.RsaPub
		pubBytes, err = x509.MarshalPKIXPublicKey(k)
		pubKey = k
	case *ed25519.PrivateKey:
		multicodecCode = multicodec.Ed25519Pub
		pubBytes = k.Public().(ed25519.PublicKey)
		pubKey = k.Public().(ed25519.PublicKey)
	case *ed25519.PublicKey:
		multicodecCode = multicodec.Ed25519Pub
		pubBytes = *k
		pubKey = k
	default:
		return nil, fmt.Errorf("unsupported key type %T", k)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %v", err)
	}

	log.Debugf("Public key bytes: %v", pubBytes)
	log.Debugf("Public key bytes length: %v", len(pubBytes))
	log.Debugf("Public key: %v", pubKey)
	log.Debugf("Public key type: %T", pubKey)
	log.Debugf("Public key multicodec: %v", multicodecCode)
	multicodecBytes := varint.ToUvarint(uint64(multicodecCode))
	prefixedPubBytes := append(multicodecBytes, pubBytes...)

	multibaseStr, err := multibase.Encode(multibase.Base58BTC, prefixedPubBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to multibase encode the prefixed public key bytes: %v", err)
	}

	return &PublicKeyMultibase{
		MulticodecCodeString: multicodecCode.String(),
		PublicKey:            pubKey,
		PublicKeyMultibase:   multibaseStr,
	}, nil
}

func Parse(pkmb string) (*PublicKeyMultibase, error) {

	// Step 1: Multibase Decoding
	_, decodedBytes, err := multibase.Decode(pkmb)
	if err != nil {
		return nil, fmt.Errorf("error decoding multibase string: %v", err)
	}

	log.Debugf("Multibase decoded bytes: %v", decodedBytes)
	log.Debugf("Multibase decoded bytes length: %v", len(decodedBytes))

	// Step 2: Reading and removing the Multicodec Prefix
	codec, n, err := varint.FromUvarint(decodedBytes)
	if err != nil {
		return nil, fmt.Errorf("error reading multicodec varint: %v", err)
	}

	log.Debugf("Multicodec varint: %v", codec)
	log.Debugf("Multicodec name: %v", multicodec.Code(codec).String())

	// Step 3: Extracting the Public Key Bytes based on Multicodec
	pubKeyBytes := decodedBytes[n:]
	var pub crypto.PublicKey
	// switch multicodec.Code(codec) {
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

func (pkmb *PublicKeyMultibase) String() string {
	return "did:key:" + pkmb.PublicKeyMultibase
}
