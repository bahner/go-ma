package doc

import "fmt"

var docs map[string]*Document

func init() {
	docs = make(map[string]*Document)
}

func cache(d *Document) {
	docs[d.ID] = d
}

func get(id string) (*Document, error) {
	d, ok := docs[id]
	if !ok {
		return nil, fmt.Errorf("did: did with id %s not found", id)
	}
	return d, nil
}

func exists(fragment string) bool {
	_, ok := docs[fragment]
	return ok
}
