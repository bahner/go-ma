package message

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

func (m *Message) Sign(privKey crypto.PrivKey) error {

	data_to_sign, err := m.PayloadPack()
	if err != nil {
		return err
	}

	bytes_to_sign := []byte(data_to_sign)

	sig, err := privKey.Sign(bytes_to_sign)
	if err != nil {
		return fmt.Errorf("failed to sign Message: %v", err)
	}

	encoded_sig, err := MessageEncoder(sig)
	if err != nil {
		return fmt.Errorf("failed to encode signature: %v", err)
	}

	m.Signature = encoded_sig

	return nil
}

// Verify verifies the Message's signature
func (m *Message) Verify() (bool, error) {

	did, err := did.Parse(m.From)
	if err != nil {
		return false, err
	}

	publicKey, err := did.PublicKey()
	if err != nil {
		return false, err
	}

	payload, err := m.PayloadPack()
	if err != nil {
		return false, err
	}

	return publicKey.Verify([]byte(payload), []byte(m.Signature))
}

func PublicKey(d *did.DID) (crypto.PubKey, error) {
	pid, err := peer.Decode(d.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Peer: %v", err)
	}

	return pid.ExtractPublicKey()

}
