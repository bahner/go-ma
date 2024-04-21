package utils

import "fmt"

var (
	ErrNoInput      = fmt.Errorf("error decoding: insufficient data")
	ErrInvalidSize  = fmt.Errorf("error decoding: invalid varint size")
	ErrUnknownCodec = fmt.Errorf("error obtaining codec name: unknown codec")
)
