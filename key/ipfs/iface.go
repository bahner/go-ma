package ipfs

import (
	"context"
	"fmt"

	"github.com/bahner/go-ma/api"
	iface "github.com/ipfs/kubo/core/coreiface"
	log "github.com/sirupsen/logrus"
)

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

	return lookedupKey, fmt.Errorf("lookupIPFSKey %s: %w", keyName, ErrKeyNotFoundInIPFS)
}
