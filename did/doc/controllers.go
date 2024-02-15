package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
)

func (d *Document) AddController(controller string) error {

	err := verifyController(controller)
	if err != nil {
		return fmt.Errorf("doc/AddController: %w", err)
	}

	// Check if the controller is already in the slice
	for _, c := range d.Controller {
		if c == controller {
			return nil // Controller already exists, do nothing
		}
	}

	// Append the new controller since it's not already present
	d.Controller = append(d.Controller, controller)

	return nil
}
func (d *Document) DeleteController(controller string) {
	for i, c := range d.Controller {
		if c == controller {
			d.Controller = append(d.Controller[:i], d.Controller[i+1:]...)
		}
	}
}

func (vm VerificationMethod) ValidateControllers() error {

	if len(vm.Controller) == 0 {
		return ErrControllersIsEmpty
	}

	for _, c := range vm.Controller {
		if !isValidController(c) {
			return did.ErrInvalidDID
		}
	}
	return nil
}

func verifyController(controller string) error {
	if !did.IsValidDID(controller) {
		return fmt.Errorf("controller is not a valid DID: %s. %w", controller, did.ErrInvalidDID)
	}
	return nil
}

func isValidController(controller string) bool {
	return verifyController(controller) == nil
}
