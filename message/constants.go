package message

import (
	"time"

	"github.com/bahner/go-ma"
)

const (

	// Messages which are older than a day should be ignored
	MESSAGE_TTL = time.Hour * 24

	// How we identify the messages we support
	MESSAGE_ENCRYPTION_LABEL = ma.MESSAGE_MIME_TYPE
)
