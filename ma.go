package ma

import (
	"github.com/multiformats/go-multibase"
)

const (

	// Ma version
	VERSION = "0.0.1"

	// A MIME type for a message. Just to implement it for future proofing.
	MESSAGE_MIME_TYPE = "application/x-ma-message; version=" + VERSION

	// The rendesvous string used for peer discovery in libp2p.
	MESSAGE_RENDEZVOUS = "/ma/" + VERSION

	// Use the same multicodec everywhere.
	MULTIBASE_ENCODING = multibase.Base58BTC

	// Encryption multicodecs in the private range.
	// We should probably register these with multiformats.
	// But there's just too many open questions about
	// if, how, when where and how.

	// X25519 encryption codec
	ECDHX25519ChaCha20Poly1305BLAKE3 = 0x300010

	// X448 encryption codec
	ECDHX448ChaCha20Poly1305BLAKE3 = 0x300020
)
