package entity

import (
	"context"
	"fmt"
	"sync"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	keyset "github.com/bahner/go-ma/key/set"
	log "github.com/sirupsen/logrus"
)

func CreateDocument(id string, controller string, keyset *keyset.Keyset) (*doc.Document, error) {
	// Initialize a new DID Document
	myDoc, err := doc.GetOrCreate(id, controller)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create document: %s", err)
	}

	// Add the encryption key to the document,
	// and set it as the key agreement key.
	log.Debugf("entity/document: existing keyAgreement: %v", myDoc.KeyAgreement)
	myEncVM, err := doc.NewVerificationMethod(
		id,
		id,
		key.KEY_AGREEMENT_KEY_TYPE,
		internal.GetDIDFragment(keyset.EncryptionKey.DID),
		keyset.EncryptionKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("entity/document: failed to create encryption verification method: %s", err)
	}
	// Add the controller to the verification method
	myEncVM.AddController(controller)

	// Set the key agreement key verification method
	myDoc.AddVerificationMethod(myEncVM)
	myDoc.KeyAgreement = myEncVM.ID
	log.Debugf("entity/document: set keyAgreement to %v for %s", myDoc.KeyAgreement, myDoc.ID)

	// Add the signing key to the document and set it as the assertion method.
	log.Debugf("entity/document: Creating assertionMethod for document %s", myDoc.ID)
	mySignVM, err := doc.NewVerificationMethod(
		id,
		id,
		key.ASSERTION_METHOD_KEY_TYPE,
		internal.GetDIDFragment(keyset.SigningKey.DID),
		keyset.SigningKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("entity: failed to create signing verification method: %s", err)
	}
	// Add the controller to the verification method
	mySignVM.AddController(controller)

	// Set the assertion method verification method
	myDoc.AddVerificationMethod(mySignVM)
	myDoc.AssertionMethod = mySignVM.ID
	log.Debugf("entity/document: Set assertionMethod to %v for %s", myDoc.AssertionMethod, mySignVM.ID)
	// Finally the document with the signing key.
	myDoc.Sign(keyset.SigningKey, mySignVM)

	return myDoc, nil
}

// Publish entity document. This needs to be done, when the keyset is new.
// Maybe we can check the assertionMethod and keyAgreement fields to see if
// the document is already published corretly.
func (e *Entity) PublishDocument() error {

	id, err := e.Doc.Publish(nil)
	if err != nil {
		return fmt.Errorf("entity: failed to publish document: %s", err)
	}
	log.Debugf("entity: published document: %v to %s", e.Doc, id)
	return nil
}

func (e *Entity) PublishDocumentGorutine(wg *sync.WaitGroup, cancel context.CancelFunc, opts *doc.PublishOptions) {
	defer wg.Done()
	defer cancel() // Ensure the context is canceled once this function returns

	// Launch the Publish operation in a separate goroutine
	done := make(chan struct{})
	go func() {
		e.Doc.Publish(opts) // Assuming Publish handles the context internally
		close(done)
	}()

	// Wait for the Publish operation to complete or for the context to be cancelled/timed out
	select {
	case <-opts.Ctx.Done():
		// Context is cancelled or timed out
		if opts.Ctx.Err() == context.DeadlineExceeded {
			log.Errorf("entity: deadline exceeded: %v", opts.Ctx.Err())
		} else {
			log.Errorf("entity: context cancelled: %v", opts.Ctx.Err())
		}
	case <-done:
		// Publish operation completed
		log.Infof("Published document for entity: %s", e.DID.String())
	}
}
