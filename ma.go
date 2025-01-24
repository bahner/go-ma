package ma

const (

	// Just our name
	NAME        = "ma"
	PROPER_NAME = "間"

	// Ma version
	VERSION = "0.0.1"

	// The rendezvous string used for peer discovery in libp2p.
	RENDEZVOUS = "/" + NAME + "/" + VERSION

	// Broadcast topic for global messages.
	BROADCAST_TOPIC = "/" + NAME + "/broadcast/" + VERSION

	// BLAKE3 label for symmetric key generation.
	BLAKE3_CONTENT_LABEL = RENDEZVOUS
	BLAKE3_HEADERS_LABEL = NAME
	BLAKE3_SUM_SIZE      = 32 // 256 bits

        IDENTITY_NONCE_SIZE = 12

)
