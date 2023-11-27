package ma

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

)
