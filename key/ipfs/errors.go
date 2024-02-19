package ipfs

import "errors"

var (
	ErrKeyNotFoundInIPFS  = errors.New("ipfsKey not found")
	ErrKeyMissingFragment = errors.New("ipfsKey has no fragment")
	ErrKeyMissingID       = errors.New("ipfsKey has no ID")
	ErrKeyMissingName     = errors.New("ipfsKey has no name")
)
