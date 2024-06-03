package doc

import (
	cbor "github.com/fxamacker/cbor/v2"
	"lukechampine.com/blake3"
)

// Payload generates the unsigned DID,
// This is everything except the proof.
func Payload(d Document) (Document, error) {

	payloadDoc := d

	payloadDoc.Proof = Proof{}

	return payloadDoc, nil
}

// Marshals the payload to CBOR for publication
func (d *Document) MarshalPayloadToCBOR() ([]byte, error) {
	p, err := Payload(*d)
	if err != nil {
		return nil, err
	}

	return cbor.Marshal(p)
}

func (d *Document) PayloadHash() ([]byte, error) {
	p, err := d.MarshalPayloadToCBOR()
	if err != nil {
		return nil, ErrPayloadMarshal
	}

	// Hash the payload
	hashed := blake3.Sum256(p)

	return hashed[:], nil
}
