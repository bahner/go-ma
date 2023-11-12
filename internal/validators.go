package internal

import (
	"crypto/ed25519"
	"regexp"

	"github.com/ipfs/boxo/ipns"
)

var (
	alphanumeric = regexp.MustCompile("^[a-z0-9]*$")
	nanoId       = regexp.MustCompile("^[a-zA-Z0-9_-]*$")
)

func IsAlnum(str string) bool {
	return alphanumeric.MatchString(str)
}

// ValidateNanoID checks if a string only contains valid NanoID characters.
func IsValidNanoID(str string) bool {
	return nanoId.MatchString(str)
}

func IsValidEd25519PrivateKey(privKey *ed25519.PrivateKey) bool {
	if privKey == nil || len(*privKey) != ed25519.PrivateKeySize {
		return false
	}
	return true
}

func IsValidIPNSName(name string) bool {
	_, err := ipns.NameFromString(name)
	return err == nil
}
