package key

import (
	"lukechampine.com/blake3"
)

// Generate a symmetric key from a shared secret.
func GenerateSymmetricKey(shared []byte, size int, label []byte) []byte {

	hasher := blake3.New(size, nil)
	hasher.Write(label)
	hasher.Write(shared)
	return hasher.Sum(nil)[:size]
}
