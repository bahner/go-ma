package did

import (
	"github.com/ipfs/boxo/path"
)

type DID struct {
	// The Id is the full DID
	Id string
	// The Identifier is the IPNS name without the /ipns/ prefix
	Identifier string
	// The fragment is a NanoID that is used to identify the entity locally
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
	err := Validate(didStr)
	if err != nil {
		return DID{}, err
	}

	return DID{
		Id:         didStr,
		Identifier: getIdentifier(didStr),
		Fragment:   getFragment(didStr),
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
	return Validate(d.Id)
}

func (d *DID) Name() string {
	return d.Identifier + "#" + d.Fragment
}
