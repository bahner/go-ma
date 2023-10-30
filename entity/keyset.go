package entity

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
)

func NewFromKeyset(keyset key.Keyset) (*Entity, error) {

	id, err := did.NewFromIPNSKey(keyset.IPNSKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create did from ipnsKey: %s", err)
	}

	myDoc, err := CreateEntityDocument(id, id, keyset)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create document: %s", err)
	}

	return &Entity{
		DID:    id,
		Doc:    myDoc,
		Keyset: keyset,
	}, nil
}

func NewFromPackedKeyset(data string) (*Entity, error) {

	keyset, err := key.UnpackKeyset(data)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to unpack keyset: %s", err)
	}

	return NewFromKeyset(keyset)

}
