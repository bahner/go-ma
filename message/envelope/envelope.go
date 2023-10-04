package envelope

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/internal"
)

const MIMEType = "application/x-ma-envelope"

// Bask the encrypted message and the encrypted symmetric key in a JSON envelope.
type Envelope struct {
	MIMEType     string `json:"mime_type"`
	EphemeralKey string `json:"ephemeral_key"`
	EncryptedMsg string `json:"encrypted_msg"`
}

// Use a pointer here, this might be arbitrarily big.
func New(encodedCipherText string, encodedEphemeralKey string) (*Envelope, error) {
	return &Envelope{
		MIMEType:     MIMEType,
		EphemeralKey: encodedEphemeralKey,
		EncryptedMsg: encodedCipherText,
	}, nil
}

func (e *Envelope) MarshalToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func UnmarshalFromJSON(data []byte) (*Envelope, error) {

	e := &Envelope{}

	err := json.Unmarshal(data, e)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("envelope: error unmarshalling envelope: %s\n", err))
	}

	return e, nil
}

func (e *Envelope) String() string {
	data, err := e.MarshalToJSON()
	if err != nil {
		return ""
	}
	return string(data)
}

func (e *Envelope) GetEncryptedKey() string {
	return e.EphemeralKey
}

func (e *Envelope) GetEncryptedMsg() string {
	return e.EncryptedMsg
}

func (e *Envelope) GetMIMEType() string {
	return e.MIMEType
}
