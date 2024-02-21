package set

import "errors"

var (
	ErrSetKeysMisMatch        = errors.New("keyset: DID and IPFS key DID do not match")
	ErrIPFSKeyNotFound        = errors.New("ipfsKey not found")
	ErrIPFSKeyMissingFragment = errors.New("ipfsKey has no fragment")
	ErrIPFSKeyMissingID       = errors.New("ipfsKey has no ID")
	ErrIPFSKeyMissingName     = errors.New("ipfsKey has no name")
)
