package entity

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/keyset"
	shell "github.com/ipfs/go-ipfs-api"
)

// This creates a New Live Identity for you. This is what you want to use,
// when you create new entitites.

// This function requires a live ipfs node to be running.

// So not only does it create a new DID, it also creates a new IPNS key, which
// you can use to publish your DID Document with.
type Entity struct {
	DID    *did.DID
	Keyset *keyset.Keyset
}

// This creates a new Entity from a method and an identifier.
// This is the base function for all the rest.

func New(method string, name string) (*Entity, error) {

	if !internal.IsValidNanoID(name) {
		return nil, fmt.Errorf("entity: invalid name Try nanonid.New(): %s", name)
	}

	ipfsKey, err := internal.IPNSGetOrCreateKey(name)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to get or create key: %v", err)
	}

	entityDID := did.NewFromIPNSKey(method, ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create DID from ipnsKey: %s", err)
	}

	maKey, err := keyset.New(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create key from ipnsKey: %s", err)
	}

	return &Entity{
		DID:    entityDID,
		Keyset: maKey,
	}, nil
}

func NewFromDID(d *did.DID) (*Entity, error) {

	ipfsKey, err := internal.IPNSGetOrCreateKey(d.Fragment)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to get or create key: %v", err)
	}

	maKey, err := keyset.New(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create key from ipnsKey: %s", err)
	}

	return &Entity{
		DID:    d,
		Keyset: maKey,
	}, nil
}

func NewFromKey(method string, a *shell.Key) (*Entity, error) {

	k, err := keyset.New(a)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create key from ipnsKey: %s", err)
	}

	d := did.NewFromIPNSKey(method, a)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create DID from ipnsKey: %s", err)
	}

	return &Entity{
		DID:    d,
		Keyset: k,
	}, nil
}
