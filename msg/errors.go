package msg

import (
	"fmt"

	"github.com/bahner/go-ma"
)

var (
	ErrBroadcastHasRecipient = fmt.Errorf("broadcast message must not have a recipient")
	ErrBroadcastInvalidTopic = fmt.Errorf("broadcast topic must be %s", ma.BROADCAST_TOPIC)
	ErrBroadcastInvalidType  = fmt.Errorf("broadcast message must not %s", BROADCAST)
	ErrEmptyID               = fmt.Errorf("id must be non-empty")
	ErrInvalidID             = fmt.Errorf("invalid message id")
	ErrFetchDoc              = fmt.Errorf("failed to fetch entity document")
	ErrInvalidMessageType    = fmt.Errorf("invalid Message type")
	ErrInvalidSender         = fmt.Errorf("invalid sender")
	ErrInvalidRecipient      = fmt.Errorf("invalid recipient")
	ErrMissingContentType    = fmt.Errorf("empty ContentType")
	ErrMissingContent        = fmt.Errorf("empty ContentType")
	ErrMissingFrom           = fmt.Errorf("mmissing From sender")
	ErrMissinSignature       = fmt.Errorf("mmissing signature")
	ErrNilMessage            = fmt.Errorf("nil Message provided")
	ErrNilEnvelope           = fmt.Errorf("nil Envelope provided")
	ErrSameActor             = fmt.Errorf("header From and To be different")
	ErrVersionInvalid        = fmt.Errorf("version not %s", ma.VERSION)
	ErrVersionTooHigh        = fmt.Errorf("version is higher %s", ma.VERSION)
	ErrVersionTooLow         = fmt.Errorf("version is less than %s", ma.VERSION)
)
