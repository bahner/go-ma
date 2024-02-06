package api

import (
	"github.com/bahner/go-ma"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {

	pflag.String("api-maddr", ma.DEFAULT_IPFS_API_MULTIADDR, "Multiaddr of the IPFS API.")
	viper.SetDefault("api.maddr", ma.DEFAULT_IPFS_API_MULTIADDR)
	viper.BindPFlag("api.maddr", pflag.Lookup("api-maddr"))

}
