package doc

import (
	"crypto"
	"errors"
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	log "github.com/sirupsen/logrus"
)

const (
	encryptionKeyType = "x25519-pub"
	signingKeyType    = "ed25519-pub"
)

var verificationMethodTypes = []string{"MultiKey", "Ed25519VerificationKey2020"}

func (d *Document) AddVerificationMethod(method VerificationMethod) error {
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

		codec, decoded, err := internal.DecodePublicKeyMultibase(method.PublicKeyMultibase)
		log.Debugf("doc/vm: VerificationMethodsOfType: codec: %s", codec)
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
	keys, err := d.VerificationMethodsOfType(encryptionKeyType)
	if err == nil && len(keys) > 0 { // if there is no error and we found a key
		return keys[0], nil
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

func (d *Document) isUniqueVerificationMethod(newMethod VerificationMethod) error {
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

// VerificationMethod defines the structure of a Verification Method
type VerificationMethod struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyMultibase string `json:"publicKeyMultibase"`
}

// NewVerificationMethod creates a new VerificationMethod
// id is the identifier of the verification method, eg.
// k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#signature
// id must be a valid IPNS name
// vmType must be one of MultiKey or Ed25519VerificationKey2020
func NewVerificationMethod(
	id string,
	controller string,
	vmType string,
	publicKeyMultibase string) (VerificationMethod, error) {

	if !internal.IsValidIdentifier(id) {
		return VerificationMethod{}, internal.ErrInvalidID
	}

	return VerificationMethod{
		ID:                 key.DID_KEY_PREFIX + id,
		Type:               vmType,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}

func (vm VerificationMethod) GetID() string {
	return vm.ID
}

func (vm VerificationMethod) Fragment() string {
	return internal.GetFragmentFromDID(vm.ID)
}
