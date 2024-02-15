package set

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
)

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
