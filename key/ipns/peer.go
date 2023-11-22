package ipns

import (
	"fmt"

	"github.com/ipfs/boxo/ipns"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

func PublicKeyFromIdentifier(identifier string) (crypto.PubKey, error) {

	// Decode the PeerID from the Identifier which is the IPNS name
	pid, err := peer.Decode(identifier)
	if err != nil {
		return nil, fmt.Errorf("key/ipns: failed to decode peer ID from identifier: %w", err)
	}

	// Extract the public key from the peer ID
	pubKey, err := pid.ExtractPublicKey()
	if err != nil {
		return nil, fmt.Errorf("key/ipns: failed to extract public key from peer ID: %w", err)
	}

	// // Convert the crypto.PubKey into a byte slice
	// pubKey, err = cryptoPublicKey.Raw()
	// if err != nil {
	// 	return nil, fmt.Errorf("key/ipns: failed to convert crypto public key to byte slice: %w", err)
	// }

	return pubKey, nil
}

// func IdentifierFromPublicKey(pubKey ed25519.PublicKey) (ipns.Name, error) {
func IdentifierFromPublicKey(pubKey crypto.PubKey) (ipns.Name, error) {

	pid, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		return ipns.Name{}, fmt.Errorf("key/ipns: failed to generate peer ID from public key: %w", err)
	}
	ipnsName := ipns.NameFromPeer(pid)

	return ipnsName, nil
}
