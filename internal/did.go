package internal

import (
	"strings"

	"github.com/bahner/go-ma"
	nanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateFragment() string {
	fragment, _ := nanoid.New()

	return "#" + fragment
}

func GetFragmentFromDID(did string) string {

	didName := strings.TrimPrefix(did, ma.DID_PREFIX)

	elements := strings.Split(didName, "#")

	return elements[len(elements)-1]
}

func GetIdentifierFromDID(did string) string {

	didName := strings.TrimPrefix(did, ma.DID_PREFIX)

	elements := strings.Split(didName, "#")

	return elements[0]
}
