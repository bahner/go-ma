package envelope

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
)

// Bask the encrypted message and the encrypted symmetric key in a JSON envelope.
type Envelope struct {
	MIMEType string `cbor:"mime_type"`
	Seal     string `cbor:"seal"`
	Message  string `cbor:"message"`
}

// Use a pointer here, this might be arbitrarily big.
func New(encodedCipherText string, encodedEphemeralKey string) (*Envelope, error) {
	return &Envelope{
		MIMEType: ma.ENVELOPE_MIME_TYPE,
		Seal:     encodedEphemeralKey,
		Message:  encodedCipherText,
	}, nil
}

func (e *Envelope) MarshalToCBOR() ([]byte, error) {
	return cbor.Marshal(e)
}

func (e *Envelope) MarshalToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func UnmarshalFromCBOR(data []byte) (*Envelope, error) {

	e := &Envelope{}

	err := cbor.Unmarshal(data, e)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("envelope: error unmarshalling envelope: %s\n", err))
	}

	return e, nil
}

func (e *Envelope) String() string {
	data, err := e.MarshalToCBOR()
	if err != nil {
		return ""
	}
	return string(data)
}

func (e *Envelope) GetEncryptedKey() string {
	return e.Seal
}

func (e *Envelope) GetEncryptedMsg() string {
	return e.Message
}

func (e *Envelope) GetMIMEType() string {
	return e.MIMEType
}
