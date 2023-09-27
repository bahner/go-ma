package doc

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	"github.com/multiformats/go-multibase"
)

func (doc *Document) Sign(signKey *key.Key) error {
	p, err := doc.MarshalPayloadToJSON()
	if err != nil {
		return err
	}

	// Compute the hash of the payload
	h := SIGNATURE_HASH.New()
	h.Write(p)
	hashed := h.Sum(nil)

	// Sign the hash of the payload
	signature, err := rsa.SignPKCS1v15(rand.Reader, signKey.RSAPrivateKey, SIGNATURE_HASH, hashed)
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error signing payload: %s", err))
	}

	multibaseEncodedSignature, err := multibase.Encode(SIGNATURE_ENCODING, signature)
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error encoding signature: %s", err))
	}

	doc.Signature = multibaseEncodedSignature

	return nil
}
