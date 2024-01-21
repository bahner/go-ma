package entity

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	keyset "github.com/bahner/go-ma/key/set"
	log "github.com/sirupsen/logrus"
)

func NewFromKeyset(keyset *keyset.Keyset) (*Entity, error) {

	id, err := did.New(keyset.DID.String())

	if err != nil {
		return nil, fmt.Errorf("entity: failed to create did from ipnsKey: %s", err)
	}

	myDoc, err := CreateDocument(id.String(), id.String(), keyset)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create document: %s", err)
	}

	log.Infof("Created new entity for %s", id.String())
	return &Entity{
		DID:    id,
		Doc:    myDoc,
		Keyset: keyset,
	}, nil
}

func NewFromPackedKeyset(data string) (*Entity, error) {

	keyset, err := keyset.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to unpack keyset: %s", err)
	}

	return NewFromKeyset(keyset)

}
