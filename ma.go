package ma

import (
	"github.com/multiformats/go-multibase"
)

const (

	// Ma version
	VERSION = "0.0.1"

	// The did prefix to use
	DID_SCHEME = "did"
	DID_METHOD = "ma"
	DID_PREFIX = DID_SCHEME + ":" + DID_METHOD + ":"

	// A MIME type for a message. Just to implement it for future proofing.
	MESSAGE_ENVELOPE_MIME_TYPE = "application/x-ma-envelope"
	MESSAGE_MIME_TYPE          = "application/x-ma-message"

	// The rendesvous string used for peer discovery in libp2p.
	RENDEZVOUS = "/ma/" + VERSION

	// Use the same multicodec everywhere.
	MULTIBASE_ENCODING = multibase.Base58BTC

	// Encryption multicodecs in the private range.
	// We should probably register these with multiformats.
	// But there's just too many open questions about
	// if, how, when where and how.

	// X25519 encryption codec
	ECDHX25519ChaCha20Poly1305BLAKE3 = 0x300100

	// X448 encryption codec
	ECDHX448ChaCha20Poly1305BLAKE3 = 0x300200
)
