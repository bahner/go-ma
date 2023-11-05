package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/entity"
	keyset "github.com/bahner/go-ma/key/set"
)

func main() {

	fmt.Fprintln(os.Stderr, "******************************************************************")
	fmt.Fprintln(os.Stderr, "*The following strings contains secrets and should not be shared.*")
	fmt.Fprintln(os.Stderr, "*              It is only meant for testing.                     *")
	fmt.Fprintln(os.Stderr, "******************************************************************")

	var name string

	flag.StringVar(&name, "name", "", "Name of the entity to create")
	flag.Parse()

	// Create a new person, object - an entity
	ID, err := did.NewFromName(name)
	if err != nil {
		panic(err)
	}

	// Create a new keyset for the entity from the name (fragment)
	keyset, err := keyset.New(ID.Fragment)
	if err != nil {
		panic(err)
	}

	myEntity, err := entity.NewFromKeyset(keyset)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stderr, "Publishing Entity DIDDocument to IPFS. Please wait ...")
	err = myEntity.PublishEntityDocument()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, " done.")

	document, err := myEntity.Doc.JSON()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "Published the following Entity DIDDocument: ")
	fmt.Fprintln(os.Stderr, string(document))

	packedEntity, err := myEntity.Pack()
	if err != nil {
		panic(err)
	}

	fmt.Println(packedEntity)

}
