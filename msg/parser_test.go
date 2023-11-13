package msg_test

import (
	"testing"

	"github.com/bahner/go-ma/msg"
)

func TestGenMessage(t *testing.T) {
	m := msg.ValidExampleMessage()
	packed_message, err := m.Pack()
	if packed_message == "" {
		t.Errorf("Packed message is empty")
	}

	if err != nil {
		t.Errorf("Pack failed: %v", err)
	}

	if packed_message != msg.Message_test_packed_valid_message {
		t.Error("Packed message does not match expected value")
	}

}
func TestParse(t *testing.T) {
	// Use a valid packed message for testing.
	parsed, err := msg.Parse(msg.Message_test_packed_valid_message)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	// Check if parsed matches your expectation (this assumes you have equality comparison for Message)
	expected := msg.ValidExampleMessage()
	if *parsed != *expected {
		t.Errorf("Expected parsed to be %v, got %v", expected, parsed)
	}
}

func TestIsValid(t *testing.T) {
	msg := msg.ValidExampleMessage()

	if err := msg.IsValid(); err != nil {
		t.Errorf("IsValid failed: %v", err)
	}

	// You can add more test cases for invalid Messages.
}

func TestVerifyMessageVersion(t *testing.T) {
	msg := msg.ValidExampleMessage()

	if err := msg.VerifyMessageVersion(); err != nil {
		t.Errorf("VerifyMessageVersion failed: %v", err)
	}
}

func TestVerifyTimestamps(t *testing.T) {
	msg := msg.ValidExampleMessage()

	if err := msg.VerifyTimestamps(); err != nil {
		t.Errorf("VerifyTimestamps failed: %v", err)
	}
}
