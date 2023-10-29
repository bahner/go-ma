package internal

import (
	nanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateFragment() string {
	fragment, _ := nanoid.New()

	return "#" + fragment
}
