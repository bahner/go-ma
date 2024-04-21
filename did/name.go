package did

import (
	"strings"

	"github.com/ipfs/boxo/ipns"
)

// Get the identifier from the DID string
// The prefix is not required, it'll just be stripped off.
func getName(did string) (ipns.Name, error) {

	didName := strings.TrimPrefix(did, PREFIX)

	elements := strings.Split(didName, "#")

	nameVal := elements[0]

	return ipns.NameFromString(nameVal)
}
