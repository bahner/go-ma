package doc

import (
	"crypto/rsa"
	"fmt"

	"github.com/bahner/go-ma/did/pkm"
)

// This is really useful sugar as these are the keys used for encryption.
// In our case this should normally just return one key.
// We can try and decrypt with each key until we find one that works.
func (d *Document) VerificationMethodRSAPublicKeys() ([]*rsa.PublicKey, error) {

	var pks []*rsa.PublicKey

	for _, method := range d.VerificationMethod {

		pubKey, err := pkm.Parse(method.PublicKeyMultibase)
		if err != nil {
			return pks, fmt.Errorf("doc: Error parsing public key: %s", err)
		}

		rsaPubKey := pubKey.PublicKey.(*rsa.PublicKey)
		pks = append(pks, rsaPubKey)
	}

	return pks, nil
}
