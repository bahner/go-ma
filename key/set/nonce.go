package set

import (
	"github.com/bahner/go-ma/did"
	"lukechampine.com/blake3"
)

// This is the recommended nonceSize for AES GCM
const NONCE_SIZE = 12

// Creates a deterministic Nonce for the Keyset encryption by
// hashing the DID and returning the first 12 bytes.
func Nonce(did did.DID) []byte {

	idHash := blake3.Sum256([]byte(did.Id))
	return idHash[:NONCE_SIZE]
}
