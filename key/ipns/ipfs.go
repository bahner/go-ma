package ipns

import (
	"bytes"
	"fmt"

	"github.com/bahner/go-ma/internal"
	"github.com/libp2p/go-libp2p/core/crypto"
)

// Import the key into IPFS under it's IPNS name.
// Doesn't try to be clever. If the same is already
// there - do nothing. If a key with the same name exist
// then fail. User will have to delete it manually or choose
// a different name.
func (i *Key) ExportToIPFS(name string) error {

	privKeyBytes, err := crypto.MarshalPrivateKey(i.PrivKey)
	if err != nil {
		return fmt.Errorf("failed to marshal private key: %v", err)
	}

	keyReader := bytes.NewReader(privKeyBytes)

	// Get the key from IPFS
	shell := internal.GetShell()
	err = shell.KeyImport(internal.GetContext(), name, keyReader)
	if err != nil {
		return fmt.Errorf("failed to import key: %v", err)
	}

	return nil

}
