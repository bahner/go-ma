package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
)

func main() {

	fmt.Fprint(os.Stderr, "******************************************************************\n")
	fmt.Fprint(os.Stderr, "*The following strings contains secrets and should not be shared.*\n")
	fmt.Fprint(os.Stderr, "*              It is only meant for testing.                     *\n")
	fmt.Fprint(os.Stderr, "******************************************************************\n")

	var name string

	flag.StringVar(&name, "name", "", "Name of the entity to create")
	flag.Parse()

	// Create a new person, object - an entity
	ID, err := did.NewFromName(name)
	if err != nil {
		panic(err)
	}

	// Create a new keyset for the entity
	keyset, err := key.NewKeyset(ID)
	if err != nil {
		panic(err)
	}

	packedKeyset, err := keyset.Pack()
	if err != nil {
		panic(err)
	}
	fmt.Println(packedKeyset)

}
