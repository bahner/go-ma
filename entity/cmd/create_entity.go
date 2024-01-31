package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/entity"
	keyset "github.com/bahner/go-ma/key/set"
	log "github.com/sirupsen/logrus"
)

func main() {

	fmt.Fprintln(os.Stderr, "********************************************************************")
	fmt.Fprintln(os.Stderr, "* The following strings contains secrets and should not be shared. *")
	fmt.Fprintln(os.Stderr, "*               It is only meant for testing.                      *")
	fmt.Fprintln(os.Stderr, "********************************************************************")

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

	// Set options for proper timeout of IPNS publish
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(600)*time.Second)
	opts := doc.DefaultPublishOptions()
	opts.Ctx = ctx

	// Setup a waitgroup to wait for the goroutine to finish or be cancelled
	var wg sync.WaitGroup
	wg.Add(1)

	go myEntity.PublishDocumentGorutine(&wg, cancel, opts)
	if err != nil {
		panic(err)
	}
	wg.Wait()

	packedEntity, err := myEntity.Pack()
	if err != nil {
		panic(err)
	}

	fmt.Println(packedEntity)

}
