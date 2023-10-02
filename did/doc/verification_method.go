package doc

import (
	"crypto"
	"errors"
	"fmt"

	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/internal"
	"github.com/multiformats/go-multibase"
)

// Specification of encryption and signing key types
var encryptionKeyTypes = []string{"x25519-pub", "x448-pub"}
var signingKeyType = "ed25519-pub"

func (d *Document) AddVerificationMethod(method vm.VerificationMethod) error {
	// Before appending the method, check if id or publicKeyMultibase is unique
	if err := d.isUniqueVerificationMethod(method); err != nil {
		return fmt.Errorf("doc/vm: error adding verification method: %s", err)
	}

	d.VerificationMethod = append(d.VerificationMethod, method)

	return nil
}

func (d *Document) VerificationMethodsOfType(multicodec string) ([]crypto.PublicKey, error) {
	var keys []crypto.PublicKey

	for _, method := range d.VerificationMethod {

		_, data, err := multibase.Decode(method.PublicKeyMultibase)
		if err != nil {
			return nil, fmt.Errorf("doc/vm: error decoding public key multibase: %s", err)
		}

		codec, decoded, err := internal.MulticodecDecode(data)
		if err != nil {
			return nil, fmt.Errorf("doc/vm: error decoding public key: %s", err)
		}

		if codec == multicodec {
			keys = append(keys, decoded)
		}

	}

	return keys, nil
}

// EncryptionKey returns the encryption key from the document's VerificationMethod.
func (d *Document) EncryptionKey() (crypto.PublicKey, error) {
	for _, multicodecStr := range encryptionKeyTypes {
		keys, err := d.VerificationMethodsOfType(multicodecStr)
		if err == nil && len(keys) > 0 { // if there is no error and we found a key
			return keys[0], nil
		}
	}
	return nil, errors.New("no encryption key found")
}

// SigningKey returns the signing key from the document's VerificationMethod.
func (d *Document) SigningKey() (crypto.PublicKey, error) {
	keys, err := d.VerificationMethodsOfType(signingKeyType)
	if err != nil || len(keys) == 0 {
		return nil, errors.New("no signing key found")
	}
	return keys[0], nil
}

func (d *Document) isUniqueVerificationMethod(newMethod vm.VerificationMethod) error {
	for _, existingMethod := range d.VerificationMethod {
		if existingMethod.ID == newMethod.ID {
			return errors.New("duplicate id found in Verification Methods")
		}
		if existingMethod.PublicKeyMultibase == newMethod.PublicKeyMultibase {
			return errors.New("duplicate publicKeyMultibase found in Verification Methods")
		}
	}
	return nil // Return nil if no duplicate found
}
