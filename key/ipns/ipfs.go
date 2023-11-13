package ipns

import (
	"bytes"
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/libp2p/go-libp2p/core/crypto"
	log "github.com/sirupsen/logrus"
)

// Import the key into IPFS under it's IPNS name.
// Doesn't try to be clever. If the same is already
// there - do nothing. If a key with the same name exist
// then fail. User will have to delete it manually or choose
// a different name.
// Takes 2nd parameters which is to forcePublish by deleting
// existing keys with the same name.
func (i *Key) ExportToIPFS(name string, forcePublish bool) error {

	if KeyWithNameExists(name) && !forcePublish {
		return fmt.Errorf("key/ipns: key with name %s already exists and publication not forced", name)
	}

	// This test doesn't necessarily mean that it's the same key,
	// that we wanna write to, but it's a warning sign.
	if KeyWithIdentifierExists(internal.GetDIDIdentifier(i.DID)) && !forcePublish {
		return fmt.Errorf(
			"key/ipns: key with identifier %s already exists and publication not forced",
			internal.GetDIDIdentifier(i.DID))
	}

	// Serialise the private key to bytes
	privKeyBytes, err := crypto.MarshalPrivateKey(i.PrivKey)
	if err != nil {
		return fmt.Errorf("failed to marshal private key: %v", err)
	}

	// Create a reader from the bytes, as the publication requires a file object
	keyReader := bytes.NewReader(privKeyBytes)

	// Get the key from IPFS
	shell := internal.GetShell()

	// If key with the same name exists, delete it.
	// Doublecheck the forcePublish flag. This is redundant,
	// but it's better to be safe than sorry.
	if KeyWithNameExists(name) && forcePublish {
		_, err = shell.KeyRm(internal.GetContext(), name)
		if err != nil {
			return fmt.Errorf("failed to delete existing key: %v", err)
		}
		log.Infof("key/ipns: deleted existing key with name %s", name)
	}

	err = shell.KeyImport(internal.GetContext(), name, keyReader)
	if err != nil {
		return fmt.Errorf("failed to import key: %v", err)
	}
	log.Debugf("key/ipns: imported key with name %s", name)

	return nil

}

func (i *Key) IsUnique() bool {

	if KeyWithNameExists(internal.GetDIDFragment(i.DID)) {
		return false
	}

	if KeyWithIdentifierExists(internal.GetDIDIdentifier(i.DID)) {
		return false
	}

	return true

}

func KeyWithNameExists(name string) bool {

	shell := internal.GetShell()

	// Get the key from IPFS
	keySlice, err := shell.KeyList(internal.GetContext())
	if err != nil {
		return false
	}

	for _, ipnsKey := range keySlice {
		if ipnsKey.Name == name {
			return true
		}
	}

	return false
}

func KeyWithIdentifierExists(identifier string) bool {

	shell := internal.GetShell()

	// Get the key from IPFS
	keySlice, err := shell.KeyList(internal.GetContext())
	if err != nil {
		return false
	}

	for _, ipnsKey := range keySlice {
		if ipnsKey.Id == identifier {
			return true
		}
	}

	return false
}
