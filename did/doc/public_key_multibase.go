package doc

import (
	"fmt"

	mb "github.com/multiformats/go-multibase"
)

func verifyPublicKeyMultibase(multibase string) error {

	if multibase == "" {
		return ErrPublicKeyMultibaseEmpty
	}

	_, _, err := mb.Decode(multibase)
	if err != nil {
		return fmt.Errorf("vm/Verify: %w", err)
	}

	return nil
}
