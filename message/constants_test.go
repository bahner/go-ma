package message_test

import (
	"testing"
	"time"

	"github.com/bahner/go-ma/message"
)

func TestConstants(t *testing.T) {
	if message.MESSAGE_VERSION != "0.0.1" {
		t.Errorf("Expected version to be '0.0.1', got %s", message.MESSAGE_VERSION)
	}

	if message.MESSAGE_MIME_TYPE != "application/x-ma-message" {
		t.Errorf("Expected type to be 'ma/message', got %s", message.MESSAGE_MIME_TYPE)
	}

	if message.MESSAGE_TTL != time.Hour*24 {
		t.Errorf("Expected TTL to be 24 hours, got %s", message.MESSAGE_TTL)
	}

}
