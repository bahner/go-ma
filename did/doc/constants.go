package doc

import "time"

var DID_CONTEXT = []string{
	"https://w3id.org/did/v1",
}

const (
	VerificationMethodDefaultTTL = time.Duration(7) * time.Hour * 24
)
