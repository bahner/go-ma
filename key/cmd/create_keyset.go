package main

import (
	"flag"
	"fmt"
	"os"

	keyset "github.com/bahner/go-ma/key/set"
	log "github.com/sirupsen/logrus"
)

func main() {

	fmt.Fprint(os.Stderr, "******************************************************************\n")
	fmt.Fprint(os.Stderr, "*The following strings contains secrets and should not be shared.*\n")
	fmt.Fprint(os.Stderr, "*              It is only meant for testing.                     *\n")
	fmt.Fprint(os.Stderr, "******************************************************************\n")

	log.SetLevel(log.ErrorLevel)

	name := flag.String("name", "", "Name of the entity to create")
	forceUpdate := flag.Bool("forceUpdate", false, "Force publish to IPFS")
	flag.Parse()

	// Create a new keyset for the entity
	keyset, err := keyset.New(*name, *forceUpdate)
	if err != nil {
		panic(err)
	}
	log.Debugf("main: keyset: %v", keyset)

	packedKeyset, err := keyset.Pack()
	if err != nil {
		panic(err)
	}
	fmt.Println(packedKeyset)

}
