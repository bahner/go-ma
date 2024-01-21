package entity

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	keyset "github.com/bahner/go-ma/key/set"
	log "github.com/sirupsen/logrus"
)

func CreateDocument(id *did.DID, controller *did.DID, keyset *keyset.Keyset) (*doc.Document, error) {
	// Initialize a new DID Document
	myDoc, err := doc.New(id.String(), id.String())
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create document: %s", err)
	}
	log.Debugf("entity: myDoc: %v", myDoc)

	// Add the encryption key to the document,
	// and set it as the key agreement key.
	myEncVM, err := doc.NewVerificationMethod(
		id.String(),
		id.String(),
		key.KEY_AGREEMENT_KEY_TYPE,
		internal.GetDIDFragment(keyset.EncryptionKey.DID),
		keyset.EncryptionKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create encryption verification method: %s", err)
	}
	myDoc.AddVerificationMethod(myEncVM)
	myDoc.KeyAgreement = myEncVM.ID
	log.Debugf("entity: myEncVM: %v", myDoc.KeyAgreement)

	// Add the signing key to the document and set it as the assertion method.
	mySignVM, err := doc.NewVerificationMethod(
		id.String(),
		id.String(),
		key.ASSERTION_METHOD_KEY_TYPE,
		internal.GetDIDFragment(keyset.SigningKey.DID),
		keyset.SigningKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create signing verification method: %s", err)
	}
	myDoc.AddVerificationMethod(mySignVM)
	myDoc.AssertionMethod = mySignVM.ID
	log.Debugf("entity: mySigVM: %v", myDoc.AssertionMethod)

	// Finally the document with the signing key.
	myDoc.Sign(keyset.SigningKey, mySignVM)

	return myDoc, nil
}

// Publish entity document. This needs to be done, when the keyset is new.
// Maybe we can check the assertionMethod and keyAgreement fields to see if
// the document is already published corretly.
func (e *Entity) PublishDocument() error {

	id, err := e.Doc.Publish(false)
	if err != nil {
		return fmt.Errorf("entity: failed to publish document: %s", err)
	}
	log.Debugf("entity: published document: %v to %s", e.Doc, id)
	return nil
}
