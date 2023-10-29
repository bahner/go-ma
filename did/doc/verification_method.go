package doc

import (
	"errors"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
)

// VerificationMethod defines the structure of a Verification Method
type VerificationMethod struct {
	// The full name of the verification method, eg. did:key:123456789abcdefghi#signature
	ID string `json:"id"`
	// The type of verification method. We only use MultiKey. It's unofficial, but it works.
	// https://w3c-ccg.github.io/did-method-key/
	Type string `json:"type"`
	// The controller of the verification method. This is the DID of the entity that controls the key.
	// Should probably always be the DID itself, but maybe the DID controller.
	Controller string `json:"controller"`
	// The public key of the verification method, encoded in multibase, as
	// per the did core spec.
	PublicKeyMultibase string `json:"publicKeyMultibase"`
}

// NewVerificationMethod creates a new VerificationMethod
// id is the identifier of the verification method, eg.
// k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr
// id must be a valid IPNS name.
// A random fragment will be added to the id
// vmType must be of type MultiKey, so that the format never changes -
// even if the underlying key type changes.
func NewVerificationMethod(
	id string,
	controller string,
	publicKeyMultibase string) (VerificationMethod, error) {

	if !internal.IsValidIPNSName(id) {
		return VerificationMethod{}, internal.ErrInvalidID
	}

	return VerificationMethod{
		ID:                 key.DID_KEY_PREFIX + id + internal.GenerateFragment(),
		Controller:         controller,
		Type:               ma.VERIFICATION_METHOD_KEY_TYPE,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}

func (d *Document) AddVerificationMethod(method VerificationMethod) error {
	// Before appending the method, check if id or publicKeyMultibase is unique
	if err := d.isUniqueVerificationMethod(method); err != nil {
		return fmt.Errorf("doc/vm: error adding verification method: %s", err)
	}

	d.VerificationMethod = append(d.VerificationMethod, method)

	return nil
}

func (d *Document) GetVerificationMethodbyID(vmid string) (VerificationMethod, error) {

	for _, method := range d.VerificationMethod {

		if method.ID == vmid {
			return method, nil
		}
	}

	return VerificationMethod{}, fmt.Errorf("doc/vm: no verification method found with id: %s", vmid)
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

func (vm VerificationMethod) GetID() string {
	return vm.ID
}

func (vm VerificationMethod) Fragment() string {
	return internal.GetFragmentFromDID(vm.ID)
}
