package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
)

func (d *Document) AddController(controller string) error {

	if d.Controller == nil {
		return fmt.Errorf("controller slice is nil")
	}

	if !did.IsValidDID(controller) {
		return fmt.Errorf("controller is not a valid DID: %s", controller)
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
