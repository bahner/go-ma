package msg

import (
	"context"
	"fmt"

	"github.com/bahner/go-ma/did/doc"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/ipfs/boxo/namesys"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"golang.org/x/crypto/curve25519"
)

// Bask the encrypted message and the encrypted symmetric key in a CBOR envelope.
type Envelope struct {
	EphemeralKey     []byte
	EncryptedContent []byte
	EncryptedHeaders []byte
}

func (e *Envelope) Verify() error {
	if e.EphemeralKey == nil || e.EncryptedContent == nil || e.EncryptedHeaders == nil {
		return fmt.Errorf("envelope: missing fields in envelope")
	}

	if len(e.EphemeralKey) != curve25519.PointSize {
		return fmt.Errorf("envelope: invalid ephemeral key length")
	}

	if e.EncryptedContent == nil {
		return fmt.Errorf("envelope: missing encrypted content")
	}

	if e.EncryptedHeaders == nil {
		return fmt.Errorf("envelope: missing encrypted headers")
	}

	return nil
}

// Use a pointer here, this might be arbitrarily big.
func (m *Message) Enclose(resolver *namesys.Resolver) (*Envelope, error) {

	err := m.Verify()
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: %w", err)
	}

	// AT this point we *need* to fetch the recipient's document, otherwise we can't encrypt the message.
	// But this fetch should probably have a timeout, so we don't get stuck here - or a caching function.
	to, _, err := doc.Fetch(m.To)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: %w", err)
	}

	recipientPublicKeyBytes, err := to.KeyAgreementPublicKeyBytes()
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: %w", err)
	}

	// Generate ephemeral keys to be used for his message
	ephemeralPublic, symmetricKey, err := generateEphemeralKeys(recipientPublicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: %w", err)
	}

	msgHeaders, err := m.marshalHeadersToCBOR()
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: %w", err)
	}

	encryptedMsgHeaders, err := encrypt(msgHeaders, symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: %w", err)
	}

	encryptedContent, err := encrypt(m.Content, symmetricKey)
	if err != nil {
		return nil, fmt.Errorf("msg_enclose: %w", err)
	}

	return &Envelope{
		EphemeralKey:     ephemeralPublic,
		EncryptedHeaders: encryptedMsgHeaders,
		EncryptedContent: encryptedContent,
	}, nil
}

func (e *Envelope) Send(ctx context.Context, t *pubsub.Topic) error {

	eBytes, err := cbor.Marshal(e)
	if err != nil {
		return fmt.Errorf("msg_send: %w", err)
	}

	// t.Publish(ctx, eBytes, nil)
	t.Publish(ctx, eBytes)

	return nil
}

// Takes the envelope as a byte array and returns a pointer to an Envelope struct
// Basically this is what you do with a receieved message envelope, eg. in an Open() function.
func UnmarshalEnvelopeFromCBOR(data []byte) (*Envelope, error) {

	e := &Envelope{}

	err := cbor.Unmarshal(data, e)
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_envelope:  %s", err)
	}

	return e, nil
}

func UnmarshalAndVerifyEnvelopeFromCBOR(data []byte) (*Envelope, error) {

	e, err := UnmarshalEnvelopeFromCBOR(data)
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_and_verify_envelope: %s", err)
	}

	if e == nil {
		return nil, fmt.Errorf("msg_unmarshal_and_verify_envelope: %w", ErrNilEnvelope)
	}

	err = e.Verify()
	if err != nil {
		return nil, fmt.Errorf("msg_unmarshal_and_verify_envelope: %s", err)
	}

	return e, nil
}

func (e *Envelope) getContent(privkey []byte) ([]byte, error) {
	return decrypt(e.EncryptedContent, e.EphemeralKey, privkey)
}

func (e *Envelope) getHeaders(privkey []byte) (*Headers, error) {

	bytes, err := decrypt(e.EncryptedHeaders, e.EphemeralKey, privkey)
	if err != nil {
		return nil, err
	}

	var hdrs *Headers = new(Headers)

	err = cbor.Unmarshal(bytes, hdrs)
	if err != nil {
		return nil, err
	}

	return hdrs, nil
}
