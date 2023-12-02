package key

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	coreiface "github.com/ipfs/kubo/core/coreiface"
	log "github.com/sirupsen/logrus"
)

func LookupID(name string) (coreiface.Key, error) {

	keys, err := List()
	if err != nil {
		return nil, fmt.Errorf("ipfs: failed to list : %w", err)
	}
	for _, key := range keys {
		if key.Name() == name {
			return key, nil
		}
	}

	return nil, fmt.Errorf("ipfs: key %s not found", name)
}

func LookupName(keyName string) (coreiface.Key, error) {

	var lookedupKey coreiface.Key

	keys, err := List()
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

func KeyWithNameExists(name string) bool {

	_, err := LookupName(name)
	return err == nil
}

func KeyWithIdentifierExists(identifier string) bool {

	_, err := LookupID(identifier)
	return err == nil

}

// Get or create a key with the given name in the IPFS node.
func GetOrCreate(name string) (coreiface.Key, error) {

	var err error

	ik, _ := LookupName(name)

	if ik == nil {
		ik, err = internal.GetIPFSAPI().Key().Generate(internal.GetContext(), name)
		if err != nil {
			return nil, fmt.Errorf("ipfs: failed to create key %s: %w", name, err)
		}
	}

	log.Debugf("ipfs: Created key %v", ik)

	return ik, nil
}
