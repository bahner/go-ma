package set

import (
	"fmt"

	"github.com/bahner/go-ma/multi"
	cbor "github.com/fxamacker/cbor/v2"
	log "github.com/sirupsen/logrus"
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

	log.Debugf("Unmarshaled keyset: %v", k)

	return k, nil
}

func (k Keyset) Pack() (string, error) {

	data, err := k.MarshalToCBOR()
	if err != nil {
		return "", fmt.Errorf("KeysetPack: %w", err)
	}

	return multi.MultibaseEncode(data)
}

func Unpack(data string) (Keyset, error) {

	decoded, err := multi.MultibaseDecode(data)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetUnpack: %w", err)
	}

	keyset, err := UnmarshalFromCBOR(decoded)
	log.Debugf("Unpacked keyset: %v", keyset)

	return keyset, err
}
