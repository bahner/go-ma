package internal

import (
	"testing"
)

// func TestEncode(t *testing.T) {
// 	input := []byte("hello")
// 	expectedEncoding := multibase.Base58BTC // Replace with the expected multibase encoding for your MESSAGE_ma.MULTIBASE_ENCODING
// 	encoded_result, _ := multibase.Encode(multibase.Encoding(expectedEncoding), input)
// 	expected_encoding, expected_result, _ := multibase.Encode(expected)
// 	actual, err := Encode(input)
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
// 	actualEncoding, actualBytes, err := Decode(encoded)
// 	if err != nil {
// 		t.Fatalf("Decode failed: %v", err)
// 	}
// 	if actualEncoding != expectedEncoding || string(actualBytes) != string(expectedBytes) {
// 		t.Errorf("Expected %v %v, got %v %v", expectedEncoding, expectedBytes, actualEncoding, actualBytes)
// 	}
// }

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
