package msg

import (
	"time"
)

const (

	// A MIME type for a message. Just to implement it for future proofing.
	MIME_TYPE = "application/x-ma-message; version=0.0.1"

	// Messages which are older than a day should be ignored
	MESSAGE_TTL = time.Hour * 24

	// How we identify the messages we support
	MESSAGE_ENCRYPTION_LABEL = MIME_TYPE
)
