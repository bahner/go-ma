package did

import (
	"fmt"
	"strings"

	"github.com/ipfs/boxo/ipns"
)

// Validates the DID string and returns the ipns.Name and fragment
func Parse(didStr string) (ipns.Name, string, error) {

	if didStr == "" {
		return ipns.Name{}, "", ErrEmptyDID
	}

	// Check that the string has the correct prefix
	if !strings.HasPrefix(didStr, PREFIX) {
		return ipns.Name{}, "", ErrInvalidPrefix
	}

	_id := strings.TrimPrefix(didStr, PREFIX)
	parts := strings.Split(_id, "#")

	if len(parts) == 0 {
		return ipns.Name{}, "", ErrMissingIdentifier
	}
	if len(parts) == 1 {
		return ipns.Name{}, "", ErrMissingFragment
	}

	if len(parts) > 2 {
		return ipns.Name{}, "", ErrInvalidFormat
	}

	// Extract the name from the string
	name, err := getName(didStr)
	if err != nil {
		return ipns.Name{}, "", fmt.Errorf("invalid identifier: %s, %w", parts[0], ErrInvalidName)
	}

	// Extract the fragment from the string
	fragment := getFragment(didStr)
	if !isValidFragment("#" + fragment) {
		return ipns.Name{}, "", fmt.Errorf("invalid fragment: %s, %w", fragment, ErrInvalidFragment)
	}

	return name, fragment, nil
}

func Validate(didStr string) error {
	_, _, err := Parse(didStr)
	return err
}
