package set

import "errors"

var (
	ErrSetKeysMisMatch = errors.New("keyset: DID and IPFS key DID do not match")
)
