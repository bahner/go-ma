package envelope

import (
	"lukechampine.com/blake3"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/message"
)

func Enclose(m *message.Message) (*Envelope, error) {
	return kyberEd25519Encrypt(m)
}

func generateSymmetricKey(shared []byte, size int) []byte {

	// Hash the shared secret with Blake3 in a uniform way.

	// The label is the MIME Type, just so we have our own namespace.
	label := []byte(ma.MESSAGE_MIME_TYPE)
	hasher := blake3.New(size, nil)
	hasher.Write(label)
	hasher.Write(shared)
	return hasher.Sum(nil)[:size]
}
