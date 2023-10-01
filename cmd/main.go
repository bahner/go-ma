package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/key"
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

	env_key, err := key.GenerateEncryptionKey("x25519", "enc-1")
	if err != nil {
		fmt.Printf("Error getting encryption key: %v", err)
	}

	encVM, err := vm.New(my.DID.Id, env_key.Name(), env_key.PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating encryption verification method: %v", err)
	}
	err = myDoc.AddVerificationMethod(encVM)
	if err != nil {
		fmt.Printf("Error adding encryption verification method: %v", err)
	}

	env_key, err = key.GenerateEncryptionKey("x448", "enc-1")
	if err != nil {
		fmt.Printf("Error getting encryption key: %v", err)
	}

	encVM, err = vm.New(my.DID.Id, env_key.Name(), env_key.PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating encryption verification method: %v", err)
	}
	err = myDoc.AddVerificationMethod(encVM)
	if err != nil {
		fmt.Printf("Error adding encryption verification method: %v", err)
	}

	sigKey, err := key.GenerateSignatureKey("ed25519", "sig1")
	if err != nil {
		fmt.Printf("Error getting signature key: %v", err)
	}

	sigVM, err := vm.New(my.DID.Id, sigKey.Name(), sigKey.PublicKeyMultibase())
	if err != nil {
		fmt.Printf("Error creating signature verification method: %v", err)
	}
	err = myDoc.AddVerificationMethod(sigVM)
	if err != nil {
		fmt.Printf("Error adding signature verification method: %v", err)
	}

	myDoc.Publish()

}
