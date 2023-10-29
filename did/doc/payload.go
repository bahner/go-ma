package doc

import (
	"encoding/json"

	"github.com/bahner/go-ma/did/doc/proof"
)

// Payload generates the unsigned DID,
// This is everything non-metadata in the DID document.
// We don't use a pointer here, so that we don't have to reiterate the
// struct in the function. We just need to change the signature.
func Payload(d Document) (Document, error) {

	d.Proof = []proof.Proof{}

	return d, nil
}

// ToJSON converts the DID to JSON format
func (d *Document) MarshalPayloadToJSON() ([]byte, error) {
	p, err := Payload(*d)
	if err != nil {
		return nil, err
	}

	return json.Marshal(p)
}
