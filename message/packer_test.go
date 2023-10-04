package message_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/message"
)

var input_message = message.ValidExampleMessage()

func TestMarshalToJSON(t *testing.T) {

	expected, _ := json.Marshal(input_message)
	actual, err := input_message.MarshalToJSON()
	if err != nil {
		t.Fatalf("MarshalToJSON failed: %v", err)
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("Expected %s, got %s", string(expected), string(actual))
	}
}

func TestPack(t *testing.T) {

	jsonData, _ := input_message.MarshalToJSON()
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
	packed := message.Message_test_packed_valid_message
	expected := message.ValidExampleMessage()

	actual, err := message.Unpack(packed)
	if err != nil {
		t.Fatalf("Unpack failed: %v", err)
	}

	expectedJSON, _ := json.Marshal(expected)
	actualJSON, _ := json.Marshal(actual)

	if !bytes.Equal(expectedJSON, actualJSON) {
		t.Errorf("Expected %s, got %s", string(expectedJSON), string(actualJSON))
	}
}
