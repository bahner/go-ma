package did

import (
	"github.com/ipfs/boxo/ipns"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

const PREFIX = "did:ma:"

type DID struct {
	// The Id is the full DID
	Id string
	// The Name is the IPNS name of the entity
	Name ipns.Name
	// The fragment is a NanoID that is used to identify the entity locally
	Fragment string
}

// New creates a new DID from a Name and a fragment.
// This is the recommended way to create a new DID.
func New(name ipns.Name, fragment string) DID {

	return DID{
		Id:       PREFIX + name.String() + "#" + fragment,
		Name:     name,
		Fragment: fragment,
	}
}

// This creates a new DID from an identifier.
// This is the base function for all the rest.
// The identitifier is the IPNS name and the fragment is the key shortname, eg
// did:ma:k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner
//
// Remember that is needs to pre-exist in IPFS or be published to IPFS to be useful.
func NewFromString(didStr string) (DID, error) {

	// Firstly validate the DID
	name, fragment, err := Parse(didStr)
	if err != nil {
		return DID{}, err
	}

	return New(name, fragment), nil
}

func NewFromPrivateKey(key crypto.PrivKey, fragment string) (DID, error) {

	p, err := peer.IDFromPrivateKey(key)
	if err != nil {
		return DID{}, err
	}

	return New(ipns.NameFromPeer(p), fragment), nil
}

func (d *DID) Validate() error {
	return Validate(d.Id)
}
