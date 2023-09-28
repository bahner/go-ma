package envelope

type EnvelopeInterface interface {
	GetEncryptedKey() string
	GetEncryptedMsg() string
	GetMimeType() string
	MarshalToJSON() ([]byte, error)
	String() string
}
