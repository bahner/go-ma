package msg_test

import (
	"bytes"
	"testing"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/msg"
	cbor "github.com/fxamacker/cbor/v2"
)

var input_message = msg.ValidExampleMessage()

func TestMarshalToCBOR(t *testing.T) {

	expected, _ := cbor.Marshal(input_message)
	actual, err := input_message.MarshalToCBOR()
	if err != nil {
		t.Fatalf("MarshalToJSON failed: %v", err)
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("Expected %s, got %s", string(expected), string(actual))
	}
}

func TestPack(t *testing.T) {

	jsonData, _ := input_message.MarshalToCBOR()
	expected, _ := internal.MultibaseEncode(jsonData) // Assuming `Encode` works correctly
	actual, err := input_message.Pack()
	if err != nil {
		t.Fatalf("Pack failed: %v", err)
	}

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestUnpack(t *testing.T) {
	packed := msg.Message_test_packed_valid_message
	expected := msg.ValidExampleMessage()

	actual, err := msg.Unpack(packed)
	if err != nil {
		t.Fatalf("Unpack failed: %v", err)
	}

	expectedJSON, _ := cbor.Marshal(expected)
	actualJSON, _ := cbor.Marshal(actual)

	if !bytes.Equal(expectedJSON, actualJSON) {
		t.Errorf("Expected %s, got %s", string(expectedJSON), string(actualJSON))
	}
}
