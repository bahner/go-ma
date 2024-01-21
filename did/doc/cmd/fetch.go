package main

import (
	"flag"
	"fmt"

	doc "github.com/bahner/go-ma/did/doc"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.ErrorLevel)

	id := flag.String("id", "", "IPNS name of the document to fetch")
	logLevel := flag.String("loglevel", "error", "Set the log level (debug, info, warn, error, fatal, panic)")

	flag.Parse()
	_level, err := log.ParseLevel(*logLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(_level)
	log.Debugf("main: log level set to %v", _level)

	// Create a new keyset for the entity
	doc, err := doc.Fetch(*id)
	if err != nil {
		panic(err)
	}
	log.Debugf("main: fetched document: %v", doc)

	fmt.Printf(doc.String())
}
