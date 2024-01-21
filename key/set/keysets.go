package set

import "fmt"

var keysets map[string]*Keyset

func init() {
	keysets = make(map[string]*Keyset)
}

func Add(keyset *Keyset) {
	keysets[keyset.id()] = keyset
}

func Get(id string) (*Keyset, error) {
	k, ok := keysets[id]
	if !ok {
		return nil, fmt.Errorf("keyset: keyset with id %s not found", id)
	}
	return k, nil
}

func Exists(id string) bool {
	_, ok := keysets[id]
	return ok
}

func Remove(id string) {
	delete(keysets, id)
}

func List() map[string]*Keyset {
	return keysets
}

func Clear() {
	keysets = make(map[string]*Keyset)
}

func GetByName(name string) (*Keyset, error) {
	for _, k := range keysets {
		if k.IPFSKey.Name == name {
			return k, nil
		}
	}
	return nil, fmt.Errorf("keyset: keyset with name %s not found", name)
}

func GetByID(id string) (*Keyset, error) {
	for _, k := range keysets {
		if k.IPFSKey.IPNSName == id {
			return k, nil
		}
	}
	return nil, fmt.Errorf("keyset: keyset with id %s not found", id)
}

// This id is the CID of the IPFS key, not the IPNS identifier.
// They should not be confused, hence this function is private.
func (k *Keyset) id() string {
	return k.IPFSKey.IPNSName
}
