package internal

import (
	"testing"
)

func TestIsAlnum(t *testing.T) {
	if !IsAlnum("hello123") {
		t.Errorf("IsAlnum failed for alphanumeric string")
	}
	if IsAlnum("hello-123") {
		t.Errorf("IsAlnum succeeded for non-alphanumeric string")
	}
}

func TestIsValidMultibase(t *testing.T) {
	validMultibase := "k51qzi5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd"
	if !IsValidMultibase(validMultibase) {
		t.Errorf("IsValidMultibase failed for valid multibase string")
	}

	invalidMultibase := "i5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd"
	if IsValidMultibase(invalidMultibase) {
		t.Errorf("IsValidMultibase succeeded for invalid multibase string")
	}
}

func TestIsValidNanoID(t *testing.T) {
	if !IsValidNanoID("ABC123_-") {
		t.Errorf("IsValidNanoID failed for valid NanoID string")
	}
	if IsValidNanoID("ABC 123") {
		t.Errorf("IsValidNanoID succeeded for invalid NanoID string")
	}
}
