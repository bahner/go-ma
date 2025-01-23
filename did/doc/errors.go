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

	ErrVerificationMethodIsEmpty           = errors.New("verification method is empty")
	ErrVerificationMethodFetchFailed       = errors.New("error fetching verification method")
	ErrVerificationMethodMissingType       = errors.New("verification method missing type")
	ErrVerificationMethoddUnkownID         = errors.New("verification method ID unknown")
	ErrVerificationMethodInvalidController = errors.New("verification method invalid controller")
	ErrAssertionMethodIsEmpty              = errors.New("assertion method is empty")
	ErrKeyAgreementIsEmpty                 = errors.New("key agreement is empty")
	ErrProofIsEmpty                        = errors.New("proof is empty")

	ErrPublicKeyMultibaseInvalid  = errors.New("failed to decode public key multibase")
	ErrPublicKeyMultibaseEmpty    = errors.New("public key multibase is empty")
	ErrPublicKeyMultibaseMismatch = errors.New("public key multibase mismatch")
	ErrMultiCodecInvalid          = fmt.Errorf("codec must be %s", key.KEY_AGREEMENT_MULTICODEC.String())

	ErrPublicKeyLengthInvalid = fmt.Errorf("public keysize must be %d", ed25519.PublicKeySize)

	ErrDoumentAlreadyPublished = errors.New("document already published")

	ErrContextIsEmpty    = errors.New("context missing")
	ErrIDIsEmpty         = errors.New("id missing")
	ErrControllerIsEmpty = errors.New("controller missing")
	ErrNodeIDEmpty       = errors.New("node ID missing")
	ErrHostTypeMissing   = errors.New("no such node type")
	ErrInvalidHostType   = errors.New("invalid endpoint protocol")
	ErrIdentityIsEmpty   = errors.New("identity missing")
	ErrIdentityInvalid   = errors.New("invalid identity")

	ErrDocumentIsNil            = errors.New("document is nil")
	ErrDocumentInvalid          = errors.New("invalid document")
	ErrDocumentSignatureInvalid = errors.New("invalid document signature")
)
