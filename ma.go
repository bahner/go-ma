package ma

import (
	"time"

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
	HASH_ALGORITHM_MULTICODEC_STRING = "blake3"
	BLAKE3_LABEL                     = "ma"
	BLAKE3_SUM_SIZE                  = 32 // 256 bits

	// Keytype labels
	VERIFICATION_METHOD_KEY_TYPE       = "MultiKey"
	KEY_AGREEMENT_MULTICODEC_STRING    = "x25519-pub"
	ASSERTION_METHOD_MULTICODEC_STRING = "ed25519-pub"

	// Timing
	// Let the keys last a month by default. That gives ample time to rotate keys.
	VERIFICATION_METHOD_DEFAULT_TTL = time.Duration(30) * time.Hour * 24
)
