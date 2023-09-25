package internal_test

import (
	"testing"
	"time"

	"github.com/bahner/go-ma/internal"
)

// func TestEncode(t *testing.T) {
// 	input := []byte("hello")
// 	expectedEncoding := multibase.Base58BTC // Replace with the expected multibase encoding for your MESSAGE_SIGNATURE_ENCODING
// 	encoded_result, _ := multibase.Encode(multibase.Encoding(expectedEncoding), input)
// 	expected_encoding, expected_result, _ := multibase.Encode(expected)
// 	actual, err := internal.Encode(input)
// 	if err != nil {
// 		t.Fatalf("Encode failed: %v", err)
// 	}
// 	if actual != expected {
// 		t.Errorf("Expected %s, got %s", expected, actual)
// 	}

// 	if expected_encoding != expectedEncoding {
// 		t.Errorf("Expected %v, got %v", expectedEncoding, expected_encoding)
// 	}
// }

// func TestDecode(t *testing.T) {
// 	encoded := "base58_encoded_string_here"       // Replace with a valid multibase-encoded string for testing
// 	expectedEncoding := multibase.Base58BTC       // Replace with the expected multibase encoding
// 	expectedBytes := []byte("decoded_bytes_here") // Replace with the expected decoded bytes
// 	actualEncoding, actualBytes, err := internal.Decode(encoded)
// 	if err != nil {
// 		t.Fatalf("Decode failed: %v", err)
// 	}
// 	if actualEncoding != expectedEncoding || string(actualBytes) != string(expectedBytes) {
// 		t.Errorf("Expected %v %v, got %v %v", expectedEncoding, expectedBytes, actualEncoding, actualBytes)
// 	}
// }

func TestIsAlnum(t *testing.T) {
	if !internal.IsAlnum("hello123") {
		t.Errorf("IsAlnum failed for alphanumeric string")
	}
	if internal.IsAlnum("hello-123") {
		t.Errorf("IsAlnum succeeded for non-alphanumeric string")
	}
}

func TestIsValidMultibase(t *testing.T) {
	validMultibase := "k51qzi5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd"
	if !internal.IsValidMultibase(validMultibase) {
		t.Errorf("IsValidMultibase failed for valid multibase string")
	}

	invalidMultibase := "i5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd"
	if internal.IsValidMultibase(invalidMultibase) {
		t.Errorf("IsValidMultibase succeeded for invalid multibase string")
	}
}

func TestIsValidNanoID(t *testing.T) {
	if !internal.IsValidNanoID("ABC123_-") {
		t.Errorf("IsValidNanoID failed for valid NanoID string")
	}
	if internal.IsValidNanoID("ABC 123") {
		t.Errorf("IsValidNanoID succeeded for invalid NanoID string")
	}
}

func TestCreateTimeFromIsoString(t *testing.T) {
	validTime := "2022-01-01T01:01:01Z"
	expectedTime, _ := time.Parse(time.RFC3339, validTime)
	actualTime, err := internal.CreateTimeFromIsoString(validTime)
	if err != nil {
		t.Fatalf("CreateTimeFromIsoString failed: %v", err)
	}
	if !actualTime.Equal(expectedTime) {
		t.Errorf("Expected %v, got %v", expectedTime, actualTime)
	}

	invalidTime := "not_a_time"
	_, err = internal.CreateTimeFromIsoString(invalidTime)
	if err == nil {
		t.Errorf("CreateTimeFromIsoString succeeded for invalid time string")
	}
}
