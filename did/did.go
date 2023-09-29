package did

import (
	"errors"
	"fmt"
	"strings"

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
func New(method string, identifier string) *DID {

	id, fragment, _ := parseIdentifier(identifier)

	return &DID{
		Scheme:     PREFIX,
		Method:     method,
		Id:         id,
		Identifier: identifier,
		Fragment:   fragment,
	}
}

// If you already have a key, you can use this to create a DID.
func NewFromIPNSKey(method string, keyName *shell.Key) *DID {

	new_id := keyName.Id + "#" + keyName.Name

	return New(method, new_id)

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
	if scheme != PREFIX {
		return &DID{}, errors.New("invalid DID format, missing 'did' scheme prefix")
	}

	// Check the method is alphanumeric and 'ma'
	if !internal.IsAlnum(method) {
		return &DID{}, errors.New("invalid DID format, method must be alphanumeric")
	}

	return New(method, identifier), nil
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

	// We could add fragment rules here. They're supposed to be a NanoID.

	return id, fragment, nil
}

func (d *DID) String() string {

	return d.Scheme + ":" + d.Method + ":" + d.Identifier

}

func (d *DID) IsValid() bool {
	_, err := Parse(d.Identifier)
	return err == nil
}

func (d *DID) IsIdenticalTo(did DID) bool {

	return d.Id == did.Id
}

// AreIdentical checks if two DIDs have the same ID, ignoring the fragment.
// This is a stretched interpretation of the word Identical, but
// I couldn't help myself. It means they are derived from the same key,
// which makes them Identical in my book.
func AreIdentical(did1 *DID, did2 *DID) bool {
	return did1.Id == did2.Id
}

func (d *DID) PublicKey() (crypto.PubKey, error) {

	// Decode the PeerID from the ID which is the IPNS name
	pid, err := peer.Decode(d.Id)
	if err != nil {
		return nil, err
	}

	return peer.ID.ExtractPublicKey(pid)

}

// Sometimes we just want the identifier, not the whole DID.
func ExtractID(didStr string) (string, error) {

	did, err := Parse(didStr)
	if err != nil {
		return "", internal.LogError(fmt.Sprintf("did: not a valid identifier: %v", err))
	}

	return did.Id, nil
}
