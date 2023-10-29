package entity

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
)

// This creates a New Live Identity for you. This is what you want to use,
// when you create new entitites.

// This function requires a live ipfs node to be running.

// So not only does it create a new DID, it also creates a new IPNS key, which
// you can use to publish your DID Document with.
type Entity struct {
	DID    *did.DID
	Doc    *doc.Document
	Keyset key.Keyset
}

// This creates a new Entity from an identifier.
// Set controller as the world DID or as self.

func New(id *did.DID, controller *did.DID) (*Entity, error) {

	ipfsKey, err := internal.IPNSGetOrCreateKey(id.Fragment) // The fragment is the key shortname
	if err != nil {
		return nil, fmt.Errorf("entity: failed to get or create key in IPFS: %v", err)
	}
	log.Debugf("entity: ipfsKey: %v", ipfsKey)

	myKeyset, err := key.NewKeysetFromIPFSKey(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create key from ipnsKey: %s", err)
	}
	log.Debugf("entity: myKeyset: %v", myKeyset)

	myDoc, err := doc.New(id.String(), id.String())
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create document: %s", err)
	}
	log.Debugf("entity: myDoc: %v", myDoc)

	myEncVM, err := vm.New(id.Id, "encryption", myKeyset.EncryptionKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create encryption verification method: %s", err)
	}
	log.Debugf("entity: myEncVM: %v", myEncVM)

	err = myDoc.AddVerificationMethod(myEncVM)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to add encryption verification method to document: %s", err)
	}

	mySigVM, err := vm.New(id.Id, "signature", myKeyset.SignatureKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create signature verification method: %s", err)
	}
	log.Debugf("entity: mySigVM: %v", mySigVM)

	err = myDoc.AddVerificationMethod(mySigVM)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to add signature verification method to document: %s", err)
	}

	return &Entity{
		DID:    id,
		Doc:    myDoc,
		Keyset: myKeyset,
	}, nil
}

func NewFromKey(method string, ipfsKey *shell.Key) (*Entity, error) {

	id, err := did.NewFromIPNSKey(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create did from ipnsKey: %s", err)
	}

	return New(id, id)
}
