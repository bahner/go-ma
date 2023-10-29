package internal

import (
	"fmt"
	"time"
)

func CreateTimeFromIsoString(timestring string) (time.Time, error) {
	_time, err := time.Parse(time.RFC3339, timestring)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timeformat: %v", err)
	}

	return _time, nil
}

// Returns the current time in ISO 8601 format, fit for use in a DID document.
func NowIsoString() string {
	return time.Now().UTC().Format(time.RFC3339)
}
