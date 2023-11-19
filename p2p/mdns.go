package p2p

import (
	"context"
	"sync"

	"github.com/bahner/go-ma"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	log "github.com/sirupsen/logrus"
)

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

// interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// Initialize the MDNS service
func initMDNS(h host.Host, rendezvous string) chan peer.AddrInfo {
	// register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	// An hour might be a long long period in practical applications. But this is fine for us
	ser := mdns.NewMdnsService(h, rendezvous, n)
	if err := ser.Start(); err != nil {
		panic(err)
	}
	return n.PeerChan
}
func DiscoverMDNSPeers(ctx context.Context, wg *sync.WaitGroup, h host.Host) chan peer.AddrInfo {

	log.Debugf("Discovering MDNS peers for servicename: %s", ma.RENDEZVOUS)
	defer wg.Done()

	peerChan := initMDNS(h, ma.RENDEZVOUS)
	for {
		select {
		case peer := <-peerChan:
			log.Infof("Found MDNS peer: %s connecting", peer.ID.String())

			err := h.Connect(ctx, peer)
			if err != nil {
				log.Debugf("Failed connecting to %s, error: %v\n", peer.ID.String(), err)
			} else {
				log.Infof("Connected to MDNS peer: %s", peer.ID.String())
				log.Info("MDNS peer discovery successful.")
				return nil
			}
		case <-ctx.Done():
			log.Info("Context cancelled, stopping MDNS peer discovery.")
			return nil
		}
	}
}
