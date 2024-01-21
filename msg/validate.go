package msg

import (
	"errors"
	"fmt"
	"time"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	semver "github.com/blang/semver/v4"
)

// check if a Message headers are valid
func (h *Headers) validate() error {

	var err error

	if h == nil {
		return errors.New("nil Message provided")
	}

	if h.MimeType != ma.MESSAGE_MIME_TYPE {
		return errors.New("invalid Message type")
	}

	// Check that message body headers are valid
	if h.ContentType == "" {
		return errors.New("ContentType must be non-empty")
	}

	// Verify ID
	err = h.verifyID()
	if err != nil {
		return err
	}

	// Verify actors
	err = h.verifyActors()
	if err != nil {
		return err
	}

	// Message version check
	err = h.verifyTimestamps()
	if err != nil {
		return err
	}

	// Message version check
	err = h.verifyMessageVersion()
	if err != nil {
		return err
	}

	return nil
}

// Compare messageVersion.  Return nil if ok else an error
func (h *Headers) verifyMessageVersion() error {

	messageSemver, err := h.SemVersion()
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
		return fmt.Errorf("message version %s is greater than supported version %s", messageSemver, supportedSemver)
	}

	if messageSemver.LT(supportedSemver) {
		return fmt.Errorf("message version %s is less than supported version %s", messageSemver, supportedSemver)
	}

	return fmt.Errorf("message version %s is not supported", messageSemver)
}

func (h *Headers) verifyTimestamps() error {
	// Time-based checks
	now := time.Now()

	created_time := h.CreatedTime()

	if created_time.After(now) {
		return errors.New("CreatedTime must be in the past")
	}

	expires_time := h.ExpiresTime()

	if expires_time.Before(created_time) {
		return errors.New("ExpiresTime must be after CreatedTime")
	}

	return nil
}

func (h *Headers) verifyActors() error {

	var err error

	_, err = did.New(h.From)
	if err != nil {
		return err
	}
	_, err = did.New(h.To)
	if err != nil {
		return err
	}

	// Must they?
	if h.From == h.To {
		return errors.New("actors From and To must be different")
	}

	return nil
}

// Check that ID is valid
func (h *Headers) verifyID() error {
	if h.ID == "" {
		return errors.New("ID must be non-empty")
	}

	if !internal.IsValidNanoID(h.ID) {
		return errors.New("ID must be a valid NanoID")
	}

	return nil
}
