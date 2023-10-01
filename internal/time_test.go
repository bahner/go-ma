package internal

import (
	"testing"
	"time"
)

func TestCreateTimeFromIsoString(t *testing.T) {
	validTime := "2022-01-01T01:01:01Z"
	expectedTime, _ := time.Parse(time.RFC3339, validTime)
	actualTime, err := CreateTimeFromIsoString(validTime)
	if err != nil {
		t.Fatalf("CreateTimeFromIsoString failed: %v", err)
	}
	if !actualTime.Equal(expectedTime) {
		t.Errorf("Expected %v, got %v", expectedTime, actualTime)
	}

	invalidTime := "not_a_time"
	_, err = CreateTimeFromIsoString(invalidTime)
	if err == nil {
		t.Errorf("CreateTimeFromIsoString succeeded for invalid time string")
	}
}
