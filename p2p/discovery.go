package p2p

import (
	"context"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
	log "github.com/sirupsen/logrus"
)

// StartPeerDiscovery starts peer discovery using  DHT
// Try and be smart and exist, when one of the discovery methods is done
func StartPeerDiscovery(ctx context.Context, h host.Host) error {
	log.Debug("Starting peer discovery...")

	// Create a new cancellable context that inherits the timeout from ctx
	ctx, cancel := context.WithCancel(ctx)
	defer cancel() // Ensure any remaining operations are cancelled upon return

	done := make(chan bool, 2)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		DiscoverDHTPeers(ctx, wg, h)
		done <- true
		cancel() // Cancel the other discovery method
	}()

	// Wait for either discovery method to complete or a timeout
	select {
	case <-ctx.Done():
		log.Warn("Peer discovery timed out or was cancelled.")
		return ctx.Err()
	case <-done:
		log.Info("Peer discovery successfully completed.")
		return nil
	}
}
