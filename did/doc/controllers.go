package doc

func (d *Document) AddController(controller string) {
	// Check if the controller is already in the slice
	for _, c := range d.Controller {
		if c == controller {
			return // Controller already exists, do nothing
		}
	}

	// Append the new controller since it's not already present
	d.Controller = append(d.Controller, controller)
}
func (d *Document) DeleteController(controller string) {
	for i, c := range d.Controller {
		if c == controller {
			d.Controller = append(d.Controller[:i], d.Controller[i+1:]...)
		}
	}
}
