package ma

import (
	"github.com/multiformats/go-multibase"
)

const (
	// A MIME type for a message. Just to implement it for future proofing.
	MESSAGE_MIME_TYPE = "application/x-ma-message"

	// The rendesvous string used for peer discovery in libp2p.
	// So it has nothing to do with MIME, but it's make parsing
	// the string easier, if people want to be clever about it.
	MESSAGE_RENDEZVOUS = MESSAGE_MIME_TYPE + "; rendezvous="

	// Where as people can be clever about the rendezvous string,
	// We should just use the same everyone. But maybe we need
	// split them up later. There is just too many open questions.
	RENDEZVOUS = MESSAGE_RENDEZVOUS + "SPACE"

	// Use the same multicodec everywhere.
	MULTIBASE_ENCODING = multibase.Base58BTC

	// Encryption multicodecs in the private range.
	// We should probably register these with multiformats.
	// But there's just too many open questions about
	// if, how, when where and how.
	// Multicodecs are 64 bit unsigned integers.
	ECDHX25519ChaCha20Poly1305BLAKE3256bit = 0x300001
	ECDHX25519ChaCha20Poly1305BLAKE3512bit = 0x300002
	ECDHX448ChaCha20Poly1305BLAKE3256bit   = 0x300010
	ECDHX448ChaCha20Poly1305BLAKE3512bit   = 0x300011
)
