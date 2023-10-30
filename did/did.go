package did

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	nanoid "github.com/matoous/go-nanoid/v2"
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

	ipnsKey, err := internal.GetOrCreateIPNSKey(name)
	if err != nil {
		return nil, fmt.Errorf("did/new: failed to get or create key in IPFS: %w", err)
	}

	return NewFromIdentifier(ipnsKey.Id + "#" + ipnsKey.Name)

}

// This creates a new DID from an identifier.
// This is the base function for all the rest.
// The identitifier is the IPNS name and the fragment is the key shortname, eg
// k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner
// This implies that you have already created an IPNS key.
func NewFromIdentifier(name string) (*DID, error) {

	identifier, fragment, err := parseName(name)
	if err != nil {
		return &DID{}, fmt.Errorf("did/new: failed to parse identifier: %w", err)
	}

	return &DID{
		Identifier: identifier,
		Fragment:   fragment,
		Name:       name,
	}, nil
}

func NewFromName(name string) (*DID, error) {

	ipnsKey, err := internal.GetOrCreateIPNSKey(name)
	if err != nil {
		return &DID{}, fmt.Errorf("did/new: failed to parse identifier: %w", err)
	}

	return &DID{
		Identifier: ipnsKey.Id,
		Fragment:   name,
		Name:       ipnsKey.Id + "#" + name,
	}, nil
}

// If you already have a key, you can use this to create a DID.
func NewFromIPNSKey(keyName *shell.Key) (*DID, error) {

	new_id := keyName.Id + "#" + keyName.Name

	return NewFromIdentifier(new_id)

}

func Parse(didStr string) (*DID, error) {

	// Manually splitting by ":"
	// net/url doesn't handle this.
	parts := strings.Split(didStr, ":")
	if len(parts) < 3 {
		return &DID{}, errors.New("invalid DID format, missing method or ID")
	}

	scheme := parts[0]
	method := parts[1]
	name := parts[2]

	// Verify scheme
	if scheme != ma.DID_SCHEME {
		return &DID{}, fmt.Errorf("invalid DID format, scheme must be %s", ma.DID_SCHEME)
	}

	// Check the method is alphanumeric and 'ma'
	if !internal.IsAlnum(method) {
		return &DID{}, fmt.Errorf("invalid DID format, method must be alphanumeric: %s", method)
	}

	return NewFromIdentifier(name)
}

func parseName(identifier string) (string, string, error) {
	// Check if the identifier contains a fragment
	parts := strings.Split(identifier, "#")
	if len(parts) > 2 {
		return "", "", errors.New("invalid DID format, identifier contains more than one fragment")
	}

	id := parts[0]
	if !internal.IsValidMultibase(id) {
		return "", "", errors.New("invalid DID format, identifier is not a valid multibase string")
	}

	fragment := ""
	if len(parts) == 2 {
		fragment = parts[1]
	}

	if !internal.IsValidNanoID(fragment) {
		return "", "", errors.New("invalid DID format, fragment is not a valid fragment")
	}

	return id, fragment, nil
}

func (d *DID) String() string {

	return ma.DID_PREFIX + d.Identifier + "#" + d.Fragment

}

func (d *DID) IsValid() bool {
	_, err := Parse(d.Identifier)
	return err == nil
}

func (d *DID) PublicKey() (crypto.PubKey, error) {

	// Decode the PeerID from the ID which is the IPNS name
	pid, err := peer.Decode(d.Identifier)
	if err != nil {
		return nil, err
	}

	return peer.ID.ExtractPublicKey(pid)

}

func (d *DID) IsIdenticalTo(did DID) bool {

	return AreIdentical(d, &did)
}
