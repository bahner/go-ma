package doc

import (
	"fmt"

	"github.com/ipfs/go-cid"
)

func (d *Document) SetIdentity(identity string) error {

	err := validateIdentity(identity)
	if err != nil {
		return fmt.Errorf("doc/AddIdentity: %w", err)
	}

	d.Identity = identity

	return nil
}

func validateIdentity(identity string) error {

	_, err := cid.Parse(identity)

	return err
}
