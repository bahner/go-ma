package api

import (
	"github.com/spf13/pflag"
)

func init() {

	pflag.String("api-maddr", defaultIPFSAPIMaddr, "Multiaddr of the IPFS API.")

}
