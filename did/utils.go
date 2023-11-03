package did

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"google.golang.org/appengine/log"
)

// // Sometimes we just want the identifier, not the whole DID.
// func ExtractID(didStr string) (string, error) {

// 	did, err := Parse(didStr)
// 	if err != nil {
// 		return "", internal.LogError(fmt.Sprintf("did: not a valid identifier: %v\n", err))
// 	}

// 	return did.Identifier, nil
// }

func IsValidDID(didStr string) bool {

	err := ValidateDID(didStr)
	return err == nil
}

func ValidateDID(didStr string) error {

	log.Debugf("did/validate: didStr: %s", didStr)
	if !strings.HasPrefix(didStr, ma.DID_PREFIX) {
		return fmt.Errorf("invalid DID format, must start with %s", ma.DID_PREFIX)
	}

	identifier := strings.TrimPrefix(didStr, ma.DID_PREFIX)
	log.Debugf("did/validate: identifier: %s", identifier)

	name := strings.Split(identifier, "#")
	id := name[0]
	log.Debugf("did/validate: id: %s", id)
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
