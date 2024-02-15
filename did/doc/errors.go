package doc

import (
	"crypto/ed25519"
	"errors"
	"fmt"

	"github.com/bahner/go-ma/key"
)

var (
	ErrPayloadMarshal     = errors.New("error marshalling payload to CBOR")
	ErrPayloadMultiencode = errors.New("error multiencoding hashed payload")
	ErrPayloadSign        = errors.New("error signing payload")

	ErrVerificationMethodFetchFailed       = errors.New("error fetching verification method")
	ErrVerificationMethodMissingType       = errors.New("verification method missing type")
	ErrVerificationMethoddUnkownID         = errors.New("verification method ID unknown")
	ErrVerificationMethodInvalidController = errors.New("verification method invalid controller")

	ErrPublicKeyMultibaseInvalid  = errors.New("failed to decode public key multibase")
	ErrPublicKeyMultibaseEmpty    = errors.New("public key multibase is empty")
	ErrPublicKeyMultibaseMismatch = errors.New("public key multibase mismatch")
	ErrMultiCodecInvalid          = fmt.Errorf("codec must be %s", key.ASSERTION_METHOD_KEY_MULTICODEC_STRING)

	ErrPublicKeyLengthInvalid = fmt.Errorf("public keysize must be %d", ed25519.PublicKeySize)

	ErrDoumentAlreadyPublished = errors.New("document already published")

	ErrControllersIsEmpty = errors.New("controller missing")

	ErrDocumentIsNil          = errors.New("document is nil")
	ErrDocumentInvalid          = errors.New("invalid document")
	ErrDocumentSignatureInvalid = errors.New("invalid document signature")
)
