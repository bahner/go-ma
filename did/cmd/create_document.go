package main

import (
	"fmt"

	"github.com/bahner/go-ma/entity"
	ipfskey "github.com/bahner/go-ma/key/ipfs"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.ErrorLevel)

	ipfsKey, err := ipfskey.Fetch("bahner")
	if err != nil {
		fmt.Printf("failed to create new IPNS key: %v\n", err)
	}

	_, err = entity.NewFromIPFSKey(ipfsKey)
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}

}
