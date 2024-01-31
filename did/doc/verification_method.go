package doc

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	mb "github.com/multiformats/go-multibase"
	log "github.com/sirupsen/logrus"
)

// VerificationMethod defines the structure of a Verification Method
type VerificationMethod struct {
	// The full name of the verification method, eg. did:ma:123456789abcdefghi#signature-key-id
	ID string `cbor:"id" json:"id"`
	// The type of verification method. We only use MultiKey. It's unofficial, but it works.
	// https://w3c-ccg.github.io/did-method-key/
	Type string `cbor:"type" json:"type"`
	// The controller of the verification method. This is the DID of the entity that controls the key.
	// Should probably always be the DID itself, but maybe the DID controller.
	Controller []string `cbor:"controller,toarray" json:"controller"`
	// Created is the time the verification method was created
	PublicKeyMultibase string `cbor:"publicKeyMultibase" json:"publicKeyMultibase"`
}

// NewVerificationMethod creates a new VerificationMethod
// id is the identifier of the verification method, eg.
// k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr
// id must be a valid IPNS name.
// A random fragment will be added to the id
// vmType must be of type MultiKey, so that the format never changes -
// even if the underlying key type changes.
func NewVerificationMethod(
	// DID of the subject to create a verification method for
	id string,
	// DID of the controller of the verification method
	controller string,
	// The DID suite key type specified in the verification method
	vmType string,
	// Name (fragment) of the verification method. If "", a random fragment will be generated
	// Fragment should not be prefixed with "#"
	fragment string,
	// The public key of the verification method, encoded in multibase
	publicKeyMultibase string,
) (VerificationMethod, error) {

	if !did.IsValidDID(id) {
		return VerificationMethod{}, did.ErrInvalidID
	}

	identifier := did.GetIdentifier(id)

	// Create a random fragment if none is provided
	if fragment == "" {
		fragment = did.GenerateFragment()
	} else {
		fragment = "#" + fragment
	}

	if !did.IsValidFragment(fragment) {
		return VerificationMethod{}, did.ErrInvalidFragment
	}

	return VerificationMethod{
		ID:                 ma.DID_PREFIX + identifier + fragment,
		Controller:         []string{controller},
		Type:               vmType,
		PublicKeyMultibase: publicKeyMultibase,
	}, nil
}

// Add a controller to the verification method. The DID will be
// checked for validity before adding it to the slice.
func (v *VerificationMethod) AddController(controller string) error {

	// Check if the controller is empty
	if controller == "" {
		return fmt.Errorf("doc/vm: controller is empty")
	}

	// Make sure the controller is a valid DID
	if !did.IsValidDID(controller) {
		return fmt.Errorf("doc/vm: controller is not a valid DID")
	}

	// Check if the controller is already in the slice
	for _, c := range v.Controller {
		if c == controller {
			return nil
		}
	}

	v.Controller = append(v.Controller, controller)

	return nil
}

func (v *VerificationMethod) DeleteController(controller string) {
	for i, c := range v.Controller {
		if c == controller {
			v.Controller = append(v.Controller[:i], v.Controller[i+1:]...)
		}
	}
}

func (d *Document) AddVerificationMethod(method VerificationMethod) error {

	// Make sure the verification method is valid
	err := method.Verify()
	if err != nil {
		return fmt.Errorf("doc/vm: %w", err)
	}

	// Before appending the method, check if id or publicKeyMultibase is unique
	if d.isUniqueVerificationMethod(method) {
		d.VerificationMethod = append(d.VerificationMethod, method)
	}

	return nil
}

func (d *Document) DeleteVerificationMethod(method VerificationMethod) error {
	for i, m := range d.VerificationMethod {
		if m.ID == method.ID {
			d.VerificationMethod = append(d.VerificationMethod[:i], d.VerificationMethod[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("doc/vm: error deleting verification method: %s", method.ID)
}

func (d *Document) GetVerificationMethodbyID(vmid string) (VerificationMethod, error) {

	for _, method := range d.VerificationMethod {

		if method.ID == vmid {
			return method, nil
		}
	}

	return VerificationMethod{}, fmt.Errorf("doc/vm: no verification method found with id: %s", vmid)
}

func (d *Document) isUniqueVerificationMethod(newMethod VerificationMethod) bool {

	for _, existingMethod := range d.VerificationMethod {
		if existingMethod.ID == newMethod.ID {
			return false
		}
		if existingMethod.PublicKeyMultibase == newMethod.PublicKeyMultibase {
			return false
		}
	}
	log.Debugf("doc/vm: verification method %s is unique", newMethod.ID)
	return true
}

func (vm VerificationMethod) Fragment() string {
	return did.GetFragment(vm.ID)
}

func (vm VerificationMethod) Equal(other VerificationMethod) bool {
	if vm.ID != other.ID {
		log.Debugf("vm/Equal: ID not equal")
		return false
	}

	if vm.Type != other.Type {
		log.Debugf("vm/Equal: Type not equal")
		return false
	}

	if !compareSlices(vm.Controller, other.Controller) {
		log.Debugf("vm/Equal: Controllers not equal")
		return false
	}

	if vm.PublicKeyMultibase != other.PublicKeyMultibase {
		log.Debugf("vm/Equal: PublicKeyMultibase not equal")
		return false
	}

	return true
}

// Simply verify that the verification method seems valid
func (vm VerificationMethod) Verify() error {
	if vm.ID == "" {
		return fmt.Errorf("vm/Verify: ID is empty")
	}

	if !did.IsValidDID(vm.ID) {
		return fmt.Errorf("vm/Verify: ID is not a valid DID")
	}

	if vm.Type == "" {
		return fmt.Errorf("vm/Verify: Type is empty")
	}

	err := vm.ValidateControllers()
	if err != nil {
		return fmt.Errorf("vm/Verify: %w", err)
	}

	if vm.PublicKeyMultibase == "" {
		return fmt.Errorf("vm/Verify: PublicKeyMultibase is empty")
	}

	_, _, err = mb.Decode(vm.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf("vm/Verify: %w", err)
	}

	return nil

}

func (vm VerificationMethod) IsValid() bool {

	return vm.Verify() == nil
}

func (vm VerificationMethod) ValidateControllers() error {

	if len(vm.Controller) == 0 {
		return fmt.Errorf("vm/ValidateControllers: Controller is empty")
	}

	for _, c := range vm.Controller {
		if !did.IsValidDID(c) {
			return fmt.Errorf("vm/ValidateControllers: Controller is not a valid DID")
		}
	}
	return nil
}
