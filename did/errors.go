package did

import (
	"errors"
	"fmt"

	"github.com/bahner/go-ma"
)

var (
	ErrInvalidFragment   = errors.New("fragment must be a valid NanoID")
	ErrInvalidDID        = errors.New("invalid DID")
	ErrEmptyDID          = errors.New("empty DID")
	ErrInvalidPrefix     = fmt.Errorf("invalid prefix. Must start with  %s", ma.DID_PREFIX)
	ErrInvalidFormat     = errors.New("invalid DID format, must contain both an identifier and a fragment and nothing else")
	ErrInvalidIdentifier = errors.New("identifier must be a valid IPNS name")
)
