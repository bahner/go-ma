package envelope

import "github.com/bahner/go-ma/message"

type EnvelopeInterface interface {
	GetEncryptedKey() string
	GetEncryptedMsg() string
	GetMimeType() string
	MarshalToJSON() ([]byte, error)
	String() string
	Open() (*message.Message, error)
}
