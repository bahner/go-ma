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
		return Keyset{}, fmt.Errorf("KeysetUnmarshalFromCBOR: %w", err)
	}

	return k, nil
}

func (k Keyset) Pack() (string, error) {

	data, err := k.MarshalToCBOR()
	if err != nil {
		return "", fmt.Errorf("KeysetPack: %w", err)
	}

	return internal.MultibaseEncode(data)
}

func Unpack(data string) (Keyset, error) {

	decoded, err := internal.MultibaseDecode(data)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetUnpack: %w", err)
	}

	return UnmarshalFromCBOR(decoded)
}
