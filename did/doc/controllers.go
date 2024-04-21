package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
)

func (d *Document) AddController(controller string) error {

	err := did.Validate(controller)
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
		err := did.Validate(c)
		if err != nil {
			return fmt.Errorf("doc/ValidateControllers: %w", err)
		}
	}
	return nil
}
