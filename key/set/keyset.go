package set

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/key/ipfs"
	log "github.com/sirupsen/logrus"
)

// KeySet struct the encryption and signing keys are actual keys,
// but the IPFSKey is a reference to the IPFS key and holds names and paths.
// The key itself resides in IPFS.
type Keyset struct {
	DID           *did.DID
	IPFSKey       *ipfs.Key
	EncryptionKey *key.EncryptionKey
	SigningKey    *key.SigningKey
}

// Creates new keyset from a name (typically fragment of a DID)
// This requires that the key is already in IPFS and that IPFS is running.
func GetOrCreate(name string) (*Keyset, error) {

	var err error

	if exists(name) {
		log.Debugf("Keyset %s found in cache", name)
		return get(name)
	}

	ipfsKey, err := ipfs.GetOrCreate(name)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to get or create key in IPFS: %w", err)
	}
	log.Infof("Created new key in IPFS: %v", ipfsKey)

	ks, err := newFromIPFSKey(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to create new keyset: %w", err)
	}

	// Add key to cache
	cache(ks)

	return ks, nil
}
