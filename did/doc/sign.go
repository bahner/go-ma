package doc

import (
	"crypto/ed25519"
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/multiformats/go-multibase"
)

func (doc *Document) Sign(signKey *ed25519.PrivateKey) error {
	p, err := doc.MarshalPayloadToJSON()
	if err != nil {
		return err
	}

	// Sign the payload with an ed25519 key
	signature := ed25519.Sign(*signKey, p)
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
