package main

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/did/coll"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/did/pubkey"
	"github.com/bahner/go-ma/did/vm"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	did := did.New("space", "bahner")
	doc, _ := doc.New(did.String())
	myIPNSKey, err := internal.IPNSGetOrCreateKey("bahner")
	if err != nil {
		fmt.Println(err)
	}
	myKey, err := key.New(myIPNSKey)
	if err != nil {
		fmt.Println(err)
	}
	myPublicKeyMultibase, err := pubkey.New(myKey.RSAPrivateKey)
	if err != nil {
		fmt.Println(err)
	}
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
	cid, _ := internal.IPFSPublishString(docString)

	data, _ := internal.IPNSPublishCID(cid, "bahner", true)

	fmt.Printf("IPNS: %s\n", data)

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
