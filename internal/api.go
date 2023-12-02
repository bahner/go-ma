package internal

import (
	"sync"

	"github.com/ipfs/kubo/client/rpc"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	"go.deanishe.net/env"
)

const defaultIPFSAPIMultiAddr = "/ipv/tcp/127.0.0.1/45005" // Default to Brave browser, Kubo is localhost:5001

var (
	once                   sync.Once
	IPFSAPIMultiAddrString string
	ipfsAPI                *rpc.HttpApi
)

// initializeApi sets up the ipfsAPI and ipfsAPISocket.
func initializeApi() {
	IPFSAPIMultiAddrString = env.Get("IPFS_API_MULTIADDR", defaultIPFSAPIMultiAddr)

	ipfsAPIMultiAddr, err := ma.NewMultiaddr(IPFSAPIMultiAddrString)
	if err != nil {
		log.Fatalf("ipfs: failed to parse IPFS API socket: %v", err)
	}

	ipfsAPI, err = rpc.NewApi(ipfsAPIMultiAddr)
	if err != nil {
		log.Fatalf("ipfs: failed to initialize IPFS API: %v", err)
	}

}

func GetIPSAPI() *rpc.HttpApi {
	once.Do(initializeApi)
	return ipfsAPI
}
