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
	MESSAGE_MIME_TYPE  = "application/x-ma-message; version=0.0.1"

	// The rendezvous string used for peer discovery in libp2p.
	RENDEZVOUS = "/ma/" + VERSION

	// Use the same multicodec everywhere.
	MULTIBASE_ENCODING = multibase.Base58BTC

	// BLAKE3 label for symmetric key generation.
	HASH_ALGORITHM_MULTICODEC_STRING = "blake3"
	BLAKE3_LABEL                     = "ma"
	BLAKE3_SUM_SIZE                  = 32 // 256 bits

	// Verification Method types
	KEY_AGREEMENT_MULTICODEC_STRING    = "x25519-pub"
	VERIFICATION_KEY_MULTICODEC_STRING = "ed25519-pub"
	// We don't use these, but it's here for reference. If a
	// more modern suite is created, we should ise them.
	// These 2 requires SHA-256, but we use blake3.
	// KEY_AGREEMENT_KEY_TYPE = "X25519KeyAgreementKey2020"
	// VERIFICATION_KEY_TYPE  = "Ed25519VerificationKey2020"
	KEY_AGREEMENT_KEY_TYPE = "MultiKey"
	VERIFICATION_KEY_TYPE  = "MultiKey"
)
