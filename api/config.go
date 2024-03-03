package api

import (
	"github.com/bahner/go-ma"
	"github.com/spf13/pflag"
)

func init() {

	pflag.String("api-maddr", ma.DEFAULT_IPFS_API_MULTIADDR, "Multiaddr of the IPFS API.")

}
