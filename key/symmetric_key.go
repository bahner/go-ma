package key

import (
	"github.com/bahner/go-ma"
	"lukechampine.com/blake3"
)

// Generate a symmetric key from a shared secret.
func GenerateSymmetricKey(shared []byte, size int) []byte {

	// Add a label, so we have out own namespace
	label := []byte(ma.BLAKE3_LABEL)
	hasher := blake3.New(size, nil)
	hasher.Write(label)
	hasher.Write(shared)
	return hasher.Sum(nil)[:size]
}
