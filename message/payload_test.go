package message_test

import (
	"bytes"
	"testing"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
	cbor "github.com/fxamacker/cbor/v2"
)

// Helper function to create a valid Message instance for testing.
func validMessageWithSignature() *message.Message {
	return &message.Message{
		ID:        "validNanoID",
		MimeType:  ma.MESSAGE_MIME_TYPE,
		From:      "did:ma:from",
		To:        "did:ma:to",
		Created:   1698684192,
		Expires:   1698687792,
		Body:      "Hello",
		Version:   ma.VERSION,
		Signature: "signature",
	}
}

func TestPayload(t *testing.T) {
	msg := validMessageWithSignature()

	payload, err := message.Payload(*msg)
	if err != nil {
		t.Fatalf("Payload failed: %v", err)
	}

	if payload.Signature != "" {
		t.Errorf("Expected empty signature in payload, got %s", payload.Signature)
	}
}

func TestMarshalPayloadToCBOR(t *testing.T) {
	msg := validMessageWithSignature()

	jsonData, err := msg.MarshalPayloadToCBOR()
	if err != nil {
		t.Fatalf("MarshalPayloadToJSON failed: %v", err)
	}

	payload, _ := message.Payload(*msg)
	expected, _ := cbor.Marshal(payload)

	if !bytes.Equal(expected, jsonData) {
		t.Errorf("Expected %s, got %s", string(expected), string(jsonData))
	}
}

func TestPayloadPack(t *testing.T) {
	msg := validMessageWithSignature()

	PayloadPack, err := msg.PayloadPack()
	if err != nil {
		t.Fatalf("PayloadPack failed: %v", err)
	}

	jsonData, _ := msg.MarshalPayloadToCBOR()
	expected, _ := internal.MultibaseEncode(jsonData) // Assuming `Encode` works correctly

	if PayloadPack != expected {
		t.Errorf("Expected %s, got %s", expected, PayloadPack)
	}
}
