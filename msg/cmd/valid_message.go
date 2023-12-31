package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/bahner/go-ma/msg"
)

func main() {

	m := msg.ValidExampleMessage()
	packed, err := m.Pack()
	if err != nil {
		log.Fatalf("Error packing message: %s", err)
	}

	json_message, _ := m.MarshalToCBOR()
	json_message_string := string(json_message)
	fmt.Println(json_message_string)
	fmt.Println(packed)

}
