package internal

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	log "github.com/sirupsen/logrus"
)

const defaultIPFSAPISocket = "localhost:45005" // Default to Brave browser, Kubo is localhost:5001

// FIXME: This is taken completely out of the blue. They must be properly researched.

// This is the life time of the IPNS record.
const defaultIPNSRecordLifetime = time.Duration(24 * time.Hour) // A day. The default

// This is the max cache time of the IPNS record.
const defaultIPNSRecordTTL = time.Duration(1 * time.Hour) // An hour.

var ctx = context.Background()

var (
	once          sync.Once
	ipfsAPISocket string
	ipfsAPI       *shell.Shell
	exists        bool
)

// initializeApi sets up the ipfsAPI and ipfsAPISocket.
func initializeApi() {
	ipfsAPISocket, exists = os.LookupEnv("IPFS_API_SOCKET")
	if !exists {
		ipfsAPISocket = defaultIPFSAPISocket
	}
	ipfsAPI = shell.NewShell(ipfsAPISocket)

	_, err := ipfsAPI.ID()
	if err != nil {
		log.Fatalf("ipfs: failed to connect to IPFS API: %v\n", err)
	}
}

// PublishToIPFS publishes the provided data string to IPFS and returns the CID.
// This is a simple function which only publishes strings, which is why
// it is in the internal package.

func IPFSPublishString(data string) (string, error) {
	once.Do(initializeApi)

	cid, err := ipfsAPI.Add(strings.NewReader(data))
	if err != nil {
		log.Printf("ipfs: failed to add data to IPFS: %v\n", err)
		return "", err
	}

	return cid, nil
}

func IPLDPutJSON(data []byte) (string, error) {

	return IPLDPut(data, "json", "dag-cbor")

}

func IPLDPutCBOR(data []byte) (string, error) {

	return IPLDPut(data, "cbor", "dag-cbor")

}

func IPLDPut(data []byte, input string, output string) (string, error) {

	once.Do(initializeApi)

	cid, err := ipfsAPI.DagPut(data, input, output)

	if err != nil {
		log.Printf("ipld: failed to add data IPLD linked data: %v\n", err)
		return "", err
	}

	return cid, nil

}

// // Now if ever there was a sugar function.
// func IPFSPublishBytes(data []byte) (string, error) {

//		return IPFSPublishString(string(data))
//	}
func IPNSPublishCID(contentHash string, key string, resolve bool) (*shell.PublishResponse, error) {
	once.Do(initializeApi)

	// res, err := ipfsAPI.PublishWithDetails(contentHash, key, lifetime, ttl, resolve)
	// if err != nil {
	// 	return nil, LogError("ipfs: failed to publish to IPNS: %v", err)
	// }

	return ipfsAPI.PublishWithDetails(contentHash,
		key, // Alias name for IPNS key
		defaultIPNSRecordLifetime,
		defaultIPNSRecordTTL,
		resolve)
}
func IPNSListKeys() ([]*shell.Key, error) {
	once.Do(initializeApi)

	return ipfsAPI.KeyList(ctx)
}
func IPNSFindKeyID(name string) (string, error) {

	keys, err := IPNSListKeys()
	if err != nil {
		return "", fmt.Errorf("ipfs: failed to list : %v", err)
	}
	for _, key := range keys {
		if key.Name == name {
			return key.Id, nil
		}
	}

	return "", fmt.Errorf("ipfs: key %s not found", name)
}
func IPNSFindKeyName(id string) (string, error) {

	keys, err := IPNSListKeys()
	if err != nil {
		return "", fmt.Errorf("ipfs: failed to list : %v", err)
	}
	for _, key := range keys {
		if key.Id == id {
			return key.Name, nil
		}
	}

	return "", fmt.Errorf("ipfs: key %s not found", id)
}
func IPNSLookupKeyName(keyName string) (*shell.Key, error) {

	var lookedupKey *shell.Key

	keys, err := IPNSListKeys()
	if err != nil {
		return lookedupKey, fmt.Errorf("ipfs: failed to list : %v", err)
	}

	// A little deeper than I usually like to nest, but hey, it's a one off.
	for _, foundKey := range keys {
		if foundKey.Name == keyName {
			return foundKey, nil
		}
	}

	return lookedupKey, fmt.Errorf("ipfs: key %s not found", keyName)
}
func GetShellKey(keyName string) (*shell.Key, error) {
	once.Do(initializeApi)

	var key *shell.Key

	key, err := IPNSLookupKeyName(keyName)
	if err == nil {
		log.Debugf("ipfs: found existing key for %s ", keyName)
		return key, nil
	}

	return nil, fmt.Errorf("ipfs: failed to find key %s", keyName)
}

func GetShell() *shell.Shell {
	once.Do(initializeApi)
	return ipfsAPI
}

func GetContext() context.Context {
	return ctx
}
