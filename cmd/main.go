package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/did/pkm"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/message/entity"

	log "github.com/sirupsen/logrus"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	log.SetLevel(log.InfoLevel)

	var err error
	os.Setenv("IPFS_API_HOST", "localhost:45005")

	// shell := internal.GetShell()

	// Create a new person, object - an entity
	me, err := entity.New("space", "me")
	if err != nil {
		fmt.Printf("Error creating new identity in space: %v", err)
	}

	fmt.Printf("My DID: %s\n", me.DID.Identifier)

	// Create a new DID Document for the entity
	ddoc, err := doc.New(me.DID.Id)
	if err != nil {
		fmt.Printf("Error creating new DID Document: %v", err)
	}

	// This is a little overengineering,
	// Put it's also not bad. We can add sugar functions
	// for verification methods, that will make it easier
	ddocPkm, err := pkm.Parse(me.Key.PublicKeyMultibase)
	if err != nil {
		fmt.Printf("Error parsing public key multibase: %v", err)
	}

	ddocVm, err := vm.New(me.DID.Id, "#key-1", ddocPkm)
	if err != nil {
		fmt.Printf("Error creating new Verification Method: %v", err)
	}

	ddoc.AddVerificationMethod(ddocVm)
	ddoc.Sign(me.Key)

	res, err := ddoc.Publish()
	if err != nil {
		fmt.Printf("Error publishing DID Document: %v", err)
	}

	fmt.Printf("Published DID Document: %s\n", res)

}
