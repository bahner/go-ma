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

	// BLAKE3 label for symmetric key generation.
	HASH_ALGORITHM_MULTICODEC_STRING = "blake3"
	BLAKE3_LABEL                     = NAME
	BLAKE3_SUM_SIZE                  = 32 // 256 bits

	// Message constants
	MESSAGE_TYPE           = "/ma/message/" + VERSION
	ENVELOPE_MESSAGE_TYPE  = "/ma/message/envelope/" + VERSION
	BROADCAST_MESSAGE_TYPE = "/ma/message/broadcast/" + VERSION

	BROADCAST_TOPIC = BROADCAST_MESSAGE_TYPE

	MESSAGE_DEFAULT_CONTENT_TYPE = "text/plain"

	// API
	DEFAULT_IPFS_API_MULTIADDR = "/ip4/127.0.0.1/tcp/45005" // Default to Brave browser, Kubo is /ip4/127.0.0.1/tcp/5001
	ENV_PREFIX                 = "GO_MA"
)
