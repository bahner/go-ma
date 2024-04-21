package utils

import (
	"regexp"
)

var (
	nanoId = regexp.MustCompile("^[a-zA-Z0-9_-]*$")
)

func ValidateNanoID(str string) bool {
	return nanoId.MatchString(str)
}
