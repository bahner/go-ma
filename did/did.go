package did

import (
	"fmt"
	"strings"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	ipnskey "github.com/bahner/go-ma/key/ipns"
	nanoid "github.com/matoous/go-nanoid/v2"
	log "github.com/sirupsen/logrus"
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

func New() (*DID, error) {

	name, err := nanoid.New()
	if err != nil {
		return nil, fmt.Errorf("did/new: error generating nanoid: %w", err)
	}
	log.Debugf("did/new: name: %s", name)

	ipnsKey, err := ipnskey.New(name)
	if err != nil {
		return nil, fmt.Errorf("did/new: failed to get or create key in IPFS: %w", err)
	}
	log.Debugf("did/new: ipnsKey.DID: %s", ipnsKey.DID)

	return NewFromDID(ipnsKey.DID)

}

// This creates a new DID from an identifier.
// This is the base function for all the rest.
// The identitifier is the IPNS name and the fragment is the key shortname, eg
// did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner
//
// Remember that is needs to pre-exist in IPFS or be published to IPFS to be useful.
func NewFromDID(didStr string) (*DID, error) {

	// Firstly validate the DID
	err := ValidateDID(didStr)
	if err != nil {
		return &DID{}, fmt.Errorf("did/new: failed to validate DID: %w", err)
	}

	// Remove the prefix
	name := strings.TrimPrefix(didStr, ma.DID_PREFIX)

	// Extract the identifier and fragment
	identifier := internal.GetIdentifierFromDID(name)
	fragment := internal.GetFragmentFromDID(name)

	return &DID{
		Identifier: identifier,
		Fragment:   fragment,
		Name:       name,
	}, nil
}

// NewFromName creates a new DID from a name, where name is just the fragment.
// The Identifier will be generated.
func NewFromName(name string) (*DID, error) {

	log.Debugf("did/new: name: %s", name)
	ipnsKey, err := ipnskey.New(name)
	if err != nil {
		return &DID{}, fmt.Errorf("did/new: failed to parse identifier: %w", err)
	}

	identifier := internal.GetIdentifierFromDID(ipnsKey.DID)
	log.Debugf("did/new: identifier: %s", identifier)
	fragment := internal.GetFragmentFromDID(ipnsKey.DID)
	log.Debugf("did/new: fragment: %s", fragment)
	log.Debugf("did/new: ipnsKey.DID: %s", ipnsKey.DID)

	return &DID{
		Identifier: identifier,
		Fragment:   fragment,
		Name:       ipnsKey.DID,
	}, nil
}

// If you already have a key, you can use this to create a DID.
func NewFromIPNSKey(keyName ipnskey.Key) (*DID, error) {

	return NewFromDID(keyName.DID)

}

func (d *DID) String() string {

	return d.Name

}

// ValidateDID checks if the DID is valid.
func (d *DID) IsValid() bool {
	err := ValidateDID(d.String())
	return err == nil
}

func (d *DID) IsIdenticalTo(did DID) bool {

	return AreIdentical(d, &did)
}
