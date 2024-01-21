package doc

import "fmt"

func (d *Document) AddController(controller string) {
	d.Controller = append(d.Controller, controller)
}

func (d *Document) DeleteController(controller string) error {
	for i, c := range d.Controller {
		if c == controller {
			d.Controller = append(d.Controller[:i], d.Controller[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("doc: error deleting controller: %s", controller)
}
