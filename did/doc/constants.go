package doc

import (
	"crypto"

	"github.com/multiformats/go-multibase"
)

const (
	CONTEXT            = "https://w3id.org/did/v1"
	SIGNATURE_HASH     = crypto.SHA256
	SIGNATURE_ENCODING = multibase.Base58BTC
)
