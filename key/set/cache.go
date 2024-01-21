package set

import "fmt"

var keysets map[string]*Keyset

func init() {
	keysets = make(map[string]*Keyset)
}

func cache(keyset *Keyset) {
	keysets[keyset.id()] = keyset
}

func get(id string) (*Keyset, error) {
	k, ok := keysets[id]
	if !ok {
		return nil, fmt.Errorf("keyset with id %s not found in cache", id)
	}
	return k, nil
}

func exists(id string) bool {
	_, ok := keysets[id]
	return ok
}

// This id is the CID of the IPFS key, not the IPNS identifier.
// They should not be confused, hence this function is private.
func (k *Keyset) id() string {
	return k.IPFSKey.Fragment
}
