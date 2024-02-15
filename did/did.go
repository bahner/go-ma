package did

import (
	"strings"

	"github.com/bahner/go-ma"
	"github.com/ipfs/boxo/path"
)

type DID struct {
	Identifier string
	// The identifier is the IPNS name without the /ipns/ prefix
	Fragment string
}

// This creates a new DID from an identifier.
// This is the base function for all the rest.
// The identitifier is the IPNS name and the fragment is the key shortname, eg
// did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner
//
// Remember that is needs to pre-exist in IPFS or be published to IPFS to be useful.
func New(didStr string) (DID, error) {

	// Firstly validate the DID
	err := ValidateDID(didStr)
	if err != nil {
		return DID{}, err
	}

	// Remove the prefix
	identifier := strings.TrimPrefix(didStr, ma.DID_PREFIX)
	fragment := GetFragment(identifier)

	return DID{
		Identifier: identifier,
		Fragment:   fragment,
	}, nil
}

// ValidateDID checks if the DID is valid.
func (d *DID) IsValid() bool {
	return d.Verify() == nil
}

func (d *DID) IsIdenticalTo(did DID) bool {

	return AreIdentical(*d, did)
}

func (d *DID) Path(space string) string {

	return "/" + path.IPNSNamespace + "/" + d.Identifier

}

func (d *DID) Verify() error {
	return ValidateDID(d.DID())
}

func (d *DID) Name() string {
	return d.Identifier + "#" + d.Fragment
}

func (d *DID) DID() string {
	return ma.DID_PREFIX + d.Name()
}

func (d *DID) String() string {
	return d.DID()
}
