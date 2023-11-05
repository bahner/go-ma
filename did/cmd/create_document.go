package main

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/entity"
	ipnskey "github.com/bahner/go-ma/key/ipns"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.ErrorLevel)

	ipnsKey, err := ipnskey.New("bahner")

	id, err := did.NewFromIPNSKey(ipnsKey)
	if err != nil {
		fmt.Printf("failed to create new DID from IPNS key: %v\n", err)
	}

	i, err := entity.New(id, id)
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}

	docstring, err := i.Doc.MarshalPayloadToJSON()
	if err != nil {
		fmt.Printf("failed to get doc string from document, %v", err)
	}

	fmt.Println(string(docstring))

}
