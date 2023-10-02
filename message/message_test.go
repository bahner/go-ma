package message_test

import (
	"testing"
	"time"

	"github.com/Masterminds/semver"
	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/message"

	"reflect"
)

func TestNewMessage(t *testing.T) {

	msg := message.ValidExampleMessage()
	from := msg.From
	to := msg.To
	body := msg.Body
	mime := message.MimeType("text")

	msg, err := message.New(from, to, body, mime)
	if err != nil {
		t.Fatalf("Failed to create new message: %v", err)
	}

	if msg.From != from {
		t.Errorf("Expected From to be %s, got %s", from, msg.From)
	}

	if msg.To != to {
		t.Errorf("Expected To to be %s, got %s", to, msg.To)
	}

	if msg.Body != body {
		t.Errorf("Expected Body to be %s, got %s", body, msg.Body)
	}

	// Add more checks for ID, CreatedTime, ExpiresTime, etc.
}

func TestMessageSender(t *testing.T) {
	validMsg := message.ValidExampleMessage()
	from := validMsg.From
	msg := &message.Message{From: from}

	sender, err := msg.Sender()
	if err != nil {
		t.Fatalf("Failed to get sender: %v", err)
	}

	// Assuming did.Parse works correctly, replace with your actual check
	expectedSender, _ := did.Parse(from)

	if !reflect.DeepEqual(sender, expectedSender) {
		t.Errorf("Expected sender to be %v, got %v", expectedSender, sender)
	}
}

func TestMessageRecipient(t *testing.T) {
	validMsg := message.ValidExampleMessage()
	to := validMsg.To
	msg := &message.Message{To: to}

	recipient, err := msg.Recipient()
	if err != nil {
		t.Fatalf("Failed to get recipient: %v", err)
	}

	// Assuming did.Parse works correctly, replace with your actual check
	expectedRecipient, _ := did.Parse(to)

	if !reflect.DeepEqual(recipient, expectedRecipient) {
		t.Errorf("Expected sender to be %v, got %v", expectedRecipient, recipient)
	}
}

func TestMessageCreated(t *testing.T) {
	createdTime := time.Now().UTC().Format(time.RFC3339)
	msg := &message.Message{CreatedTime: createdTime}

	created, err := msg.Created()
	if err != nil {
		t.Fatalf("Failed to get created time: %v", err)
	}

	// Assuming createTimeFromIsoString works correctly, replace with your actual check
	expectedTime, _ := time.Parse(time.RFC3339, createdTime)
	if !created.Equal(expectedTime) {
		t.Errorf("Expected created time to be %v, got %v", expectedTime, created)
	}
}

func TestMessageExpires(t *testing.T) {
	expiresTime := time.Now().UTC().Add(time.Hour).Format(time.RFC3339)
	msg := &message.Message{ExpiresTime: expiresTime}

	expires, err := msg.Expires()
	if err != nil {
		t.Fatalf("Failed to get expires time: %v", err)
	}

	// Assuming createTimeFromIsoString works correctly, replace with your actual check
	expectedTime, _ := time.Parse(time.RFC3339, expiresTime)
	if !expires.Equal(expectedTime) {
		t.Errorf("Expected expires time to be %v, got %v", expectedTime, expires)
	}
}

func TestMessageSemVersion(t *testing.T) {
	msg := &message.Message{Version: ma.VERSION}

	expectedVersion, _ := semver.NewVersion(ma.VERSION)
	parsedVersion, err := msg.SemVersion()
	if err != nil {
		t.Fatalf("Failed to parse semver: %v", err)
	}
	if !expectedVersion.Equal(parsedVersion) {
		t.Errorf("Expected %s, got %s", expectedVersion, parsedVersion)
	}
}
