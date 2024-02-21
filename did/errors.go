package did

import (
	"errors"
	"fmt"

	"github.com/bahner/go-ma"
)

var (
	ErrInvalidFragment   = errors.New("fragment must be a valid NanoID")
	ErrMissingFragment   = errors.New("missing fragment")
	ErrInvalidDID        = errors.New("invalid DID")
	ErrEmptyDID          = errors.New("empty DID")
	ErrDIDIsNil          = errors.New("DID is nil")
	ErrInvalidPrefix     = fmt.Errorf("invalid prefix. Must start with  %s", ma.DID_PREFIX)
	ErrInvalidFormat     = errors.New("invalid DID format, must contain both an identifier and a fragment and nothing else")
	ErrInvalidIdentifier = errors.New("identifier must be a valid IPNS name")
	ErrMissingIdentifier = errors.New("missing identifier")
	ErrIPFSKeyNotFound   = errors.New("ipfsKey not found")
)
