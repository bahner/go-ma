package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
)

func (d *Document) AddContext(context string) error {

	err := did.Validate(context)
	if err != nil {
		return fmt.Errorf("doc/AddContext: %w", err)
	}

	// Check if the context is already in the slice
	for _, c := range d.Context {
		if c == context {
			return nil // Context already exists, do nothing
		}
	}

	// Append the new context since it's not already present
	d.Context = append(d.Context, context)

	return nil
}
func (d *Document) DeleteContext(context string) {
	for i, c := range d.Context {
		if c == context {
			d.Context = append(d.Context[:i], d.Context[i+1:]...)
		}
	}
}

func (d *Document) validateContexts() error {

	if len(d.Context) == 0 {
		return ErrContextIsEmpty
	}

	return nil
}
