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

	// Wait for the wait group or the timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		log.Warn("Peer discovery timed out.")
		return ctx.Err()
	case <-done:
		log.Info("Peer discovery successfully completed.")
		return nil
	}
}
