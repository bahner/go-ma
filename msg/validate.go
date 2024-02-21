package msg

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	semver "github.com/blang/semver/v4"
)

// check if a Message headers are valid
func (h *Headers) validate() error {

	var err error

	if h == nil {
		return ErrNilMessage
	}

	if h.MimeType != ma.MESSAGE_MIME_TYPE && h.MimeType != ma.BROADCAST_MIME_TYPE {
		return ErrMessageInvalidType
	}

	// Check that message body headers are valid
	if h.ContentType == "" {
		return ErrMissingContentType
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
	err = h.verifyMessageVersion()
	if err != nil {
		return err
	}

	return nil
}

// Compare messageVersion.  Return nil if ok else an error
func (h *Headers) verifyMessageVersion() error {

	messageSemver, err := h.semVersion()
	if err != nil {
		return err
	}

	supportedSemver, err := semver.Make(ma.VERSION)
	if err != nil {
		return fmt.Errorf("error parsing version constant: %w", err)
	}

	// Compare versions
	// If they are the same, we are good.
	if messageSemver.Equals(supportedSemver) {
		return nil
	}

	// If they are not the same, we need to check if the message version is greater or less than the supported verion.
	// For know this is just for informational purposes.
	if messageSemver.GT(supportedSemver) {
		return fmt.Errorf("message version %s too high. %w", messageSemver, ErrVersionTooHigh)
	}

	if messageSemver.LT(supportedSemver) {
		return fmt.Errorf("message version %s too low. %w", messageSemver, ErrVersionTooLow)
	}

	return fmt.Errorf("message version %s is not supported. %w", messageSemver, ErrVersionInvalid)
}

func (h *Headers) verifyActors() error {

	var err error

	_, err = did.New(h.From)
	if err != nil {
		return err
	}

	if h.ContentType == ma.BROADCAST_MIME_TYPE {
		if h.To != "" {
			return ErrBroadcastHasRecipient
		}
	} else {
		_, err = did.New(h.To)
		if err != nil {
			return ErrInvalidRecipient
		}
	}

	// FIXME
	// Should this be an error?
	if h.From == h.To {
		return ErrSameActor
	}

	return nil
}

// Check that ID is valid
func (h *Headers) verifyID() error {
	if h.ID == "" {
		return ErrEmptyID
	}

	if !internal.IsValidNanoID(h.ID) {
		return ErrInvalidID
	}

	return nil
}
