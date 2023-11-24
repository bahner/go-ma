package msg

import (
	"errors"
	"fmt"
	"time"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/msg/mime"
	semver "github.com/blang/semver/v4"
)

// Validate checks if a Message instance is valid
func (m *Message) Validate() error {

	var err error

	if m == nil {
		return errors.New("nil Message provided")
	}

	if m.MimeType != mime.MESSAGE_MIME_TYPE {
		return errors.New("invalid Message type")
	}

	// Check that message body is not empty
	if m.Body == nil {
		return errors.New("body must be non-empty")
	}

	// Verify ID
	err = m.VerifyID()
	if err != nil {
		return err
	}

	// Verify actors
	err = m.VerifyActors()
	if err != nil {
		return err
	}

	// Message version check
	err = m.VerifyTimestamps()
	if err != nil {
		return err
	}

	// Message version check
	err = m.VerifyMessageVersion()
	if err != nil {
		return err
	}

	return nil
}

// Compare messageVersion.  Return nil if ok else an error
func (m *Message) VerifyMessageVersion() error {

	messageSemver, err := m.SemVersion()
	if err != nil {
		return err
	}

	supportedSemver, err := semver.Make(ma.VERSION)
	if err != nil {
		return fmt.Errorf("error parsing version constant: %s", err)
	}

	// Compare versions
	// If they are the same, we are good.
	if messageSemver.Equals(supportedSemver) {
		return nil
	}

	// If they are not the same, we need to check if the message version is greater or less than the supported verion.
	// For know this is just for informational purposes.
	if messageSemver.GT(supportedSemver) {
		return fmt.Errorf("Message version %s is greater than supported version %s", messageSemver, supportedSemver)
	}

	if messageSemver.LT(supportedSemver) {
		return fmt.Errorf("Message version %s is less than supported version %s", messageSemver, supportedSemver)
	}

	return fmt.Errorf("Message version %s is not supported", messageSemver)
}

func (m *Message) VerifyTimestamps() error {
	// Time-based checks
	now := time.Now()

	created_time, err := m.CreatedTime()
	if err != nil {
		return fmt.Errorf("invalid CreatedTime: %w", err)
	}

	if created_time.After(now) {
		return errors.New("CreatedTime must be in the past")
	}

	expires_time, err := m.ExpiresTime()
	if err != nil {
		return fmt.Errorf("invalid CreatedTime: %w", err)
	}

	if expires_time.Before(created_time) {
		return errors.New("ExpiresTime must be after CreatedTime")
	}

	return nil
}

func (m *Message) VerifyActors() error {

	var err error

	_, err = did.NewFromDID(m.From)
	if err != nil {
		return err
	}
	_, err = did.NewFromDID(m.To)
	if err != nil {
		return err
	}

	// Must they?
	if m.From == m.To {
		return errors.New("actors From and To must be different")
	}

	return nil

}

// Check that ID is valid
func (m *Message) VerifyID() error {
	if m.ID == "" {
		return errors.New("ID must be non-empty")
	}

	if !internal.IsValidNanoID(m.ID) {
		return errors.New("ID must be a valid NanoID")
	}

	return nil
}

func (m *Message) IsValid() bool {
	return m.Validate() == nil
}
