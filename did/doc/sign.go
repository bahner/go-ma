package doc

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	"github.com/multiformats/go-multibase"
)

func (doc *Document) Sign(signKey key.SignatureKey) error {
	p, err := doc.MarshalPayloadToJSON()
	if err != nil {
		return err
	}

	// Sign the payload with an ed25519 key
	signature, err := signKey.Sign(p)
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error signing payload: %s\n", err))
	}

	multibaseEncodedSignature, err := multibase.Encode(ma.MULTIBASE_ENCODING, signature)
	if err != nil {
		return internal.LogError(fmt.Sprintf("doc sign: Error encoding signature: %s\n", err))
	}

	doc.Signature = multibaseEncodedSignature

	return nil
}
