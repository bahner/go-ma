package ipns

import (
	"bytes"
	"fmt"

	"github.com/bahner/go-ma/internal"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/libp2p/go-libp2p/core/crypto"
	log "github.com/sirupsen/logrus"
)

// Import the key into IPFS under it's IPNS name.
// Doesn't try to be clever. If the same is already
// there - do nothing. If a key with the same name exist
// then fail. User will have to delete it manually or choose
// a different name.
// Takes 2nd parameters which is to forceUpdate by deleting
// existing keys with the same name.
func (i *Key) ExportToIPFS(forceUpdate bool) error {

	name := internal.GetDIDFragment(i.DID)
	identifier := internal.GetDIDIdentifier(i.DID)

	if i.Exists() {
		log.Debugf("key/ipns: key with name %s and identifier %s already exists. Nothing done.", name, identifier)
		return nil
	}

	// If an existing key with the same identifier exists, at this point
	// we know that it belongs to a different name / alias. That is an error.
	// Inform the user of the existing key's name as a courtesy.
	if !i.IsUnique() {

		existingKey, err := GetKeyByIdentifier(internal.GetDIDIdentifier(i.DID))
		if err != nil {
			return fmt.Errorf("key/ipns: failed to get key by identifier: %v", err)
		}

		return fmt.Errorf(
			"key/ipns: key with identifier %s already exists with another name: %s",
			internal.GetDIDIdentifier(i.DID),
			existingKey.Name,
		)
	}

	// If an existing key with the same name exists, but it has a different
	// identifier, then we must require forceUpdate to be true.
	if KeyWithNameExists(name) && !forceUpdate {
		return fmt.Errorf("key/ipns: key with name %s already exists and publication isn't forced", name)
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
	// Doublecheck the forceUpdate flag. This is redundant,
	// but it's better to be safe than sorry.
	if KeyWithNameExists(name) && forceUpdate {
		_, err = shell.KeyRm(internal.GetContext(), name)
		if err != nil {
			return fmt.Errorf("failed to delete existing key: %v", err)
		}
		log.Infof("key/ipns: deleted existing key with name %s because of forced publication", name)
	}

	// Import the key into IPFS
	err = shell.KeyImport(internal.GetContext(), name, keyReader)
	if err != nil {
		return fmt.Errorf("failed to import key: %v", err)
	}
	log.Debugf("key/ipns: imported key with name %s", name)

	return nil

}

// Checks if another IPNS key with either the same name or identifier exists in the IPFS daemon.
func (i *Key) IsUnique() bool {

	if KeyWithNameExists(internal.GetDIDFragment(i.DID)) {
		return false
	}

	if KeyWithIdentifierExists(internal.GetDIDIdentifier(i.DID)) {
		return false
	}

	return true

}

// Checks if a key with the same name/alias exists in the IPFS daemon.
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

// Checks if a key with the same identifier exists in the IPFS daemon.
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

// Fetch a key by it's name / alias.
func GetKeyByName(name string) (*shell.Key, error) {

	shell := internal.GetShell()
	ctx := internal.GetContext()

	// Get the key from IPFS
	shellKeys, err := shell.KeyList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get key list: %v", err)
	}

	for _, ipnsKey := range shellKeys {
		if ipnsKey.Name == name {
			return ipnsKey, nil
		}
	}

	return nil, fmt.Errorf("key with name %s not found", name)
}

// Fetch a key by it's identifier.
func GetKeyByIdentifier(identifier string) (*shell.Key, error) {

	shell := internal.GetShell()
	ctx := internal.GetContext()

	// Get the key from IPFS
	shellKeys, err := shell.KeyList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get key list: %v", err)
	}

	for _, ipnsKey := range shellKeys {
		if ipnsKey.Id == identifier {
			return ipnsKey, nil
		}
	}

	return nil, fmt.Errorf("key with identifier %s not found", identifier)
}

// Check if a key with the same name and identifier exists.
func KeyExists(name string, identifier string) bool {

	log.Debugf("key/ipns: checking if key with name %s and identifier %s exists", name, identifier)
	// Check if the key already exists. Return OK if it does.
	existingKey, err := GetKeyByName(name)

	if err == nil {
		// This means that an existing key with the same name exists.

		// If the existing key has the same identifier, then we're good.
		if existingKey.Id == identifier {
			log.Debugf("key/ipns: key with name %s and identifier %s already exists IN IPFS", name, identifier)
			return true
		}
	}

	return false
}
