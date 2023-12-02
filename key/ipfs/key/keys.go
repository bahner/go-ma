package key

import (
	"github.com/bahner/go-ma/internal"
	coreiface "github.com/ipfs/kubo/core/coreiface"
)

func List() ([]coreiface.Key, error) {

	return internal.GetIPFSAPI().Key().List(internal.GetContext())
}
