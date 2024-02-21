package did

// This file contains the functions to create and fetch DIDs from IPFS.
// So this is a "live" part of the DID system, as opposed to the "static" part

import (
	"context"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/api"
	iface "github.com/ipfs/kubo/core/coreiface"
	log "github.com/sirupsen/logrus"
)

// Creates a new DID my actually creating or fetching a new key in IPFS
// And get the names - notably the IPNS name from there.
func GetOrCreate(name string) (DID, error) {
	// Get or create the key in IPFS
	ik, err := getOrCreateIPFSKey(name)
	if err != nil {
		return DID{}, fmt.Errorf("GetOrCreate: %w", err)
	}

	ipnsName := ik.Path().Segments()[1]

	return New(ma.DID_PREFIX + ipnsName + "#" + name)
}

// Get or create a key with the given name in the IPFS node.
func getOrCreateIPFSKey(name string) (iface.Key, error) {

	var err error

	ik, err := lookupIPFSKey(name)
	if err != nil {
		log.Debugf("getOrCreateIPFSKey: Ignoring: %v", err)
	}

	if ik == nil {
		ik, err = api.GetIPFSAPI().Key().Generate(context.Background(), name)
		if err != nil {
			return nil, fmt.Errorf("getOrCreateIPFSKey %s: %w", name, err)
		}
	}

	log.Debugf("getOrCreateIPFSKey: Created %v", ik)

	return ik, nil
}

// Looks up a key by name in the IPFS node.
func lookupIPFSKey(keyName string) (iface.Key, error) {

	var lookedupKey iface.Key

	keys, err := api.GetIPFSAPI().Key().List(context.Background())
	if err != nil {
		return lookedupKey, fmt.Errorf("lookupIPFSKey : %w", err)
	}

	// A little deeper than I usually like to nest, but hey, it's a one off.
	for _, foundKey := range keys {
		if foundKey.Name() == keyName {
			return foundKey, nil
		}
	}

	return lookedupKey, fmt.Errorf("lookupIPFSKey %s: %w", keyName, ErrIPFSKeyNotFound)
}
