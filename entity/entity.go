package entity

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/key"
	shell "github.com/ipfs/go-ipfs-api"
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
	Keyset key.Keyset
}

// This creates a new Entity from an identifier.
// Set controller as the world DID or as self.

func New(id *did.DID, controller *did.DID) (*Entity, error) {

	// Now we create a keyset for the entity.
	// The keyset creation will lookup the IPNS key again and also
	// create keys for signing and encryption.
	myKeyset, err := key.NewKeysetFromDID(id)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create key from ipnsKey: %s", err)
	}
	log.Debugf("entity: myKeyset: %v", myKeyset)

	// Initialize a new DID Document
	myDoc, err := doc.New(id.String(), id.String())
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create document: %s", err)
	}
	log.Debugf("entity: myDoc: %v", myDoc)

	// Add the encryption key to the document,
	// and set it as the key agreement key.
	myEncVM, err := doc.NewVerificationMethod(id.Identifier,
		id.String(),
		ma.KEY_AGREEMENT_KEY_TYPE,
		myKeyset.EncryptionKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create encryption verification method: %s", err)
	}
	myDoc.AddVerificationMethod(myEncVM)
	myDoc.KeyAgreement = myEncVM.ID
	log.Debugf("entity: myEncVM: %v", myDoc.KeyAgreement)

	// Add the signing key to the document and set it as the assertion method.
	mySignVM, err := doc.NewVerificationMethod(id.Identifier,
		id.String(),
		ma.VERIFICATION_KEY_TYPE,
		myKeyset.SigningKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create signing verification method: %s", err)
	}
	myDoc.AddVerificationMethod(mySignVM)
	myDoc.AssertionMethod = mySignVM.ID
	log.Debugf("entity: mySigVM: %v", myDoc.AssertionMethod)

	// Finally the document with the signing key.
	myDoc.Sign(myKeyset.SigningKey, mySignVM)

	return &Entity{
		DID:    id,
		Doc:    myDoc,
		Keyset: myKeyset,
	}, nil
}

func NewFromKey(ipfsKey *shell.Key) (*Entity, error) {

	id, err := did.NewFromIPNSKey(ipfsKey)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create did from ipnsKey: %s", err)
	}

	return New(id, id)
}
