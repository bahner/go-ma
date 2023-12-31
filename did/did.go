package did

import (
	"fmt"
	"strings"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
)

type DID struct {
	// The identifier is the IPNS name and the fragment, as
	// provided as input to this function.
	Identifier string
	// The Fragment is the key shortname and internal name for the key
	Fragment string
	// Name is just Identifier#Fragment it's a convenience
	Name string
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
	identifier := internal.GetDIDIdentifier(name)
	fragment := internal.GetDIDFragment(name)

	return &DID{
		Identifier: identifier,
		Fragment:   fragment,
		Name:       name,
	}, nil
}

func (d *DID) String() string {

	return ma.DID_PREFIX + d.Name

}

// ValidateDID checks if the DID is valid.
func (d *DID) IsValid() bool {
	err := ValidateDID(d.String())
	return err == nil
}

func (d *DID) IsIdenticalTo(did DID) bool {

	return AreIdentical(d, &did)
}

func GetOrCreate(didStr string) (*DID, error) {

	// We don't care if the DID exists or not, we just want to create it.
	d, _ := Get(didStr)
	if d != nil {
		return d, nil
	}

	return New(didStr)

}
