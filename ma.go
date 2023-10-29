package ma

import (
	"github.com/multiformats/go-multibase"
)

const (

	// Just our name
	NAME = "ma"

	// Ma version
	VERSION = "0.0.1"

	// The did prefix to use
	DID_SCHEME = "did"
	DID_METHOD = "ma"
	DID_PREFIX = DID_SCHEME + ":" + DID_METHOD + ":"

	// A MIME type for a message. Just to implement it for future proofing.
	ENVELOPE_MIME_TYPE = "application/x-ma-envelope"
	MESSAGE_MIME_TYPE  = "application/x-ma-message"

	// The rendezvous string used for peer discovery in libp2p.
	RENDEZVOUS = "/ma/" + VERSION

	// Use the same multicodec everywhere.
	MULTIBASE_ENCODING = multibase.Base58BTC

	// BLAKE3 label for symmetric key generation.
	BLAKE3_LABEL    = "ma"
	BLAKE3_SUM_SIZE = 32 // 256 bits
)
