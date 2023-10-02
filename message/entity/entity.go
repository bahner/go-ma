package entity

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	shell "github.com/ipfs/go-ipfs-api"
)

// This creates a New Live Identity for you. This is what you want to use,
// when you create new entitites.

// This function requires a live ipfs node to be running.

// So not only does it create a new DID, it also creates a new IPNS key, which
// you can use to publish your DID Document with.
type Entity struct {
	DID    *did.DID
	Keyset key.Keyset
}

// This creates a new Entity from an identifier.
// This is the base function for all the rest.

func New(name string) (*Entity, error) {

	if !internal.IsValidNanoID(name) {
		return nil, fmt.Errorf("entity: invalid name Try nanonid.New(): %s", name)
	}

	ipfsKey, err := internal.IPNSGetOrCreateKey(name)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to get or create key in IPFS: %v", err)
	}

	entityDID := did.NewFromIPNSKey(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create DID from ipnsKey: %s", err)
	}

	myKeyset, err := key.NewFromIPFSKey(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create key from ipnsKey: %s", err)
	}

	return &Entity{
		DID:    entityDID,
		Keyset: myKeyset,
	}, nil
}

func NewFromDID(d *did.DID) (*Entity, error) {

	return New(d.Fragment)
}

func NewFromKey(method string, ipfsKey *shell.Key) (*Entity, error) {

	return New(ipfsKey.Name)
}
