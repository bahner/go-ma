package msg

import (
	"fmt"
	"strings"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/utils"
	semver "github.com/blang/semver/v4"
)

// check if a Message headers are valid
func (h *Headers) validate() error {

	var err error

	if h == nil {
		return ErrNilMessage
	}

	// Verify ID
	err = verifyID(h.Id)
	if err != nil {
		return err
	}

	err = verifyType(h.Type)
	if err != nil {
		return err
	}

	// Message version check. Check the type first
	err = verifyMessageVersion(h.Type)
	if err != nil {
		return err
	}

	err = verifyContentType(h.ContentType)
	if err != nil {
		return err
	}

	// Verify actors
	err = h.verifyActors()
	if err != nil {
		return err
	}

	return nil
}

// Compare messageVersion.  Return nil if ok else an error
// Takes the type string as input
func verifyMessageVersion(t string) error {

	// Split the string on "/"
	parts := strings.Split(t, "/")

	// Make a semver version from the last element
	messageSemver, err := semver.Make(parts[len(parts)-1])
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

	_, err = did.NewFromString(h.From)
	if err != nil {
		return err
	}

	if h.ContentType == BROADCAST {
		if h.To != "" {
			return ErrBroadcastHasRecipient
		}
	} else {
		_, err = did.NewFromString(h.To)
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
func verifyID(id string) error {
	if id == "" {
		return ErrEmptyID
	}

	if !utils.ValidateNanoID(id) {
		return ErrInvalidID
	}

	return nil
}

func verifyType(t string) error {

	if t == BROADCAST ||
		t == MESSAGE {
		return nil
	}

	return ErrInvalidMessageType

}

// We don't want to parse this. That's up to the receiver
// But we do want to check that it is there.
func verifyContentType(ct string) error {

	if ct == "" {
		return ErrMissingContentType
	}

	return nil
}

// We don't want to parse this. That's up to the receiver
// But we do want to check that it is there.
func verifyContent(c []byte) error {

	if c == nil {
		return ErrMissingContent
	}

	return nil
}
