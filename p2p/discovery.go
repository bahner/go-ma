package p2p

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	log "github.com/sirupsen/logrus"
)

func StartPeerDiscovery(ctx context.Context, h host.Host) error {

	log.Debug("Starting peer discovery...")

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go DiscoverDHTPeers(ctx, wg, h)
	go DiscoverMDNSPeers(ctx, wg, h)
	wg.Wait()
	log.Info("Peer discovery finished.")

	return nil
}
