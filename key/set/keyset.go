package set

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	"github.com/bahner/go-ma/key/ipfs"
	cbor "github.com/fxamacker/cbor/v2"
	log "github.com/sirupsen/logrus"
)

// KeySet struct the encryption and signing keys are actual keys,
// but the IPFSKey is a reference to the IPFS key and holds names and paths.
// The key itself resides in IPFS.
type Keyset struct {
	DID           *did.DID
	IPFSKey       *ipfs.Key
	EncryptionKey *key.EncryptionKey
	SigningKey    *key.SigningKey
}

// Creates new keyset from a name (typically fragment of a DID)
// This requires that the key is already in IPFS and that IPFS is running.
func GetOrCreate(name string) (*Keyset, error) {

	var err error

	k, err := GetByName(name)
	if err != nil {
		log.Debugf("keyset/get: error message from GetByName: %v", err)
	}

	var ipfsKey *ipfs.Key

	if k == nil {
		ipfsKey, err = ipfs.GetOrCreate(name)
		if err != nil {
			return nil, fmt.Errorf("keyset/new: failed to get or create key in IPFS: %w", err)
		}
		log.Debugf("keyset/new: created new key in IPFS: %v", ipfsKey)
	}

	return NewFromKey(ipfsKey)
}

// This creates a new keyset from an existing IPFS key.
func NewFromKey(k *ipfs.Key) (*Keyset, error) {

	identifier, err := k.RootCID()
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to get root CID from IPFS key: %w", err)
	}

	encryptionKey, err := key.NewEncryptionKey(identifier)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to generate encryption key: %w", err)
	}

	signatureKey, err := key.NewSigningKey(identifier)
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to generate signature key: %w", err)
	}

	d, err := did.GetOrCreate(k.DID())
	if err != nil {
		return nil, fmt.Errorf("keyset/new: failed to get or create DID: %w", err)
	}

	return &Keyset{
		DID:           d,
		IPFSKey:       k,
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

func (k *Keyset) ID() string {
	return k.IPFSKey.ID
}
