package msg

import (
	"fmt"

	cbor "github.com/fxamacker/cbor/v2"
)

// Bask the encrypted message and the encrypted symmetric key in a CBOR envelope.
type Envelope struct {
	EphemeralKey     []byte
	EncryptedContent []byte
	EncryptedHeaders []byte
}

func (e *Envelope) marshalToCBOR() ([]byte, error) {
	return cbor.Marshal(e)
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

func (e *Envelope) getContent(privkey []byte) ([]byte, error) {
	return decrypt(e.EncryptedContent, e.EphemeralKey, privkey)
}
func (e *Envelope) getHeaders(privkey []byte) (*Headers, error) {

	bytes, err := decrypt(e.EncryptedHeaders, e.EphemeralKey, privkey)
	if err != nil {
		return nil, fmt.Errorf("envelope: error decrypting headers: %w", err)
	}

	var hdrs *Headers = new(Headers)

	err = cbor.Unmarshal(bytes, hdrs)
	if err != nil {
		return nil, fmt.Errorf("envelope: error unmarshalling headers: %w", err)
	}

	return hdrs, nil
}
