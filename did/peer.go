package did

import (
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multibase"
	log "github.com/sirupsen/logrus"
)

func (d *DID) PeerID() (peer.ID, error) {
	_, decodedBytes, err := multibase.Decode(d.Identifier)
	if err != nil {
		log.Debugf("(Failed to decode DID %s: %v", d.Identifier, err)
		return "", err
	}

	c, err := cid.Cast(decodedBytes)
	if err != nil {
		return "", err
	}
	return peer.FromCid(c)
}
