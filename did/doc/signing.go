package doc

import (
	"crypto/rand"
	"crypto/rsa"
)

func (doc *Document) Sign(privKey *rsa.PrivateKey) error {
	p, err := doc.MarshalPayloadToJSON()
	if err != nil {
		return err
	}

	// Sign the marshalled payload.
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, SIGNATURE_HASHER, p)
	if err != nil {
		return err
	}

	doc.Signature = string(signature)

	return nil
}
