package did

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"github.com/ipfs/boxo/ipns"
	nanoid "github.com/matoous/go-nanoid/v2"
)

var (
	fragment           = regexp.MustCompile("^#[a-zA-Z0-9_-]*$")
	ErrInvalidID       = errors.New("invalid ID")
	ErrInvalidFragment = errors.New("invalid fragment")
)

// // Sometimes we just want the identifier, not the whole DID.
// func ExtractID(didStr string) (string, error) {

// 	did, err := Parse(didStr)
// 	if err != nil {
// 		return "", fmt.Errorf("did: not a valid identifier: %v\n", err))
// 	}

// 	return did.Identifier, nil
// }

func IsValidDID(didStr string) bool {

	err := ValidateDID(didStr)
	return err == nil
}

func ValidateDID(didStr string) error {

	if !strings.HasPrefix(didStr, ma.DID_PREFIX) {
		return fmt.Errorf("invalid DID format, must start with %s", ma.DID_PREFIX)
	}

	identifier := strings.TrimPrefix(didStr, ma.DID_PREFIX)

	name := strings.Split(identifier, "#")
	id := name[0]
	fragment := name[1]

	if len(name) != 2 {
		return errors.New("invalid DID format, must contain both an identifier and a fragment and nothing else")
	}

	if !internal.IsValidIPNSName(id) {
		return fmt.Errorf("invalid DID format, identifier is not a valid IPNS name: %s", id)
	}

	if !internal.IsValidNanoID(name[1]) {
		return fmt.Errorf("invalid DID format, fragment is not a valid fragment: %s", fragment)
	}

	return nil
}

// AreIdentical checks if two DIDs have the same ID, ignoring the fragment.
// This is a stretched interpretation of the word Identical, but
// I couldn't help myself. It means they are derived from the same key,
// which makes them Identical in my book.
func AreIdentical(did1 *DID, did2 *DID) bool {
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

func IsValidIdentifier(identifier string) bool {

	parts := strings.Split(identifier, "#")
	if len(parts) != 2 {
		return false
	}

	// Check that the identifier has a valid fragment
	if !IsValidFragment(parts[1]) {
		return false
	}

	// Check that the id is a valid IPNS name
	_, err := ipns.NameFromString(identifier)

	// Last check so check that it has not errors
	return err == nil
}
