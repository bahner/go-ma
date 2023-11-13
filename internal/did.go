package internal

import (
	"strings"

	"github.com/bahner/go-ma"
)

func GetDIDFragment(did string) string {

	didName := strings.TrimPrefix(did, ma.DID_PREFIX)

	elements := strings.Split(didName, "#")

	return elements[len(elements)-1]
}

func GetDIDIdentifier(did string) string {

	didName := strings.TrimPrefix(did, ma.DID_PREFIX)

	elements := strings.Split(didName, "#")

	return elements[0]
}
