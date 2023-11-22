package internal

import (
	"fmt"
	"time"
)

func CreateTimeFromIsoString(timestring string) (time.Time, error) {
	_time, err := time.Parse(time.RFC3339, timestring)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid timeformat: %w", err)
	}

	return _time, nil
}

// Returns the current time in UTC.
func Now() time.Time {
	return time.Now().UTC()
}

// Returns the current time in ISO 8601 format, fit for use in a DID document.
func NowIsoString() string {
	return Now().Format(time.RFC3339)
}

// Converts a time.Time to an ISO 8601 format string, fit for use in a DID document.
// No timezones, just Zulu time.
func TimeIsoString(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}
