package msg

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
)

// Bask the encrypted message and the encrypted symmetric key in a JSON envelope.
type Envelope struct {
	EphemeralKey []byte `cbor:"ephemeralKey" json:"ephemeralKey"`
	Content      []byte `cbor:"body" json:"body"`
	Headers      []byte `cbor:"headers" json:"headers"`
}

func (e *Envelope) MarshalToCBOR() ([]byte, error) {
	return cbor.Marshal(e)
}

func (e *Envelope) MarshalToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// Takes the envelope as a byte array and returns a pointer to an Envelope struct
// Basically this is what you do with a receieved message envelope, eg. in an Open() function.
func UnmarshalEnvelopeFromCBOR(data []byte) (*Envelope, error) {

	e := &Envelope{}

	err := cbor.Unmarshal(data, e)
	if err != nil {
		return nil, fmt.Errorf("envelope: error unmarshalling envelope: %s", err)
	}

	return e, nil
}

// Returns the contents as a multibase encoded string
// Returns empty string if an error occurs
func (e *Envelope) String() string {
	data, err := e.MarshalToCBOR()
	if err != nil {
		return ""
	}

	str, err := internal.MultibaseEncode(data)
	if err != nil {
		return ""
	}

	return str
}

// Returns the contents as a byte array
// Returns nil if an error occurs
func (e *Envelope) Bytes() ([]byte, error) {

	return e.MarshalToCBOR()
}

func (e *Envelope) GetContent(privkey []byte) ([]byte, error) {
	return decrypt(e.Content, e.EphemeralKey, privkey)
}
func (e *Envelope) GetHeaders(privkey []byte) ([]byte, error) {
	return decrypt(e.Headers, e.EphemeralKey, privkey)
}
