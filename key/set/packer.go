package set

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/utils"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/libp2p/go-libp2p/core/crypto"
	log "github.com/sirupsen/logrus"
)

func (ks *Keyset) MarshalToCBOR() ([]byte, error) {
	identityBytes, err := crypto.MarshalPrivateKey(ks.Identity)
	if err != nil {
		return nil, fmt.Errorf("marshal identity key: %w", err)
	}
	temp := struct {
		Identity      []byte
		DID           did.DID
		EncryptionKey key.EncryptionKey
		SigningKey    key.SigningKey
	}{
		Identity:      identityBytes,
		DID:           ks.DID,
		EncryptionKey: ks.EncryptionKey,
		SigningKey:    ks.SigningKey,
	}

	return cbor.Marshal(temp)
}

func UnmarshalFromCBOR(data []byte) (Keyset, error) {
	temp := struct {
		Identity      []byte
		DID           did.DID
		EncryptionKey key.EncryptionKey
		SigningKey    key.SigningKey
	}{}

	err := cbor.Unmarshal(data, &temp)
	if err != nil {
		return Keyset{}, fmt.Errorf("unmarshal keyset: %w", err)
	}

	identity, err := crypto.UnmarshalPrivateKey(temp.Identity)
	if err != nil {
		return Keyset{}, fmt.Errorf("unmarshal identity key: %w", err)
	}

	return Keyset{
		Identity:      identity,
		DID:           temp.DID,
		EncryptionKey: temp.EncryptionKey,
		SigningKey:    temp.SigningKey,
	}, nil
}

func (k Keyset) Pack() (string, error) {

	data, err := k.MarshalToCBOR()
	if err != nil {
		return "", fmt.Errorf("KeysetPack: %w", err)
	}

	return utils.MultibaseEncode(data)
}

func Unpack(data string) (Keyset, error) {

	decoded, err := utils.MultibaseDecode(data)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetUnpack: %w", err)
	}

	keyset, err := UnmarshalFromCBOR(decoded)
	log.Debugf("Unpacked keyset: %v", keyset)

	return keyset, err
}
