package ipns

import (
	"crypto/rand"
	"fmt"

	"github.com/bahner/go-ma/internal"
	cbor "github.com/fxamacker/cbor/v2"
	"github.com/libp2p/go-libp2p/core/crypto"
)

// The Key.ID is the DID of the entity it's used for.
//
// This Key can have it's secret key exported as text
// for storage on a yellow sticker on your monitor
// where it belongs.
type Key struct {
	DID       string         `cbor:"DID" json:"DID"`
	PrivKey   crypto.PrivKey `cbor:"PrivKey" json:"PrivKey"`
	PublicKey crypto.PubKey  `cbor:"PublicKey" json:"PublicKey"`
}

// New creates a new Key. It generates a new key pair
// and derives the IPNS name from the public key.
// This function does .not require an IPFS node to be running.
func New(name string) (*Key, error) {

	if !internal.IsValidNanoID(name) {
		return nil, fmt.Errorf("key/new: invalid name: %v", name)
	}

	privKey, pubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("key/new: failed to generate key pair: %w", err)
	}

	DID, err := CreateDIDFromPublicKeyAndName(pubKey, name)
	if err != nil {
		return nil, fmt.Errorf("key/new: failed to create ID from public key and name: %w", err)
	}

	return &Key{
		DID:       DID,
		PrivKey:   privKey,
		PublicKey: pubKey,
	}, nil
}

// MarshalCBOR customizes the CBOR marshaling for IPNSKey.
func (k Key) MarshalCBOR() ([]byte, error) {
	// First, marshal the private key to bytes.
	privKeyBytes, err := crypto.MarshalPrivateKey(k.PrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	pubKeyBytes, err := crypto.MarshalPublicKey(k.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	// Create a map that represents the struct's data.
	dataMap := map[string]interface{}{
		"DID":       k.DID,
		"PrivKey":   privKeyBytes,
		"PublicKey": pubKeyBytes,
	}

	return cbor.Marshal(dataMap)
}

// UnmarshalCBOR customizes the CBOR unmarshaling for Key.
func (k *Key) UnmarshalCBOR(data []byte) error {
	var dataMap map[string]interface{}
	if err := cbor.Unmarshal(data, &dataMap); err != nil {
		return fmt.Errorf("failed to unmarshal CBOR to map: %w", err)
	}

	// Extract and unmarshal the private key.
	privKeyData, privOk := dataMap["PrivKey"].([]byte)
	if !privOk {
		return fmt.Errorf("private key bytes are not in expected format")
	}
	privKey, err := crypto.UnmarshalPrivateKey(privKeyData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal private key: %w", err)
	}

	// Extract and unmarshal the public key.
	pubKeyData, pubOk := dataMap["PublicKey"].([]byte)
	if !pubOk {
		return fmt.Errorf("public key bytes are not in expected format")
	}
	pubKey, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal public key: %w", err)
	}

	did, didOk := dataMap["DID"].(string)
	if !didOk {
		return fmt.Errorf("DID is not in expected format or not present")
	}

	k.DID = did
	k.PrivKey = privKey
	k.PublicKey = pubKey
	return nil
}

func (k *Key) Exists() bool {

	identifier := internal.GetDIDIdentifier(k.DID)
	fragment := internal.GetDIDFragment(k.DID)

	return KeyExists(fragment, identifier)
}
