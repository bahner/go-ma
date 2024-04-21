package did

import (
	"fmt"
	"regexp"
	"strings"
)

var fragment = regexp.MustCompile("^#[a-zA-Z0-9_-]*$")

// Get the fragment from the DID string
// The prefix is not required, ut'll just be stripped off.
func getFragment(did string) string {

	didName := strings.TrimPrefix(did, PREFIX)

	elements := strings.Split(didName, "#")

	return elements[len(elements)-1]
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
