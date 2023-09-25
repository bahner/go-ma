package internal

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func LogError(msg string) error {
	log.Errorf(msg)
	return fmt.Errorf(msg)
}
