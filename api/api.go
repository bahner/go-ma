package api

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ipfs/kubo/client/rpc"
	maddr "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Default IPFS API multiaddr. Default to Brave's IPFS companion extension,
// which doesn't require any special configuration or installation.
// Kubo is "/ip4/127.0.0.1/tcp/5001"

const DEFAULT_IPFS_API_MADDR = "/ip4/127.0.0.1/tcp/5001"

var (
	once    sync.Once
	ipfsAPI *rpc.HttpApi
)

func init() {

	// This must always be first as it's used everywhere else.
	pflag.String("api-maddr", DEFAULT_IPFS_API_MADDR, "Multiaddr of the IPFS API.")
	viper.SetDefault("api.maddr", DEFAULT_IPFS_API_MADDR)
	viper.BindPFlag("api.maddr", pflag.Lookup("api-maddr"))
}

func GetIPFSAPI() *rpc.HttpApi {

	// Only initialize the API once, then just return it later.
	once.Do(func() {

		ipfsAPIMultiAddr, err := maddr.NewMultiaddr(viper.GetString("api.maddr"))
		if err != nil {
			log.Fatalf("ipfs: failed to parse IPFS API socket: %v", err)
		}

		ipfsAPI, err = rpc.NewApi(ipfsAPIMultiAddr)
		if err != nil {
			log.Fatalf("ipfs: failed to initialize IPFS API: %v", err)
		}

	})

	return ipfsAPI
}

func GetIPFSAPIUrl() string {

	const (
		scheme  = "http"
		apiPath = "/api/v0"
	)

	// Split the multiaddr into components
	parts := strings.Split(viper.GetString("api.maddr"), "/")

	// Extract IP and port
	ip := parts[2]
	port := parts[4]

	// Construct a standard URL
	urlStr := fmt.Sprintf("%s://%s:%s%s", scheme, ip, port, apiPath)
	return urlStr
}
