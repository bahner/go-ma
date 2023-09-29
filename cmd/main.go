package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/did/pkm"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/message/entity"

	nanoid "github.com/matoous/go-nanoid/v2"
	log "github.com/sirupsen/logrus"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	log.SetLevel(log.DebugLevel)

	var err error
	os.Setenv("IPFS_API_HOST", "localhost:45005")

	// shell := internal.GetShell()

	// Create a new person, object - an entity
	id, _ := nanoid.New()
	me, err := entity.New("space", id)
	if err != nil {
		fmt.Printf("Error creating new identity in space: %v", err)
	}

	fmt.Printf("My DID: %s\n", me.DID.Identifier)

	// Create a new DID Document for the entity
	ddoc, err := doc.New(me.DID.String())
	if err != nil {
		fmt.Printf("Error creating new DID Document: %v", err)
	}

	// This is a little overengineering,
	// Put it's also not bad. We can add sugar functions
	// for verification methods, that will make it easier
	rsaPKM, err := pkm.New(me.Key.RSAPublicKey)
	if err != nil {
		fmt.Printf("Error parsing public key multibase: %v", err)
	}

	rsaVM, err := vm.New(me.DID.Id, "#rsa", rsaPKM)
	if err != nil {
		fmt.Printf("Error creating new Verification Method: %v", err)
	}

	signKey, err := pkm.New(me.Key.Ed25519PrivateKey)
	if err != nil {
		fmt.Printf("Error parsing public key multibase: %v", err)
	}

	signVM, err := vm.New(me.DID.Id, "#sign", signKey)
	if err != nil {
		fmt.Printf("Error creating new Verification Method: %v", err)
	}
	ddoc.AddVerificationMethod(rsaVM)
	ddoc.AddVerificationMethod(signVM)
	ddoc.Sign(me.Key.Ed25519PrivateKey)

	res, err := ddoc.Publish()
	if err != nil {
		fmt.Printf("Error publishing DID Document: %v", err)
	}

	fmt.Printf("Published DID Document: %s\n", res)

	// Lets see if we can retrieve it again

	retrieved_document, err := doc.Fetch(me.DID.String())
	if err != nil {
		fmt.Printf("Error retrieving DID Document: %v", err)
	}

	fmt.Printf("Retrieved DID Document: %s\n", retrieved_document)

}
