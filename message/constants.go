package message

import (
	"time"

	"github.com/multiformats/go-multibase"
)

const (
	// What we are working on. The name of the message type
	// in MIME format to give parsers a hint.
	MESSAGE_MIME_TYPE = "application/x-ma-message"

	// MESSAGE_VERSION of the message schema
	MESSAGE_VERSION = "0.0.1"

	// Messages which are older than a day should be ignored
	MESSAGE_TTL = time.Hour * 24

	// How we identify the messages we support
	MESSAGE_ENCRYPTION_LABEL = MESSAGE_MIME_TYPE

	// MESSAGE_ENCODER_ENCODING is the encoding used for all encoding in messages.
	MESSAGE_ENCODER_ENCODING = multibase.Base58BTC
)
