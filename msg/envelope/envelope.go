package envelope

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
)

// Bask the encrypted message and the encrypted symmetric key in a JSON envelope.
type Envelope struct {
	Seal    string `cbor:"seal" json:"seal"`
	Message string `cbor:"message" json:"message"`
}

// Use a pointer here, this might be arbitrarily big.
func New(encodedCipherText string, encodedEphemeralKey string) (*Envelope, error) {
	return &Envelope{
		Seal:    encodedEphemeralKey,
		Message: encodedCipherText,
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
		return nil, fmt.Errorf("envelope: error unmarshalling envelope: %s", err)
	}

	return e, nil
}

func (e *Envelope) Pack() (string, error) {
	marshalled, err := e.MarshalToCBOR()
	if err != nil {
		return "", fmt.Errorf("envelope: failed to marshal envelope: %w", err)
	}

	packed, err := internal.MultibaseEncode(marshalled)
	if err != nil {
		return "", fmt.Errorf("envelope: failed to multibase encode envelope: %w", err)
	}

	return packed, nil

}

func Unpack(packed string) (*Envelope, error) {

	data, err := internal.MultibaseDecode(packed)
	if err != nil {
		return nil, fmt.Errorf("envelope: failed to multibase decode envelope: %w", err)
	}

	e, err := UnmarshalFromCBOR(data)
	if err != nil {
		return nil, fmt.Errorf("envelope: failed to unmarshal envelope: %w", err)
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
