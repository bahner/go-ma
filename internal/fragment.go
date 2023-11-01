package internal

import (
	"strings"

	nanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateFragment() string {
	fragment, _ := nanoid.New()

	return "#" + fragment
}

func GetFragmentFromDID(did string) string {
	elements := strings.Split(did, "#")

	return elements[len(elements)-1]
}
