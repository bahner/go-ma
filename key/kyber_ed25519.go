package key

import (
	"crypto"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/suites"
)

type KyberEd25519PrivateKey struct {
	privKey            kyber.Scalar
	pubKey             kyber.Point
	publicKeyMultibase string
	name               string
}

func (k *KyberEd25519PrivateKey) PublicKey() crypto.PublicKey {
	// You will need to cast your kyber.Point to an appropriate crypto.PublicKey type.
	// This might be tricky since crypto.PublicKey is an interface, so you need to figure out
	// how you want to represent this public key in the standard library's crypto context.
	return k.pubKey
}

func (k *KyberEd25519PrivateKey) Encrypt(data []byte) ([]byte, error) {
	return nil, nil // FIXME
}

func (k *KyberEd25519PrivateKey) Decrypt(data []byte) ([]byte, error) {
	// Placeholder: You will need to implement the actual decryption logic.
	return nil, nil // FIXME
}

func (k *KyberEd25519PrivateKey) DID() string {
	return KEY_PREFIX + k.PublicKeyMultibase() + "#" + k.name
}

func (k *KyberEd25519PrivateKey) Name() string {
	return k.name
}

func (k *KyberEd25519PrivateKey) PublicKeyMultibase() string {
	return k.publicKeyMultibase
}

// Generate a new ed25519+Kyber keypair
// This is the default keypair type for ma.
// The name is used to generate the DID and is the "name" of the keypair.
// The Kyber version is post quantum safe and was rcently standardized by NIST,
// but it's not yet widely used. But that's why we're here, right?
func GenerateKyberEd25519PrivateKey(name string) (EncryptionKey, error) {
	suite := suites.MustFind(ma.KYBER_SUITE)
	privKey := suite.Scalar().Pick(suite.RandomStream())

	key, err := assembleKey(privKey, name)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error assembling key: %w", err)
	}

	return key, nil
}

func UnmarshalKyberEd25519PrivateKey(data []byte, name string) (EncryptionKey, error) {
	suite := suites.MustFind(ma.KYBER_SUITE)
	privKey := suite.Scalar()
	err := privKey.UnmarshalBinary(data)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error unmarshalling private key: %w", err)
	}

	encKey, err := assembleKey(privKey, name)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error assembling key: %w", err)
	}
	return encKey, nil
}

func assembleKey(privKey kyber.Scalar, name string) (EncryptionKey, error) {
	suite := suites.MustFind(ma.KYBER_SUITE)
	pubKey := suite.Point().Mul(privKey, nil)

	publicKeyMultibase, err := encodePublicKeyToMultibase(pubKey)
	if err != nil {
		return nil, fmt.Errorf("kyber_ed25519: error encoding public key to multibase: %w", err)
	}

	return &KyberEd25519PrivateKey{
		privKey:            privKey,
		pubKey:             pubKey,
		name:               name,
		publicKeyMultibase: publicKeyMultibase,
	}, nil
}

func encodePublicKeyToMultibase(pubKey kyber.Point) (string, error) {
	pubKeyBytes, err := pubKey.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("kyber_ed25519: error marshalling public key: %w", err)
	}

	return internal.EncodePublicKeyMultibase(pubKeyBytes, "kyber-ed25519-pub")
}
