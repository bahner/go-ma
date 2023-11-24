package envelope

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
)

// Bask the encrypted message and the encrypted symmetric key in a JSON envelope.
type Envelope struct {
	EphemeralKey []byte `cbor:"ephemeralKey" json:"ephemeralKey"`
	Message      []byte `cbor:"message" json:"message"`
}

// Use a pointer here, this might be arbitrarily big.
func New(encodedCipherText []byte, encodedEphemeralKey []byte) (*Envelope, error) {
	return &Envelope{
		EphemeralKey: encodedEphemeralKey,
		Message:      encodedCipherText,
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

func (e *Envelope) String() string {
	data, err := e.MarshalToCBOR()
	if err != nil {
		return ""
	}
	return string(data)
}

func (e *Envelope) Bytes() ([]byte, error) {
	return e.MarshalToCBOR()
}

func (e *Envelope) GetEncryptedMsg() []byte {
	return e.Message
}

func (e *Envelope) GetEphemeralKey() []byte {
	return e.EphemeralKey
}

func (e *Envelope) MultibaseEncode() (string, error) {
	data, err := e.MarshalToCBOR()
	if err != nil {
		return "", fmt.Errorf("envelope: error marshalling envelope: %s", err)
	}

	return internal.MultibaseEncode(data)
}

func MultibaseDecode(data string) (*Envelope, error) {
	decodedData, err := internal.MultibaseDecode(data)
	if err != nil {
		return nil, fmt.Errorf("envelope: error decoding envelope: %s", err)
	}

	e, err := UnmarshalFromCBOR(decodedData)
	if err != nil {
		return nil, fmt.Errorf("envelope: error unmarshalling envelope: %s", err)
	}

	return e, nil
}
