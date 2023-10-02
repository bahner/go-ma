package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/message/entity"
	log "github.com/sirupsen/logrus"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	log.SetLevel(log.DebugLevel)

	os.Setenv("IPFS_API_HOST", "localhost:45005")

	// shell := internal.GetShell()

	// Create a new person, object - an entity
	// id, _ := nanoid.New()
	my, err := entity.New("space", "bahner")
	if err != nil {
		fmt.Printf("Error creating new identity in space: %v", err)
	}

	fmt.Printf("I am: %s\n", my.DID.String())

	// Create a new document

	myDoc, err := doc.New(my.DID.String())
	if err != nil {
		fmt.Printf("Error creating new document: %v", err)
	}

	encVM, err := vm.New(my.DID.Id, "enc-1", my.Keyset.EncryptionKey().PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating encryption verification method: %v", err)
	}
	err = myDoc.AddVerificationMethod(encVM)
	if err != nil {
		fmt.Printf("Error adding encryption verification method: %v", err)
	}

	sigVM, err := vm.New(my.DID.Id, "sig-1", my.Keyset.SignatureKey().PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating signature verification method: %v", err)
	}
	err = myDoc.AddVerificationMethod(sigVM)
	if err != nil {
		fmt.Printf("Error adding signature verification method: %v", err)
	}

	encVM, err = vm.New(my.DID.Id, "enc-1", my.Keyset.EncryptionKey().PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating encryption verification method: %v", err)
	}
	err = myDoc.AddVerificationMethod(encVM)
	if err != nil {
		fmt.Printf("Error adding encryption verification method: %v", err)
	}

	encVM, err = vm.New(my.DID.Id, "enc-2", my.Keyset.EncryptionKey().PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating encryption verification method: %v", err)
	}
	err = myDoc.AddVerificationMethod(encVM)
	if err != nil {
		fmt.Printf("Error adding encryption verification method: %v", err)
	}

	myDoc.Publish()

}
