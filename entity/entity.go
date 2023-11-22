package entity

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
	ipnskey "github.com/bahner/go-ma/key/ipns"
	keyset "github.com/bahner/go-ma/key/set"
	cbor "github.com/fxamacker/cbor/v2"
	log "github.com/sirupsen/logrus"
)

// This creates a New Live Identity for you.
// This is what you want to use, when you create new entitites.

// This function requires a live ipfs node to be running.

// So not only does it create a new DID, it also creates a new IPNS key, which
// you can use to publish your DID Document with.
type Entity struct {
	DID    *did.DID
	Doc    *doc.Document
	Keyset *keyset.Keyset
}

// This creates a new Entity from an identifier.
// Set controller as the world DID or as self.

func New(id *did.DID, controller *did.DID, forceUpdate bool) (*Entity, error) {

	// Now we create a keyset for the entity.
	// The keyset creation will lookup the IPNS key again and also
	// create keys for signing and encryption.
	myKeyset, err := keyset.New(id.Fragment, forceUpdate)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create key from ipnsKey: %s", err)
	}
	log.Debugf("entity: myKeyset: %v", myKeyset)

	myDoc, err := CreateEntityDocument(id, controller, myKeyset)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create document: %s", err)
	}

	return &Entity{
		DID:    id,
		Doc:    myDoc,
		Keyset: myKeyset,
	}, nil
}

func NewFromKey(ipfsKey *ipnskey.Key, controller *did.DID, forceUpdate bool) (*Entity, error) {

	id, err := did.NewFromIPNSKey(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create did from ipnsKey: %s", err)
	}

	return New(id, controller, forceUpdate)
}

func (e *Entity) MarshalToCBOR() ([]byte, error) {
	data, err := cbor.Marshal(e.Keyset)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to publish document: %s", err)
	}
	return data, nil
}

func (e *Entity) UnmarshalFromCBOR(data []byte) error {
	err := cbor.Unmarshal(data, &e.Keyset)
	if err != nil {
		return fmt.Errorf("entity: failed to publish document: %s", err)
	}
	return nil
}

func (e *Entity) Pack() (string, error) {
	data, err := e.MarshalToCBOR()
	if err != nil {
		return "", fmt.Errorf("entity: failed to publish document: %s", err)
	}
	return internal.MultibaseEncode(data)
}

func Unpack(data string) (*Entity, error) {

	e := &Entity{}

	decodedData, err := internal.MultibaseDecode(data)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to publish document: %s", err)
	}
	err = cbor.Unmarshal(decodedData, e)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to unmarshal packed data: %s", err)
	}

	return e, nil
}

// Publishes the entity's DIDDocument and IPNSKey to IPFS.
// If force is true, it will publish even if the IPNSKey is already published.
func (e *Entity) Publish(force bool) error {

	// Publish the IPNSKey to IPFS for publication.
	err := e.Keyset.IPNSKey.ExportToIPFS(force)
	if err != nil {
		return fmt.Errorf("new_actor: Failed to export IPNSKey to IPFS: %w", err)
	}

	// Make sure the DIDDocument is published to IPFS if it's not already.
	_, err = e.Doc.Publish()
	if err != nil {
		return fmt.Errorf("new_actor: Failed to publish DOC: %w", err)
	}

	return nil
}
