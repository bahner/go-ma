package internal

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bahner/go-ma"
	"github.com/ipfs/kubo/client/rpc"
	maddr "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	"go.deanishe.net/env"
)

func init() {
	IPFSAPIMultiAddrString = env.Get(ma.ENV_IPFS_API_MULTIADDR, ma.DEFAULT_IPFS_API_MULTIADDR)
}

var (
	once                   sync.Once
	IPFSAPIMultiAddrString string
	ipfsAPI                *rpc.HttpApi
)

// initializeApi sets up the ipfsAPI and ipfsAPISocket.
func initializeApi() {

	ipfsAPIMultiAddr, err := maddr.NewMultiaddr(IPFSAPIMultiAddrString)
	if err != nil {
		log.Fatalf("ipfs: failed to parse IPFS API socket: %v", err)
	}

	ipfsAPI, err = rpc.NewApi(ipfsAPIMultiAddr)
	if err != nil {
		log.Fatalf("ipfs: failed to initialize IPFS API: %v", err)
	}

}

func GetIPFSAPI() *rpc.HttpApi {
	once.Do(initializeApi)
	return ipfsAPI
}

func GetIPFSAPIUrl() string {

	const (
		scheme  = "http"
		apiPath = "/api/v0"
	)

	// Split the multiaddr into components
	parts := strings.Split(IPFSAPIMultiAddrString, "/")

	// Extract IP and port
	ip := parts[2]
	port := parts[4]

	// Construct a standard URL
	urlStr := fmt.Sprintf("%s://%s:%s%s", scheme, ip, port, apiPath)
	return urlStr
}
