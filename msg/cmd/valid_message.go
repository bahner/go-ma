package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/msg"
)

func main() {

	m := msg.ValidExampleMessage()
	packed, err := m.Pack()
	if err != nil {
		fmt.Printf("Error packing message: %s\n", err)
		os.Exit(70) // EX_SOFTWARE
	}

	json_message, _ := m.MarshalToCBOR()
	json_message_string := string(json_message)
	fmt.Println(json_message_string)
	fmt.Println(packed)

}
