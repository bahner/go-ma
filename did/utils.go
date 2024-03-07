package did

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
)

var fragment = regexp.MustCompile("^#[a-zA-Z0-9_-]*$")

func IsValid(didStr string) bool {

	return Validate(didStr) == nil
}

func Validate(didStr string) error {

	if didStr == "" {
		return ErrEmptyDID
	}

	// Check that the string has the correct prefix
	if !strings.HasPrefix(didStr, ma.DID_PREFIX) {
		return ErrInvalidPrefix
	}

	name := strings.TrimPrefix(didStr, ma.DID_PREFIX)
	parts := strings.Split(name, "#")

	if len(parts) == 0 {
		return ErrMissingIdentifier
	}
	if len(parts) == 1 {
		return ErrMissingFragment
	}

	if len(parts) > 2 {
		return ErrInvalidFormat
	}

	identifier := parts[0]
	fragment := parts[1]

	err := verifyIdentifier(identifier)
	if err != nil {
		return err
	}

	if !isValidFragment("#" + fragment) {
		return fmt.Errorf("invalid fragment: %s, %w", fragment, ErrInvalidFragment)
	}

	return nil
}

// AreIdentical checks if two DIDs have the same ID, ignoring the fragment.
// This is a stretched interpretation of the word Identical, but
// I couldn't help myself. It means they are derived from the same key,
// which makes them Identical in my book.
func AreIdentical(did1 DID, did2 DID) bool {
	return did1.Identifier == did2.Identifier
}

// Get the fragment from the DID string
// The prefix is not required, ut'll just be stripped off.
func getFragment(did string) string {

	didName := strings.TrimPrefix(did, ma.DID_PREFIX)

	elements := strings.Split(didName, "#")

	return elements[len(elements)-1]
}

// Get the identifier from the DID string
// The prefix is not required, it'll just be stripped off.
func getIdentifier(did string) string {

	didName := strings.TrimPrefix(did, ma.DID_PREFIX)

	elements := strings.Split(didName, "#")

	return elements[0]
}

func isValidFragment(fragment string) bool {
	return verifyFragment(fragment) == nil
}

// This simply checks that the string is a valid nanoID,
// prefixed with a "#".
func verifyFragment(str string) error {

	ok := fragment.MatchString(str)
	if !ok {
		return fmt.Errorf("invalid fragment: %s, %w", str, ErrInvalidFragment)
	}

	return nil
}

func verifyIdentifier(id string) error {

	if !internal.IsValidIPNSName(id) {
		return fmt.Errorf("invalid identifier: %s, %w", id, ErrInvalidIdentifier)
	}

	return nil
}
