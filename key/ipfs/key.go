package ipfs

import (
	"fmt"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key/ipfs/key"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/ipfs/boxo/path"
	"github.com/libp2p/go-libp2p/core/peer"
)

type Key struct {
	ID   string    `cbor:"id"`
	Name string    `cbor:"name"`
	Path path.Path `cbor:"path"`
	DID  string    `cbor:"did"`
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

	nk, err := key.LookupName(k.Name)
	if err != nil {
		return false
	}

	ik, err := key.LookupID(k.ID)
	if err != nil {
		return false
	}

	return nk.Name() == k.Name && ik.ID().String() == k.ID
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

	return New(ma.DID_PREFIX + ik.ID().String() + "#" + name)
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

	// This becomes a validator function
	pid, err := peer.Decode(internal.GetDIDIdentifier(didStr))
	if err != nil {
		return nil, fmt.Errorf("ipfs: failed to decode peer ID %s: %w", didStr, err)
	}

	// Path functions both as a validator and a cleaner.
	keyPath, err := path.NewPathFromSegments(path.IPNSNamespace, pid.String())
	if err != nil {
		return nil, fmt.Errorf("ipfs: failed to create path from key identifier %s: %w", keyDID.Identifier, err)
	}

	return &Key{
		ID:   pid.String(),
		Name: keyDID.Fragment,
		Path: keyPath,
		DID:  didStr,
	}, nil
}
