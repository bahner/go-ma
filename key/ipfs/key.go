package ipfs

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/key/ipfs/key"
	cbor "github.com/fxamacker/cbor/v2"
	coreiface "github.com/ipfs/kubo/core/coreiface"
)

type Key struct {
	// The IPNS name of the key, not the local name
	IPNSName string
	// The ID used by kubo to identify the key
	ID string
	// Fragment is the local name of the key, which we use as the DID fragment
	Fragment string
}

// UnmarshalCBOR customizes the CBOR unmarshaling for Key.
func UnmarshalCBOR(data []byte) (*Key, error) {

	var k *Key

	if err := cbor.Unmarshal(data, &k); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CBOR to map: %w", err)
	}

	return k, nil
}

func (k *Key) Exists() bool {

	nk, err := key.LookupName(k.Fragment)
	if err != nil {
		return false
	}

	ik, err := key.LookupID(k.ID)
	if err != nil {
		return false
	}

	return nk.Name() == k.Fragment && ik.ID().String() == k.ID
}

func GetOrCreate(name string) (*Key, error) {

	var err error

	ik, _ := key.LookupName(name)

	if ik == nil {
		ik, err = key.GetOrCreate(name)
		if err != nil {
			return nil, fmt.Errorf("ipfs: failed to create key %s: %w", name, err)
		}
	}

	return NewFromIPFSKey(ik)
}

// Creates a new key in IPFS and returns a Key struct.
// This does not check that the key already exists, so
// the provided DID must be verified by the caller.
// Use GetOrCreate to create a key if it doesn't exist.

func New(didStr string) (*Key, error) {

	// This becomes a validator function
	keyDID, err := did.New(didStr)
	if err != nil {
		return nil, fmt.Errorf("ipfs: failed to create DID from name %s: %w", didStr, err)
	}
	return NewFromDID(keyDID)
}

// This takes an actual *DID as input, not just the string.
func NewFromDID(d *did.DID) (*Key, error) {

	ik, err := key.LookupName(d.Fragment)
	if err != nil {
		return nil, fmt.Errorf("ipfs: failed to lookup key %s: %w", d.Fragment, err)
	}

	return &Key{
		ID:       ik.ID().String(),
		Fragment: ik.Name(),
	}, nil
}

func NewFromIPFSKey(k coreiface.Key) (*Key, error) {

	return &Key{
		IPNSName: k.Path().Segments()[1],
		ID:       k.ID().String(),
		Fragment: k.Name(),
	}, nil
}

func (k *Key) DID() string {

	return ma.DID_PREFIX + k.IPNSName + "#" + k.Fragment
}
