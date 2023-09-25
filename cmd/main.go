package main

import (
	"fmt"

	"github.com/bahner/go-ma/did"
)

var ()

func main() {

	did, err := did.NewIdentity("space")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(did.String())

}
