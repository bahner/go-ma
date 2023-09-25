package message_test

import (
	"testing"

	"github.com/bahner/go-ma/message"
)

func TestGenMessage(t *testing.T) {
	msg := message.ValidExampleMessage()
	packed_message, err := msg.Pack()
	if packed_message == "" {
		t.Errorf("Packed message is empty")
	}

	if err != nil {
		t.Errorf("Pack failed: %v", err)
	}

	if packed_message != message.Message_test_packed_valid_message {
		t.Error("Packed message does not match expected value")
	}

}
func TestParse(t *testing.T) {
	// Use a valid packed message for testing.
	parsed, err := message.Parse(message.Message_test_packed_valid_message)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// Check if parsed matches your expectation (this assumes you have equality comparison for Message)
	expected := message.ValidExampleMessage()
	if *parsed != *expected {
		t.Errorf("Expected parsed to be %v, got %v", expected, parsed)
	}
}

func TestIsValid(t *testing.T) {
	msg := message.ValidExampleMessage()

	if err := msg.IsValid(); err != nil {
		t.Errorf("IsValid failed: %v", err)
	}

	// You can add more test cases for invalid Messages.
}

func TestVerifyMessageVersion(t *testing.T) {
	msg := message.ValidExampleMessage()

	if err := msg.VerifyMessageVersion(); err != nil {
		t.Errorf("VerifyMessageVersion failed: %v", err)
	}
}

func TestVerifyTimestamps(t *testing.T) {
	msg := message.ValidExampleMessage()

	if err := msg.VerifyTimestamps(); err != nil {
		t.Errorf("VerifyTimestamps failed: %v", err)
	}
}
