package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
	log "github.com/sirupsen/logrus"
)

func main() {

	fmt.Fprint(os.Stderr, "******************************************************************\n")
	fmt.Fprint(os.Stderr, "*The following strings contains secrets and should not be shared.*\n")
	fmt.Fprint(os.Stderr, "*              It is only meant for testing.                     *\n")
	fmt.Fprint(os.Stderr, "******************************************************************\n")

	var name string
	log.SetLevel(log.ErrorLevel)

	flag.StringVar(&name, "name", "", "Name of the entity to create")
	flag.Parse()

	// Create a new person, object - an entity
	ID, err := did.NewFromName(name)
	if err != nil {
		fmt.Printf("Error creating new DID: %v", err)
	}
	log.Debugf("main: ID: %v", ID)

	myID := ID.String()
	log.Debugf("main: myID: %s", myID)

	// Create a new keyset for the entity
	keyset, err := key.NewKeyset(ID.Fragment)
	if err != nil {
		panic(err)
	}
	log.Debugf("main: keyset: %v", keyset)

	packedKeyset, err := keyset.Pack()
	if err != nil {
		panic(err)
	}
	fmt.Println(packedKeyset)

	err = keyset.IPNSKey.ExportToIPFS(ID.Fragment)
	if err != nil {
		panic(err)
	}
}
