package did

import (
	"fmt"
	"strings"

	"github.com/bahner/go-ma"
)

type DID struct {
	// The identifier is the IPNS name without the /ipns/ prefix
	Identifier string
	// The Fragment is the key shortname and internal name for the key
	Fragment string
}

// This creates a new DID from an identifier.
// This is the base function for all the rest.
// The identitifier is the IPNS name and the fragment is the key shortname, eg
// did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner
//
// Remember that is needs to pre-exist in IPFS or be published to IPFS to be useful.
func New(didStr string) (*DID, error) {

	// Firstly validate the DID
	err := ValidateDID(didStr)
	if err != nil {
		return &DID{}, fmt.Errorf("did/new: failed to validate DID: %w", err)
	}

	// Remove the prefix
	name := strings.TrimPrefix(didStr, ma.DID_PREFIX)

	// Extract the identifier and fragment
	identifier := GetIdentifier(name)
	fragment := GetFragment(name)

	return &DID{
		Identifier: identifier,
		Fragment:   fragment,
	}, nil
}

func (d *DID) String() string {

	return ma.DID_PREFIX + d.Identifier + "#" + d.Fragment

}

// ValidateDID checks if the DID is valid.
func (d *DID) IsValid() bool {
	err := ValidateDID(d.String())
	return err == nil
}

func (d *DID) IsIdenticalTo(did DID) bool {

	return AreIdentical(d, &did)
}

func (d *DID) Path(space string) string {

	return "/" + space + "/" + d.Identifier

}
