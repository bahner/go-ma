package doc

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/libp2p/go-libp2p/core/peer"
)

const hostNumFields = 4

type Host struct {
	ID   string `cbor:"id" json:"id"`
	Type string `cbor:"type" json:"type"`
}

func NewHost(id string, t string) (Host, error) {

	node := Host{
		ID:   id,
		Type: t,
	}

	err := validateHost(node)
	if err != nil {
		return Host{}, fmt.Errorf("doc/NewHost: %w", err)
	}

	return Host{
		ID:   id,
		Type: t,
	}, nil
}

// Takes a libp2p PeerID and sets it as the PeerID of the document.
// This is the PeerID of the node to dial to communicate with the entity.
func (d *Document) SetP2PHost(peerid peer.ID) error {

	host, err := NewHost(peerid.String(), "p2p")
	if err != nil {
		return fmt.Errorf("doc/SetP2PHost: %w", err)
	}

	d.Host = host

	return nil
}

func (d *Document) GetP2PHost() (peer.ID, error) {
	if d.Host.Type != "p2p" {
		return "", ErrHostTypeMissing
	}

	return peer.IDFromBytes([]byte(d.Host.ID))
}

func validateHost(host Host) error {
	if host.Type != "p2p" {
		return ErrInvalidHostType
	}

	_, err := cid.Parse(host.ID)
	if err != nil {
		return fmt.Errorf("doc/validatePeerID: %w", err)
	}

	return nil
}

func buildHostNode(host Host) (ipld.Node, error) {
	nb := basicnode.Prototype.Map.NewBuilder()
	ma, err := nb.BeginMap(hostNumFields)
	if err != nil {
		return nil, err
	}

	ma.AssembleKey().AssignString("id")
	ma.AssembleValue().AssignString(host.ID)

	ma.AssembleKey().AssignString("type")
	ma.AssembleValue().AssignString(host.Type)

	ma.Finish()

	return nb.Build(), nil
}
