package did

import (
	"fmt"
)

var (
	ErrInvalidFragment   = fmt.Errorf("fragment must be a valid NanoID")
	ErrMissingFragment   = fmt.Errorf("missing fragment")
	ErrInvalidDID        = fmt.Errorf("invalid DID")
	ErrEmptyDID          = fmt.Errorf("empty DID")
	ErrInvalidPrefix     = fmt.Errorf("invalid prefix. Must start with  %s", PREFIX)
	ErrInvalidFormat     = fmt.Errorf("invalid DID format, must contain both an identifier and a fragment and nothing else")
	ErrInvalidName       = fmt.Errorf("identifier must be a valid IPNS name")
	ErrMissingIdentifier = fmt.Errorf("missing identifier")
)
