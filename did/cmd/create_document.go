package main

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/entity"
	"github.com/bahner/go-ma/internal"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)

	ipnsKey, err := internal.IPNSGetOrCreateKey("bahner")
	if err != nil {
		fmt.Printf("failed to get or create key in IPFS: %v\n", err)
	}

	id, err := did.NewFromIPNSKey(ipnsKey)
	if err != nil {
		fmt.Printf("failed to create new DID from IPNS key: %v\n", err)
	}

	i, err := entity.New(id, id)
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}

	fmt.Println(i.Doc.String())

}
