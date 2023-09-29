package doc

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/did/coll"
)

func (d *Document) UnmarshalJSON(data []byte) error {
	type Alias Document
	tmp := &struct {
		Context    interface{} `json:"@context"`
		Controller interface{} `json:"controller"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	// Handle the @context field.
	context, err := coll.New(tmp.Context)
	if err != nil {
		return fmt.Errorf("error creating Collection for @context: %w", err)
	}

	// Perform a safe type assertion for Context.
	contextSet, ok := context.(*coll.CollectionSet)
	if !ok {
		return fmt.Errorf("expected Context to be *coll.CollectionSet, got %T", context)
	}
	d.Context = contextSet

	// Handle the controller field.
	controller, err := coll.New(tmp.Controller)
	if err != nil {
		return fmt.Errorf("error creating Collection for controller: %w", err)
	}

	// Perform a safe type assertion for Controller.
	controllerSet, ok := controller.(*coll.CollectionSet)
	if !ok {
		return fmt.Errorf("expected Controller to be *coll.CollectionSet, got %T", controller)
	}
	d.Controller = controllerSet

	return nil
}
