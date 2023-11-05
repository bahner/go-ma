package doc

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
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

// Marshals the payload to JSON for inspection.
func (d *Document) MarshalPayloadToJSON() ([]byte, error) {
	p, err := Payload(*d)
	if err != nil {
		return nil, err
	}

	return json.Marshal(p)
}

func (d *Document) PayloadHash() ([]byte, error) {
	// Etxract the payload from the document as a JSON string
	p, err := d.MarshalPayloadToCBOR()
	if err != nil {
		return nil, fmt.Errorf("doc hashing: Error marshalling payload to JSON: %s", err)
	}

	// Hash the JSON string
	hashed := blake3.Sum256(p)
	multicodecHashed, err := internal.MulticodecEncode(ma.HASH_ALGORITHM_MULTICODEC_STRING, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("doc sign: Error multiencoding hashed payload: %s", err)
	}

	return multicodecHashed, nil
}
