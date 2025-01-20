package doc

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
)

type Node struct {
	ID   string `cbor:"id" json:"id"`
	Type string `cbor:"type" json:"type"`
}

func NewNode(id string, t string) (Node, error) {

	node := Node{
		ID:   id,
		Type: t,
	}

	err := validateNode(node)
	if err != nil {
		return Node{}, fmt.Errorf("doc/NewNode: %w", err)
	}

	return Node{
		ID:   id,
		Type: t,
	}, nil
}

// Takes a libp2p PeerID and sets it as the PeerID of the document.
// This is the PeerID of the node to dial to communicate with the entity.
func (d *Document) SetP2PNode(peerid peer.ID) error {

	node, err := NewNode(peerid.String(), "p2p")
	if err != nil {
		return fmt.Errorf("doc/SetP2PNode: %w", err)
	}

	d.Node = node

	return nil
}

func (d *Document) GetP2PNode() (peer.ID, error) {
	if d.Node.Type != "p2p" {
		return "", ErrNodeTypeMissing
	}

	return peer.IDFromBytes([]byte(d.Node.ID))
}

func validateNode(node Node) error {
	if node.Type != "p2p" {
		return ErrInvalidNodeType
	}

	_, err := cid.Parse(node.ID)
	if err != nil {
		return fmt.Errorf("doc/validatePeerID: %w", err)
	}

	return nil
}
