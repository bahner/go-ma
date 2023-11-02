package key

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
)

type Keyset struct {
	IPNSKey       IPNSKey
	EncryptionKey EncryptionKey
	SigningKey    SigningKey
}

// Creates new keyset from a name (typically fragment of a DID)
func NewKeyset(name string) (Keyset, error) {
	encryptionKey, err := GenerateEncryptionKey(name)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate encryption key: %w", err)
	}

	signatureKey, err := GenerateSigningKey(name)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to generate signature key: %w", err)
	}

	ipfsKey, err := NewIPNSKey(name)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/new: failed to get or create key in IPFS: %w", err)
	}

	return Keyset{
		IPNSKey:       ipfsKey,
		EncryptionKey: encryptionKey,
		SigningKey:    signatureKey,
	}, nil
}

func (k Keyset) MarshalToCBOR() ([]byte, error) {
	return cbor.Marshal(k)
}

func UnmarshalKeysetFromCBOR(data []byte) (Keyset, error) {
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

func UnpackKeyset(data string) (Keyset, error) {

	decoded, err := internal.MultibaseDecode(data)
	if err != nil {
		return Keyset{}, fmt.Errorf("keyset/unpack: failed to decode keyset: %w", err)
	}

	return UnmarshalKeysetFromCBOR(decoded)
}
