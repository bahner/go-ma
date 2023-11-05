package entity

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/doc"
	keyset "github.com/bahner/go-ma/key/set"
	log "github.com/sirupsen/logrus"
)

func CreateEntityDocument(id *did.DID, controller *did.DID, keyset keyset.Keyset) (*doc.Document, error) {
	// Initialize a new DID Document
	myDoc, err := doc.New(id.String(), id.String())
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create document: %s", err)
	}
	log.Debugf("entity: myDoc: %v", myDoc)

	// Add the encryption key to the document,
	// and set it as the key agreement key.
	myEncVM, err := doc.NewVerificationMethod(
		id.Identifier,
		id.String(),
		ma.KEY_AGREEMENT_KEY_TYPE,
		keyset.EncryptionKey.PublicKeyMultibase)
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

func (e *Entity) PublishEntityDocument() error {
	_, err := e.Doc.Publish()
	if err != nil {
		return fmt.Errorf("entity: failed to publish document: %s", err)
	}
	return nil
}
