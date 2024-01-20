package body

import (
	"fmt"

	cbor "github.com/fxamacker/cbor/v2"
	"lukechampine.com/blake3"
)

// This struct is a convenience wrapper for the actual content,
// in order to easily generate the headers. This is just a golangspecific
// implementation to facilitate the actual message generation.

type Body struct {
	// The content type of the message
	ContentType string `cbor:"contentType" json:"contentType"`
	// The content itself
	Content []byte `cbor:"content" json:"content"`
}

func New(content []byte, contentType string) (*Body, error) {
	return &Body{
		ContentType: contentType,
		Content:     content,
	}, nil
}

// Returns the content as a byte array
func (b *Body) Bytes() []byte {
	return b.Content
}

// Returns the Blake3 256 bit hash of the content as a hexadecimal string
// This makes is easier to compare the hashes of two messages.
// Returns empty string if an error occurs
func (b *Body) Hash() string {

	body, err := b.MarshalToCBOR()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", blake3.Sum256(body))
}

func (b *Body) MarshalToCBOR() ([]byte, error) {
	return cbor.Marshal(b)
}

func (b *Body) Size() int {
	return len(b.Content)
}

// Returns the contents as a string
func (b *Body) String() string {
	return string(b.Content)
}
