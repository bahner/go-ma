package set

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	ipnskey "github.com/bahner/go-ma/key/ipns"
	cbor "github.com/fxamacker/cbor/v2"
)

type Keyset struct {
	IPNSKey       ipnskey.Key
	EncryptionKey key.EncryptionKey
	SigningKey    key.SigningKey
}

// Creates new keyset from a name (typically fragment of a DID)
func New(name string) (Keyset, error) {

	IPNSKey, err := ipnskey.New(name)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to get or create key in IPFS: %w", err)
	}

	identifier := internal.GetDIDIdentifier(IPNSKey.DID)

	encryptionKey, err := key.NewEncryptionKey(identifier)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate encryption key: %w", err)
	}

	signatureKey, err := key.NewSigningKey(identifier)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate signature key: %w", err)
	}

	return Keyset{
		IPNSKey:       IPNSKey,
		EncryptionKey: encryptionKey,
		SigningKey:    signatureKey,
	}, nil
}

func (k Keyset) MarshalToCBOR() ([]byte, error) {
	return cbor.Marshal(k)
}

func UnmarshalFromCBOR(data []byte) (Keyset, error) {
	var k Keyset
	err := cbor.Unmarshal(data, &k)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/unmarshal: failed to unmarshal keyset: %w", err)
	}
	return k, nil
}

func (k Keyset) Pack() (string, error) {

	data, err := k.MarshalToCBOR()
	if err != nil {
		return "", fmt.Errorf("keyset/pack: failed to marshal keyset: %w", err)
	}

	return internal.MultibaseEncode(data)
}

func Unpack(data string) (Keyset, error) {

	decoded, err := internal.MultibaseDecode(data)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/unpack: failed to decode keyset: %w", err)
	}

	return UnmarshalFromCBOR(decoded)
}
