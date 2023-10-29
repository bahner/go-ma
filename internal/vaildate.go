package internal

import (
	"errors"
	"regexp"
	"strings"

	"github.com/ipfs/boxo/ipns"
	"github.com/multiformats/go-multibase"
)

var (
	ascii              = regexp.MustCompile("^[a-z]*$")
	alphanumeric       = regexp.MustCompile("^[a-z0-9]*$")
	nanoid             = regexp.MustCompile("^[a-zA-Z0-9_-]*$")
	fragment           = regexp.MustCompile("^#[a-zA-Z0-9_-]*$")
	ErrInvalidID       = errors.New("invalid ID")
	ErrInvalidFragment = errors.New("invalid fragment")
)

func IsAlnum(str string) bool {
	return alphanumeric.MatchString(str)
}

func IsValidMethod(method string) bool {
	return ascii.MatchString(method)
}

func IsValidMultibase(input string) bool {
	_, _, err := multibase.Decode(input)
	return err == nil
}

// ValidateNanoID checks if a string only contains valid NanoID characters.
func IsValidNanoID(str string) bool {
	return nanoid.MatchString(str)
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
