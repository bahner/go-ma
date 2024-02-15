package did

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"github.com/ipfs/boxo/ipns"
	nanoid "github.com/matoous/go-nanoid/v2"
)

var fragment = regexp.MustCompile("^#[a-zA-Z0-9_-]*$")

func IsValidDID(didStr string) bool {

	return ValidateDID(didStr) == nil
}

func ValidateDID(didStr string) error {

	if didStr == "" {
		return ErrEmptyDID
	}

	// Check that the string has the correct prefix

	if !strings.HasPrefix(didStr, ma.DID_PREFIX) {
		return ErrInvalidPrefix
	}

	// Identifier here is a bit of a misnomer. It's the whole name.
	identifier := strings.TrimPrefix(didStr, ma.DID_PREFIX)

	name := strings.Split(identifier, "#")
	id := name[0]
	fragment := name[1]

	if len(name) != 2 {
		return ErrInvalidFormat
	}

	if !internal.IsValidIPNSName(id) {
		return fmt.Errorf("invalid identifier: %s, %w", id, ErrInvalidIdentifier)
	}

	if !internal.IsValidNanoID(name[1]) {
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

func GenerateFragment() string {
	fragment, _ := nanoid.New()

	return "#" + fragment
}

func GetFragment(did string) string {

	didName := strings.TrimPrefix(did, ma.DID_PREFIX)

	elements := strings.Split(didName, "#")

	return elements[len(elements)-1]
}

func GetIdentifier(did string) string {

	didName := strings.TrimPrefix(did, ma.DID_PREFIX)

	elements := strings.Split(didName, "#")

	return elements[0]
}

// This simply checks that the string is a valid nanoID,
// prefixed with a "#".
func IsValidFragment(str string) bool {
	return fragment.MatchString(str)
}

func VerifyFragment(fragment string) error {
	if !IsValidFragment(fragment) {
		return ErrInvalidFragment
	}
	return nil
}

func IsValidIdentifier(identifier string) bool {
	return VerifyIdentifier(identifier) == nil
}

func VerifyIdentifier(identifier string) error {

	parts := strings.Split(identifier, "#")
	if len(parts) != 2 {
		return ErrInvalidFormat
	}

	// Check that the identifier has a valid fragment
	fragment := parts[1]
	if !IsValidFragment(fragment) {
		return fmt.Errorf("invalid fragment: %s, %w", fragment, ErrInvalidFragment)
	}

	// Check that the id is a valid IPNS name
	_, err := ipns.NameFromString(identifier)
	if err != nil {
		return fmt.Errorf("invalid identifier: %s, %w", identifier, ErrInvalidIdentifier)
	}

	return nil
}
