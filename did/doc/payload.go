package doc

import (
	"github.com/bahner/go-ma/utils"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/multiformats/go-multicodec"
	"lukechampine.com/blake3"
)

// Payload generates the unsigned DID,
// This is everything non-metadata in the DID document.
// We don't use a pointer here, so that we don't have to reiterate the
// struct in the function. We just need to change the signature.
func Payload(d Document) (Document, error) {

	d.Proof = Proof{}

	return d, nil
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
	multicodecHashed := utils.MulticodecEncode(multicodec.Blake3, hashed[:])

	return multicodecHashed, nil
}
