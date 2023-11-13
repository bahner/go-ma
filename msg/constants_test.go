package msg_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/msg"
)

func TestConstants(t *testing.T) {
	if ma.VERSION != "0.0.1" {
		t.Errorf("Expected version to be '0.0.1', got %s", ma.VERSION)
	}

	expected_message_mime_type := fmt.Sprintf("application/x-ma-message; version=%s", ma.VERSION)
	if msg.MIME_TYPE != expected_message_mime_type {
		t.Errorf("Expected type to be 'ma/message', got %s", msg.MIME_TYPE)
	}

	if msg.MESSAGE_TTL != time.Hour*24 {
		t.Errorf("Expected TTL to be 24 hours, got %s", msg.MESSAGE_TTL)
	}

}
