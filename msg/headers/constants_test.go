package headers_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/msg/headers"
	"github.com/bahner/go-ma/msg/mime"
)

func TestConstants(t *testing.T) {
	if ma.VERSION != "0.0.1" {
		t.Errorf("Expected version to be '0.0.1', got %s", ma.VERSION)
	}

	expected_message_mime_type := fmt.Sprintf("application/x-ma-message; version=%s", ma.VERSION)
	if mime.MESSAGE_MIME_TYPE != expected_message_mime_type {
		t.Errorf("Expected type to be 'ma/message', got %s", mime.MESSAGE_MIME_TYPE)
	}

	if headers.MESSAGE_TTL != time.Hour*24 {
		t.Errorf("Expected TTL to be 24 hours, got %s", headers.MESSAGE_TTL)
	}

}
