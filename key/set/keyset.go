package set

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	cbor "github.com/fxamacker/cbor/v2"
	iface "github.com/ipfs/boxo/coreiface"
)

type Keyset struct {
	IPNSKey       iface.Key
	EncryptionKey *key.EncryptionKey
	SigningKey    *key.SigningKey
}

// Creates new keyset from a name (typically fragment of a DID)
func New(name string) (*Keyset, error) {

	IPNSKey, err := ipnskey.New(name)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to get or create key in IPFS: %w", err)
	}

	return NewFromKey(IPNSKey)
}

func NewFromKey(IPNSKey iface.Key) (*Keyset, error) {

	identifier := internal.GetDIDIdentifier(IPNSKey.ID().String())

	encryptionKey, err := key.NewEncryptionKey(identifier)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to generate encryption key: %w", err)
	}

	signatureKey, err := key.NewSigningKey(identifier)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to generate signature key: %w", err)
	}

	return &Keyset{
		IPNSKey:       IPNSKey,
		EncryptionKey: encryptionKey,
		SigningKey:    signatureKey,
	}, nil
}

func (k Keyset) MarshalToCBOR() ([]byte, error) {
	return cbor.Marshal(k)
}
func UnmarshalFromCBOR(data []byte) (*Keyset, error) {
	var k *Keyset
	err := cbor.Unmarshal(data, &k)
	if err != nil {
		return nil, fmt.Errorf("keyset/unmarshal: failed to unmarshal keyset: %w", err)
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

func Unpack(data string) (*Keyset, error) {

	decoded, err := internal.MultibaseDecode(data)
	if err != nil {
		return nil, fmt.Errorf("keyset/unpack: failed to decode keyset: %w", err)
	}

	return UnmarshalFromCBOR(decoded)
}

func (k Keyset) GetOrCreateIPNSKeyFromName(name string) iface.Key {

	keys, err := getKeys()
	if err != nil {
		return nil
	}

	for _, key := range keys {
		if key.Name() == name {
			return key
		}
	}

}
