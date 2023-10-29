package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/did/doc/vm"
	"github.com/bahner/go-ma/entity"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
	"github.com/bahner/go-ma/message/envelope"
	log "github.com/sirupsen/logrus"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	log.SetLevel(log.DebugLevel)

	os.Setenv("IPFS_API_SOCKET", "localhost:45005")

	// shell := internal.GetShell()

	// Create a new person, object - an entity
	// id, _ := nanoid.New()
	bahner, err := createSubject("bahner")
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}
	job, err := createSubject("job")
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}

	msgBody := "Share and enjoy!"
	msgMimeType := "text/plain"

	myMessage, err := message.New(
		bahner.DID.String(),
		job.DID.String(),
		msgBody,
		msgMimeType)
	if err != nil {
		fmt.Printf("Error creating new message: %v\n", err)
	}

	fmt.Println(myMessage)

	letter, err := envelope.Seal(myMessage)
	if err != nil {
		fmt.Printf("Error creating new envelope: %v\n", err)
	}
	messageJSON, err := letter.MarshalToJSON()
	if err != nil {
		fmt.Printf("Error marshalling message to JSON: %v\n", err)
	}

	fmt.Println(string(messageJSON))
}

func createSubject(name string) (*entity.Entity, error) {
	// Create a new person, object - an entity
	// id, _ := nanoid.New()

	ipnsKey, err := internal.IPNSGetOrCreateKey(name)
	if err != nil {
		return nil, fmt.Errorf("failed to get or create key in IPFS: %v", err)
	}

	subject, err := entity.NewFromKey(ipnsKey)
	if err != nil {
		return nil, fmt.Errorf("error creating new identity in ma: %v", err)
	}
	log.Debugf("Created new entity: %s", subject.DID.String())
	DIDDoc, err := doc.New(subject.DID.String(), subject.DID.String())
	if err != nil {
		return nil, fmt.Errorf("error creating new identity in ma: %v", err)
	}

	encVM, err := vm.New(subject.DID.Id, "enc1", subject.Keyset.EncryptionKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("error creating new verification method: %v", err)
	}
	log.Debugf("Created new verification method: %s", encVM.ID)

	err = DIDDoc.AddVerificationMethod(encVM)
	if err != nil {
		return nil, fmt.Errorf("error adding verification method: %v", err)
	}
	log.Debugf("Added verification method: %s", encVM.ID)
	signVM, err := vm.New(subject.DID.Id, "sign1", subject.Keyset.SignatureKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("error creating new verification method: %v", err)
	}
	log.Debugf("Created new verification method: %s", signVM.ID)

	err = DIDDoc.AddVerificationMethod(signVM)
	if err != nil {
		return nil, fmt.Errorf("error adding verification method: %v", err)
	}
	log.Debugf("Added verification method: %s", signVM.ID)

	err = DIDDoc.Sign(subject.Keyset.SignatureKey)
	if err != nil {
		return nil, fmt.Errorf("error signing new identity in ma: %v", err)
	}

	_, err = DIDDoc.Publish()
	if err != nil {
		return nil, fmt.Errorf("error publishing new identity in ma: %v", err)
	}

	return subject, nil
}
