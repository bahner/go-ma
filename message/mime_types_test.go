package message_test

import (
	"testing"

	"github.com/bahner/go-ma/message"
	"github.com/stretchr/testify/assert"
)

func TestMimeTypeAliases(t *testing.T) {
	expectedAliases := []string{
		"car", "ipld_cbor", "ipld_json", "ipld", "ipns-record", "json", "message", "text",
	}
	aliases := message.MimeTypeAliases()
	assert.ElementsMatch(t, expectedAliases, aliases)
}

func TestMimeTypes(t *testing.T) {
	expectedTypes := []string{
		"application/vnd.ipld.car",
		"application/vnd.ipld.cbor",
		"application/vnd.ipld.json",
		"application/vnd.ipld.raw",
		"application/vnd.ipfs.ipns-record",
		"application/json",
		"application/x-ma-message",
		"text/plain",
	}
	types := message.MimeTypes()
	assert.ElementsMatch(t, expectedTypes, types)
}

func TestMimeTypeTuples(t *testing.T) {
	expectedTuples := [][2]string{
		{"car", "application/vnd.ipld.car"},
		{"ipld_cbor", "application/vnd.ipld.cbor"},
		{"ipld_json", "application/vnd.ipld.json"},
		{"ipld", "application/vnd.ipld.raw"},
		{"ipns-record", "application/vnd.ipfs.ipns-record"},
		{"json", "application/json"},
		{"message", "application/x-ma-message"},
		{"text", "text/plain"},
	}
	tuples := message.MimeTypeTuples()
	assert.ElementsMatch(t, expectedTuples, tuples)
}

func TestMimeType(t *testing.T) {
	tests := map[string]string{
		"car":         "application/vnd.ipld.car",
		"ipld_cbor":   "application/vnd.ipld.cbor",
		"ipld_json":   "application/vnd.ipld.json",
		"ipld":        "application/vnd.ipld.raw",
		"ipns-record": "application/vnd.ipfs.ipns-record",
		"json":        "application/json",
		"message":     "application/x-ma-message",
		"text":        "text/plain",
	}

	for alias, expectedType := range tests {
		assert.Equal(t, expectedType, message.MimeType(alias))
	}
}

func TestIsValidMimeType(t *testing.T) {
	// Valid MIME types
	validMimeTypes := []string{
		"text/plain",
		"application/json",
		"application/vnd.ipld.car",
		"text/plain; charset=utf-8",
	}
	for _, mimetype := range validMimeTypes {
		assert.True(t, message.IsValidMimeType(mimetype), "Expected MIME type to be valid: "+mimetype)
	}

	// Intentionally malformed MIME types for testing invalid cases
	invalidMimeTypes := []string{
		"text=plain",            // equals sign instead of slash
		"text/plain; charset==", // invalid parameter
	}
	for _, mimetype := range invalidMimeTypes {
		assert.False(t, message.IsValidMimeType(mimetype), "Expected MIME type to be invalid: "+mimetype)
	}
}
