package did

import "errors"

var (
	ErrInvalidID       = errors.New("invalid ID")
	ErrInvalidFragment = errors.New("invalid fragment")
	ErrInvalidDID      = errors.New("invalid DID")
)
