package ipfs

import "fmt"

var keys map[string]*Key

func init() {
	keys = make(map[string]*Key)
}

func cache(k *Key) {
	keys[k.Fragment] = k
}

func get(id string) (*Key, error) {
	d, ok := keys[id]
	if !ok {
		return nil, fmt.Errorf("did: did with id %s not found", id)
	}
	return d, nil
}

func exists(fragment string) bool {
	_, ok := keys[fragment]
	return ok
}
