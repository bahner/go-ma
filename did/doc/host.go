package doc

import (
	"fmt"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

const (
	hostNumFields     = 2
	DEFAULT_HOST_TYPE = "p2p"
)

type Host struct {
	ID   string `cbor:"id" json:"id"`
	Type string `cbor:"type" json:"type"`
}

// Creates and validates a new Host object.
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
func (d *Document) SetP2PHost(peerid peer.ID, hostType string) error {

	host, err := NewHost(peerid.String(), hostType)
	if err != nil {
		return fmt.Errorf("doc/SetP2PHost: %w", err)
	}

	d.Host = host

	return nil
}

// Takes a libp2p private key and sets the PeerID of the document to the PeerID of the key.
func (d *Document) SetP2PHostFromPrivateKey(pk crypto.PrivKey, hostType string) error {

	peerid, err := peer.IDFromPrivateKey(pk)
	if err != nil {
		return fmt.Errorf("config.publishDIDDocumentFromKeyset: %w", err)
	}

	return d.SetP2PHost(peerid, hostType)
}

func validateHost(host Host) error {

	if host.ID == "" {
		return ErrHostIDMissing
	}

	if host.Type == "" {
		return ErrHostTypeMissing
	}

	if host.Type != DEFAULT_HOST_TYPE {
		return ErrInvalidHostType
	}

	_, err := peer.Decode(host.ID)
	if err != nil {
		return fmt.Errorf("doc/validateHost: %w", err)
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
