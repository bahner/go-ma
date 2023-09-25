package internal

import (
	"errors"
	"regexp"

	"github.com/multiformats/go-multibase"
)

var (
	ascii        = regexp.MustCompile("^[a-z]*$")
	alphanumeric = regexp.MustCompile("^[a-z0-9]*$")
	nanoAlphabet = regexp.MustCompile("^[ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-]*$")
	// Accept must fewer than normal characters DID fragments.
	// These are used to name keys in verification methods.
	fragment           = regexp.MustCompile("^#[a-zA-Z0-9-]+$")
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
	return nanoAlphabet.MatchString(str)
}

func IsValidFragment(str string) bool {
	return fragment.MatchString(str)
}
