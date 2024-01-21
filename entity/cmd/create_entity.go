package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bahner/go-ma/entity"
	keyset "github.com/bahner/go-ma/key/set"
	log "github.com/sirupsen/logrus"
)

func main() {

	fmt.Fprintln(os.Stderr, "******************************************************************")
	fmt.Fprintln(os.Stderr, "*The following strings contains secrets and should not be shared.*")
	fmt.Fprintln(os.Stderr, "*              It is only meant for testing.                     *")
	fmt.Fprintln(os.Stderr, "******************************************************************")

	name := flag.String("name", "", "Name of the entity to create")
	logLevel := flag.String("loglevel", "error", "Set the log level (debug, info, warn, error, fatal, panic)")

	flag.Parse()
	_level, err := log.ParseLevel(*logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(_level)
	log.Debugf("main: log level set to %v", _level)
	// Create a new keyset for the entity from the name (fragment)
	keyset, err := keyset.GetOrCreate(*name)
	if err != nil {
		panic(err)
	}

	myEntity, err := entity.NewFromKeyset(keyset)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stderr, "Publishing Entity DIDDocument to IPFS. Please wait ...")
	err = myEntity.PublishDocument()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, " done.")

	packedEntity, err := myEntity.Pack()
	if err != nil {
		panic(err)
	}

	fmt.Println(packedEntity)

}
