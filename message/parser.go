package message

import (
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/semver"
	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
)

// Parses a received message in form of a multibase encoded JSON string
func Parse(data string) (*Message, error) {

	p, err := Unpack(data)
	if err != nil {
		return nil, err
	}

	err = p.IsValid()
	if err != nil {
		return nil, err
	}

	return p, nil
}

// IsValid checks if a Message instance is valid
func (m *Message) IsValid() error {

	var err error

	if m == nil {
		return errors.New("nil Message provided")
	}

	if m.MimeType != ma.MESSAGE_MIME_TYPE {
		return errors.New("invalid Message type")
	}

	// Check that message body is not empty
	if m.Body == "" {
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

	supportedSemver, err := semver.NewVersion(ma.VERSION)
	if err != nil {
		return fmt.Errorf("error parsing version constant: %s", err)
	}

	// Compare versions
	// If they are the same, we are good.
	if messageSemver.Equal(supportedSemver) {
		return nil
	}

	// If they are not the same, we need to check if the message version is greater or less than the supported verion.
	// For know this is just for informational purposes.
	if messageSemver.GreaterThan(supportedSemver) {
		return fmt.Errorf("Message version %s is greater than supported version %s", messageSemver, supportedSemver)
	}

	if messageSemver.LessThan(supportedSemver) {
		return fmt.Errorf("Message version %s is less than supported version %s", messageSemver, supportedSemver)
	}

	return fmt.Errorf("Message version %s is not supported", messageSemver)
}

func (m *Message) VerifyTimestamps() error {
	// Time-based checks
	now := time.Now()

	created_time, err := m.Created()
	if err != nil {
		return fmt.Errorf("invalid CreatedTime: %v", err)
	}
	if created_time.After(now) {
		return errors.New("CreatedTime must be in the past")
	}

	expires_time, err := m.Created()
	if err != nil {
		return fmt.Errorf("invalid CreatedTime: %v", err)
	}

	if expires_time.Before(created_time) {
		return errors.New("ExpiresTime must be after CreatedTime")
	}

	return nil
}

func (m *Message) VerifyActors() error {

	var err error

	_, err = did.Parse(m.From)
	if err != nil {
		return err
	}
	_, err = did.Parse(m.To)
	if err != nil {
		return err
	}

	// // FIXME: Must they?
	// if from == to {
	// 	return errors.New("From and To must be different")
	// }

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
