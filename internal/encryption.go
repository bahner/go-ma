package internal

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"

	libp2pcrypto "github.com/libp2p/go-libp2p/core/crypto"
	log "github.com/sirupsen/logrus"
)

func StdPubKeyToLibp2pPubKey(pubKey crypto.PublicKey) (libp2pcrypto.PubKey, error) {
	if pubKey == nil {
		return nil, errors.New("received nil public key")
	}

	var der []byte
	var err error

	switch key := pubKey.(type) {
	case *rsa.PublicKey:
		// Marshal *rsa.PublicKey to ASN.1 DER encoded form
		der, err = x509.MarshalPKIXPublicKey(key)
		if err != nil {
			return nil, fmt.Errorf("error marshaling RSA public key: %w", err)
		}
	case *ecdsa.PublicKey:
		// Marshal *ecdsa.PublicKey to ASN.1 DER encoded form
		der, err = x509.MarshalPKIXPublicKey(key)
		if err != nil {
			return nil, fmt.Errorf("error marshaling ECDSA public key: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported public key type %T", key)
	}

	// Unmarshal ASN.1 DER form to libp2p PubKey
	libp2pPubKey, err := libp2pcrypto.UnmarshalPublicKey(der)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling to libp2p public key: %w", err)
	}

	return libp2pPubKey, nil
}

func Libp2pPubKeyToStdRSAPubKey(pubKey libp2pcrypto.PubKey) (*rsa.PublicKey, error) {
	standardPubKey, err := libp2pcrypto.PubKeyToStdKey(pubKey)
	if err != nil {
		log.Errorf("Error converting libp2p public key to standard public key: %v", err)
		return nil, err
	}

	if rsaPubKey, ok := standardPubKey.(*rsa.PublicKey); ok {
		log.Debugf("Converted libp2p public key to standard public key: %v", rsaPubKey)
		return rsaPubKey, nil
	}

	return nil, fmt.Errorf("not an RSA public key, actual type: %T", standardPubKey)
}

func Libp2pPubKeyToStdECDSAPubKey(pubKey libp2pcrypto.PubKey) (*ecdsa.PublicKey, error) {
	standardPubKey, err := libp2pcrypto.PubKeyToStdKey(pubKey)
	if err != nil {
		log.Errorf("Error converting libp2p public key to standard public key: %v", err)
		return nil, err
	}

	if ecdsaPubKey, ok := standardPubKey.(*ecdsa.PublicKey); ok {
		log.Debugf("Converted libp2p public key to standard public key: %v", pubKey)
		return ecdsaPubKey, nil
	}

	return nil, fmt.Errorf("not aa valid public key, actual type: %T", standardPubKey)
}

func StdPrivKeyToLibp2pPrivKey(privKey crypto.PrivateKey) (libp2pcrypto.PrivKey, error) {
	if privKey == nil {
		return nil, errors.New("received nil private key")
	}

	var der []byte
	var err error

	switch key := privKey.(type) {
	case *rsa.PrivateKey:
		der = x509.MarshalPKCS1PrivateKey(key)
	case *ecdsa.PrivateKey:
		der, err = x509.MarshalECPrivateKey(key)
		if err != nil {
			return nil, fmt.Errorf("error marshaling ECDSA private key: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported private key type %T", key)
	}

	libp2pPrivKey, err := libp2pcrypto.UnmarshalPrivateKey(der)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling to libp2p private key: %w", err)
	}

	return libp2pPrivKey, nil
}

func Libp2pPrivKeyToStdRSAPrivKey(privKey libp2pcrypto.PrivKey) (*rsa.PrivateKey, error) {
	standardPrivKey, err := libp2pcrypto.PrivKeyToStdKey(privKey)
	if err != nil {
		log.Errorf("Error converting libp2p private key to standard private key: %v", err)
		return nil, err
	}

	if rsaPrivKey, ok := standardPrivKey.(*rsa.PrivateKey); ok {
		log.Debugf("Converted libp2p private key to standard private key: %v", rsaPrivKey)
		return rsaPrivKey, nil
	}

	return nil, fmt.Errorf("not an RSA private key, actual type: %T", standardPrivKey)
}

func Libp2pPrivKeyToStdECDSAPrivKey(privKey libp2pcrypto.PrivKey) (*ecdsa.PrivateKey, error) {
	standardPrivKey, err := libp2pcrypto.PrivKeyToStdKey(privKey)
	if err != nil {
		log.Errorf("Error converting libp2p private key to standard private key: %v", err)
		return nil, err
	}

	if ecdsaPrivKey, ok := standardPrivKey.(*ecdsa.PrivateKey); ok {
		log.Debugf("Converted libp2p private key to standard private key: %v", ecdsaPrivKey)
		return ecdsaPrivKey, nil
	}

	return nil, fmt.Errorf("not a valid ECDSA private key, actual type: %T", standardPrivKey)
}

func IsStdRSAPubKey(pubKey crypto.PublicKey) bool {
	if _, ok := pubKey.(*rsa.PublicKey); ok {
		return true
	}

	return false
}

func IsStdEcPubKey(pubKey crypto.PublicKey) bool {
	if _, ok := pubKey.(*ecdsa.PublicKey); ok {
		return true
	}

	return false
}

func IsStdRSAPrivKey(privKey crypto.PrivateKey) bool {
	if _, ok := privKey.(*rsa.PrivateKey); ok {
		return true
	}

	return false
}

func IsStdEcPrivKey(privKey crypto.PrivateKey) bool {
	if _, ok := privKey.(*ecdsa.PrivateKey); ok {
		return true
	}

	return false
}

func HashData(algo crypto.Hash, data []byte) ([]byte, error) {
	if !algo.Available() {
		return nil, fmt.Errorf("hash algorithm %v is not available", algo)
	}

	hasher := algo.New()
	_, err := hasher.Write(data)
	if err != nil {
		return nil, fmt.Errorf("error hashing data: %v", err)
	}

	return hasher.Sum(nil), nil
}
