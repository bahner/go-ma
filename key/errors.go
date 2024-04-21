package key

import (
	"errors"
	"fmt"
)

var (
	ErrNoType                         = errors.New("key has no type")
	ErrInvalidAssertionMethodType     = fmt.Errorf("key type must be %s", ASSERTION_METHOD_KEY_TYPE)
	ErrInvalidKeyAgreementType        = fmt.Errorf("key type must be %s", KEY_AGREEMENT_KEY_TYPE)
	ErrNoPrivateKey                   = errors.New("key has no private key")
	ErrNoPublicKey                    = errors.New("key has no public key")
	ErrNoPublicKeyMultibase           = errors.New("key has no public key multibase")
	ErrInvalidPublicKeyMultibase      = errors.New("invalid public key multibase")
	ErrInvalidEncryptionKeyMulticodec = errors.New("invalid encryption key. Must be " + KEY_AGREEMENT_MULTICODEC.String())
	ErrInvalidSigningKeyMulticodec    = errors.New("invalid signing key. Must be " + ASSERTION_METHOD_MULTICODEC.String())
	ErrInvalidKeySize                 = errors.New("invalid private key size")
	ErrPrivateKeyIsNil                = errors.New("key is nil")
)
