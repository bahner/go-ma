package did

import "fmt"

var dids map[string]*DID

func init() {
	dids = make(map[string]*DID)
}

func Add(did *DID) {
	dids[did.String()] = did
}

func Get(id string) (*DID, error) {
	d, ok := dids[id]
	if !ok {
		return nil, fmt.Errorf("did: did with id %s not found", id)
	}
	return d, nil
}

func GetByName(name string) (*DID, error) {
	for _, d := range dids {
		if d.Name == name {
			return d, nil
		}
	}
	return nil, fmt.Errorf("did: did with name %s not found", name)
}

func GetByID(identifier string) (*DID, error) {
	for _, d := range dids {
		if d.Identifier == identifier {
			return d, nil
		}
	}
	return nil, fmt.Errorf("did: did with id %s not found", identifier)
}

func Exists(id string) bool {
	_, ok := dids[id]
	return ok
}

func Remove(id string) {
	delete(dids, id)
}

func List() map[string]*DID {
	return dids
}
