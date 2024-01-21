package ipfs

import (
	"fmt"

	"github.com/bahner/go-ma"
)

type Key struct {
	// The IPNS name of the key, not the local name
	IPNSName string
	// The ID used by kubo to identify the key
	DID string
	// Fragment is the local name of the key, which we use as the DID fragment
	Fragment string
}

// Get or create a key with the given name in the IPFS node.
func GetOrCreate(name string) (*Key, error) {

	// If cached simply return the key
	if exists(name) {
		return get(name)
	}

	return Fetch(name)
}

// Fetches the key from IPFS and updates the cache.
func Fetch(name string) (*Key, error) {
	// Get or create the key in IPFS
	ik, err := getOrCreateIPFSKey(name)
	if err != nil {
		return nil, fmt.Errorf("ipfs: failed to create key %s: %w", name, err)
	}

	ipnsName := ik.Path().Segments()[1]
	// Create a new Key struct
	k := &Key{
		DID:      ma.DID_PREFIX + ipnsName + "#" + name,
		Fragment: name,
		IPNSName: ipnsName,
	}

	// Add key to cache
	cache(k)

	return k, nil
}
