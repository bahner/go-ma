package doc

import (
	"crypto"

	"github.com/multiformats/go-multibase"
)

const (
	SIGNATURE_HASH     = crypto.SHA256
	SIGNATURE_ENCODING = multibase.Base58BTC
)

var Context = []string{"https://w3id.org/did/v1"}
