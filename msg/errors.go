package msg

import (
	"errors"
	"fmt"

	"github.com/bahner/go-ma"
)

var (
	ErrNilMessage                = errors.New("nil Message provided")
	ErrMessageInvalidType        = errors.New("invalid Message type")
	ErrMessageMissingContentType = errors.New("empty ContentType")
	ErrVersionInvalid            = fmt.Errorf("version not %s", ma.VERSION)
	ErrVersionTooHigh            = fmt.Errorf("version is higher %s", ma.VERSION)
	ErrVersionTooLow             = fmt.Errorf("version is less than %s", ma.VERSION)
	ErrSameActor                 = errors.New("header From and To be different")
	ErrEmptyID                   = errors.New("id must be non-empty")
	ErrInvalidID                 = errors.New("id must be a valid NanoID")
	ErrInvalidDID                = errors.New("must be a valid DID")
	ErrMissingFrom               = errors.New("mmissing From sender")
	ErrMissingTo                 = errors.New("mmissing To reciient")
	ErrMissinSignature           = errors.New("mmissing signature")
	ErrFetchDoc                  = errors.New("failed to fetch entity document")
)
