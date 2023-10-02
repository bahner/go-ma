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

	log.SetLevel(log.InfoLevel)

	os.Setenv("IPFS_API_HOST", "localhost:45005")

	// shell := internal.GetShell()

	// Create a new person, object - an entity
	// id, _ := nanoid.New()
	my, err := entity.New("bahner")
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}

	fmt.Printf("I am: %s\n", my.DID.String())

	// Create a new document

	myDoc, err := doc.New(my.DID.String())
	if err != nil {
		fmt.Printf("Error creating new document: %v\n", err)
	}

	fmt.Printf("Id: %s\n", myDoc.ID)
	fmt.Printf("Identifier: %s\n", my.DID.Identifier)

	encVM, err := vm.New(my.DID.Id, "enc-1", my.Keyset.EncryptionKey().PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating encryption verification method: %v\n", err)
	}
	err = myDoc.AddVerificationMethod(encVM)
	if err != nil {
		fmt.Printf("Error adding encryption verification method: %v\n", err)
	}

	sigVM, err := vm.New(my.DID.Id, "sig-1", my.Keyset.SignatureKey().PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating signature verification method: %v\n", err)
	}
	err = myDoc.AddVerificationMethod(sigVM)
	if err != nil {
		fmt.Printf("Error adding signature verification method: %v\n", err)
	}

	myDoc.Publish()

}
