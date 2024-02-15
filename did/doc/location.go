package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
)

func (d *Document) SetLastKnowLocation(location string) error {

	// location must be a valid did!
	_did, err := did.New(location)
	if err != nil {
		return fmt.Errorf("doc/SetLastKnowLocation: %w", err)
	}

	d.LastKnownLocation = _did.DID

	d.UpdateVersion()

	return nil

}
