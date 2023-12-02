package internal

import (
	"fmt"

	iface "github.com/ipfs/boxo/coreiface"
)

func IPFSListKeys() ([]iface.Key, error) {

	return GetIPSAPI().Key().List(GetContext())
}

func IPFSKeyLookupID(name string) (iface.Key, error) {

	keys, err := IPFSListKeys()
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

func IPFSKeyLookupName(keyName string) (iface.Key, error) {

	var lookedupKey iface.Key

	keys, err := IPFSListKeys()
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
