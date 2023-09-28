package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/coll"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/did/pkm"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	var err error
	os.Setenv("IPFS_API_HOST", "localhost:5001")

	shell := internal.GetShell()

	me, err := did.NewIdentity("space")
	if err != nil {
		fmt.Errorf("Error creating new identity in space: %v", err)

		myRSAKey, err := key.NewRSAKey()
	myVerificationMethod, err := vm.New(did.Id, "#key1", myPublicKeyMultibase)
	if err != nil {
		fmt.Println(err)
	}
	doc.AddVerificationMethod(myVerificationMethod)
	ctrl, err := coll.New(did.String())
	if err != nil {
		fmt.Println(err)
	}
	doc.Controller = ctrl

	doc.Sign(myKey)
	docString, err := doc.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", docString)

	// ok := doc.Verify()
	// if ok == nil {
	// 	fmt.Println("Signature verified")
	// } else {
	// 	fmt.Println("Signature verification failed")
	// }

	// os.Setenv("IPFS_API_SOCKET", "localhost:45005")

	// did, err := did.NewIdentity("space")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(did.String())

}
