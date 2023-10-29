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
)

type DID struct {
	Scheme     string
	Method     string
	Id         string
	Identifier string
	Fragment   string
}

// This creates a new DID from a method and an identifier.
// This is the base function for all the rest.
// The identifier is the IPNS name and the fragment is the key shortname, eg
// k51qzi5uqu5dj9807pbuod1pplf0vxh8m4lfy3ewl9qbm2s8dsf9ugdf9gedhr#bahner
func New(identifier string) (*DID, error) {

	id, fragment, err := parseIdentifier(identifier)
	if err != nil {
		return &DID{}, fmt.Errorf("did/new: failed to parse identifier: %w", err)
	}

	return &DID{
		Scheme:     ma.DID_SCHEME,
		Method:     ma.DID_METHOD,
		Id:         id,
		Identifier: identifier,
		Fragment:   fragment,
	}, nil
}

// If you already have a key, you can use this to create a DID.
func NewFromIPNSKey(keyName *shell.Key) (*DID, error) {

	new_id := keyName.Id + "#" + keyName.Name

	return New(new_id)

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
	identifier := parts[2]

	// Verify scheme
	if scheme != ma.DID_SCHEME {
		return &DID{}, fmt.Errorf("invalid DID format, scheme must be %s", ma.DID_SCHEME)
	}

	// Check the method is alphanumeric and 'ma'
	if !internal.IsAlnum(method) {
		return &DID{}, fmt.Errorf("invalid DID format, method must be alphanumeric: %s", method)
	}

	return New(identifier)
}

func parseIdentifier(identifier string) (string, string, error) {
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

	return d.Scheme + ":" + d.Method + ":" + d.Identifier

}

func (d *DID) IsValid() bool {
	_, err := Parse(d.Identifier)
	return err == nil
}

func (d *DID) PublicKey() (crypto.PubKey, error) {

	// Decode the PeerID from the ID which is the IPNS name
	pid, err := peer.Decode(d.Id)
	if err != nil {
		return nil, err
	}

	return peer.ID.ExtractPublicKey(pid)

}

func (d *DID) IsIdenticalTo(did DID) bool {

	return d.Id == did.Id
}
