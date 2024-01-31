package ma

import "time"

const (

	// Just our name
	NAME = "ma"

	// Ma version
	VERSION = "0.0.1"

	// The did prefix to use
	DID_PREFIX = "did:" + NAME + ":"

	// The rendezvous string used for peer discovery in libp2p.
	RENDEZVOUS = "/" + NAME + "/" + VERSION

	// The topic prefix used in pubsub. It is the same as the rendezvous string.
	// But we keep it separate for future flexibility.
	TOPIC_PREFIX = RENDEZVOUS

	// BLAKE3 label for symmetric key generation.
	HASH_ALGORITHM_MULTICODEC_STRING = "blake3"
	BLAKE3_LABEL                     = "ma"
	BLAKE3_SUM_SIZE                  = 32 // 256 bits

	// MIME types

	// A MIME type for a message. Just to implement it for future proofing.
	MESSAGE_MIME_TYPE  = "application/x-ma-message; version=" + VERSION
	ENVELOPE_MIME_TYPE = "application/x-ma-envelope; version=" + VERSION

	MESSAGE_DEFAULT_CONTENT_TYPE = "text/plain"
	MESSAGE_DEFAULT_TTL          = time.Hour * 24

	// API
	DEFAULT_IPFS_API_MULTIADDR = "/ip4/127.0.0.1/tcp/45005" // Default to Brave browser, Kubo is /ip4/127.0.0.1/tcp/5001
  ENV_PREFIX = "GO_MA"
)
