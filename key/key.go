package key

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/bahner/go-ma"
	"github.com/multiformats/go-multibase"
	log "github.com/sirupsen/logrus"

	"github.com/ipfs/boxo/ipns"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

// This key struct is for libp2p keys. These are not meant
// for use in the DID document. As we just use IPNS keys
// from IPFS identity, these are just for convenience.
// One could evaluate adding an export of the IPNS key
// for import into IPFS, but that's deprecated in Kubo.
// This Key can have it's secret key exported as text
// for storage on a yellow sticker on your monitor
// where it belongs.
type Key struct {
	PrivKey        crypto.PrivKey
	PublicKey      crypto.PubKey
	EncodedPrivKey string
	IPNSName       ipns.Name
}

func New() (*Key, error) {
	privKey, pubKey, err := generateEd25519KeyPair()
	if err != nil {
		return nil, err
	}

	encodedPrivKey, err := EncodePrivKey(privKey)
	if err != nil {
		return nil, err
	}

	ipnsName, err := ipnsNameFromPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	return &Key{
		PrivKey:        privKey,
		PublicKey:      pubKey,
		EncodedPrivKey: encodedPrivKey,
		IPNSName:       ipnsName,
	}, nil
}
func NewFromEncodedPrivKey(encodedPrivKey string) (*Key, error) {
	privKey, pubKey, err := keyPairFromEncodedPrivkey(encodedPrivKey)
	if err != nil {
		return nil, err
	}

	ipnsName, err := ipnsNameFromPublicKey(pubKey)
	if err != nil {
		return nil, err
	}

	return &Key{
		PrivKey:        privKey,
		PublicKey:      pubKey,
		EncodedPrivKey: encodedPrivKey,
		IPNSName:       ipnsName,
	}, nil
}

func PrintEncodedKeyAndExit() {

	encodedPrivKey, err := GenerateEncodedKey()
	if err != nil {
		log.Fatalf("Failed to generate encoded private key: %v", err)
	}

	fmt.Println(encodedPrivKey)

	os.Exit(0)

}

func GenerateEncodedKey() (string, error) {

	pk, _, err := generateEd25519KeyPair()
	if err != nil {
		return "", err
	}
	encodedPrivKey, err := EncodePrivKey(pk)
	if err != nil {
		return "", err

	}

	return encodedPrivKey, nil
}

func DecodePrivKey(privKey string) (crypto.PrivKey, error) {

	// Decode the secret key
	_, decoded, err := multibase.Decode(privKey)
	if err != nil {
		log.Errorf("Failed to decode base58 secret key: %v", err)
		return nil, err
	}
	p, err := crypto.UnmarshalPrivateKey(decoded)
	if err != nil {
		log.Errorf("Failed to unmarshal private key: %v", err)
		return nil, err
	}

	return p, nil

}

func EncodePrivKey(privKey crypto.PrivKey) (string, error) {

	marshalledPrivKey, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		return "", fmt.Errorf("failed to marshal private key: %v", err)
	}

	encodedPrivKey, err := multibase.Encode(ma.MULTIBASE_ENCODING, marshalledPrivKey)
	if err != nil {
		return "", fmt.Errorf("failed to encode private key: %v", err)
	}

	return encodedPrivKey, nil
}

func generateEd25519KeyPair() (crypto.PrivKey, crypto.PubKey, error) {

	privKey, pubKey, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		log.Errorf("Failed to generate private key: %v", err)
		return nil, nil, err
	}

	return privKey, pubKey, nil
}

func ipnsNameFromPublicKey(pubKey crypto.PubKey) (ipns.Name, error) {

	pid, err := peer.IDFromPublicKey(pubKey)
	if err != nil {
		log.Errorf("Failed to generate peer ID from public key: %v", err)
		return ipns.Name{}, err
	}
	ipnsName := ipns.NameFromPeer(pid)

	return ipnsName, nil
}

func keyPairFromEncodedPrivkey(encodedPrivKey string) (crypto.PrivKey, crypto.PubKey, error) {

	privKey, err := DecodePrivKey(encodedPrivKey)
	if err != nil {
		return nil, nil, err
	}

	return crypto.KeyPairFromStdKey(privKey)

}
