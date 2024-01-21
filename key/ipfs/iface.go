package ipfs

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	iface "github.com/ipfs/kubo/core/coreiface"
	log "github.com/sirupsen/logrus"
)

// Get or create a key with the given name in the IPFS node.
func getOrCreateIPFSKey(name string) (iface.Key, error) {

	var err error

	ik, _ := lookupIPFSKey(name)

	if ik == nil {
		ik, err = internal.GetIPFSAPI().Key().Generate(internal.GetContext(), name)
		if err != nil {
			return nil, fmt.Errorf("ipfs: failed to create key %s: %w", name, err)
		}
	}

	log.Debugf("ipfs: Created key %v", ik)

	return ik, nil
}

// List all keys in the IPFS node.
func listIPFSKeys() ([]iface.Key, error) {

	return internal.GetIPFSAPI().Key().List(internal.GetContext())
}

// Looks up a key by name in the IPFS node.
func lookupIPFSKey(keyName string) (iface.Key, error) {

	var lookedupKey iface.Key

	keys, err := listIPFSKeys()
	if err != nil {
		return lookedupKey, fmt.Errorf("ipfs: failed to list : %w", err)
	}

	// A little deeper than I usually like to nest, but hey, it's a one off.
	for _, foundKey := range keys {
		if foundKey.Name() == keyName {
			return foundKey, nil
		}
	}

	return lookedupKey, fmt.Errorf("ipfs: key %s not found", keyName)
}
