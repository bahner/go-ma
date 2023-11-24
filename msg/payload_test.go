package msg_test

import (
	"bytes"
	"testing"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/msg"
	"github.com/bahner/go-ma/msg/mime"
	cbor "github.com/fxamacker/cbor/v2"
)

// Helper function to create a valid Message instance for testing.
func validMessageWithSignature() *msg.Message {
	return &msg.Message{
		ID:        "validNanoID",
		MimeType:  mime.MESSAGE_MIME_TYPE,
		From:      "did:ma:from",
		To:        "did:ma:to",
		Created:   1698684192,
		Expires:   1698687792,
		Body:      []byte("Hello"),
		Version:   ma.VERSION,
		Signature: "signature",
	}
}

func TestPayload(t *testing.T) {
	m := validMessageWithSignature()

	payload, err := m.Unsigned()
	if err != nil {
		t.Fatalf("Payload failed: %v", err)
	}

	if payload.Signature != "" {
		t.Errorf("Expected empty signature in payload, got %s", payload.Signature)
	}
}

func TestMarshalPayloadToCBOR(t *testing.T) {
	m := validMessageWithSignature()

	jsonData, err := m.MarshalUnsignedToCBOR()
	if err != nil {
		t.Fatalf("MarshalPayloadToJSON failed: %v", err)
	}

	payload, _ := m.Unsigned()
	expected, _ := cbor.Marshal(payload)

	if !bytes.Equal(expected, jsonData) {
		t.Errorf("Expected %s, got %s", string(expected), string(jsonData))
	}
}
