package msg

import (
	"errors"
	"fmt"

	"github.com/bahner/go-ma"
)

var (
	ErrBroadcastHasRecipient = errors.New("broadcast message must not have a recipient")
	ErrBroadcastInvalidTopic = fmt.Errorf("broadcast topic must be %s", ma.BROADCAST_TOPIC)
	ErrBroadcastInvalidType  = fmt.Errorf("broadcast message must not %s", ma.BROADCAST_MESSAGE_TYPE)
	ErrEmptyID               = errors.New("id must be non-empty")
	ErrInvalidID             = errors.New("invalid message id")
	ErrFetchDoc              = errors.New("failed to fetch entity document")
	ErrMessageInvalidType    = errors.New("invalid Message type")
	ErrInvalidSender         = errors.New("invalid sender")
	ErrInvalidRecipient      = errors.New("invalid recipient")
	ErrMissingContentType    = errors.New("empty ContentType")
	ErrMissingContent        = errors.New("empty ContentType")
	ErrMissingFrom           = errors.New("mmissing From sender")
	ErrMissinSignature       = errors.New("mmissing signature")
	ErrNilMessage            = errors.New("nil Message provided")
	ErrNilEnvelope           = errors.New("nil Envelope provided")
	ErrSameActor             = errors.New("header From and To be different")
	ErrVersionInvalid        = fmt.Errorf("version not %s", ma.VERSION)
	ErrVersionTooHigh        = fmt.Errorf("version is higher %s", ma.VERSION)
	ErrVersionTooLow         = fmt.Errorf("version is less than %s", ma.VERSION)
)
