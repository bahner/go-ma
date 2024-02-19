package ipfs

import (
	"fmt"

	"github.com/bahner/go-ma"
)

type Key struct {
	// The ID used by kubo to identify the key
	Id string
	// The IPNS name of the key, not the local name
	Identifier string
	// Fragment is the local name of the key, which we use as the DID fragment
	Fragment string
}

// Fetches the key from IPFS and updates the cache.
func GetOrCreate(name string) (Key, error) {
	// Get or create the key in IPFS
	ik, err := getOrCreateIPFSKey(name)
	if err != nil {
		return Key{}, fmt.Errorf("GetOrCreate: %w", err)
	}

	ipnsName := ik.Path().Segments()[1]
	// Create a new Key struct
	k := Key{
		Id:         ma.DID_PREFIX + ipnsName + "#" + name,
		Fragment:   name,
		Identifier: ipnsName,
	}

	return k, nil
}

func (k Key) Verify() error {
	if k.Identifier == "" {
		return ErrKeyMissingName
	}

	if k.Id == "" {
		return ErrKeyMissingID
	}

	if k.Fragment == "" {
		return ErrKeyMissingFragment
	}
	return nil
}

func (k Key) IsValid() bool {
	return k.Verify() == nil
}
