package did

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
)

// Sometimes we just want the identifier, not the whole DID.
func ExtractID(didStr string) (string, error) {

	did, err := Parse(didStr)
	if err != nil {
		return "", internal.LogError(fmt.Sprintf("did: not a valid identifier: %v\n", err))
	}

	return did.Id, nil
}

func IsValidDID(didStr string) bool {

	_, err := Parse(didStr)
	return err == nil
}

// AreIdentical checks if two DIDs have the same ID, ignoring the fragment.
// This is a stretched interpretation of the word Identical, but
// I couldn't help myself. It means they are derived from the same key,
// which makes them Identical in my book.
func AreIdentical(did1 *DID, did2 *DID) bool {
	return did1.Id == did2.Id
}
