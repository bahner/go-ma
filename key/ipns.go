package key

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"os"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/internal"
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
// This IPNSKey can have it's secret key exported as text
// for storage on a yellow sticker on your monitor
// where it belongs.
type IPNSKey struct {
	PrivKey        crypto.PrivKey
	PublicKey      crypto.PubKey
	EncodedPrivKey string
	Id             ipns.Name
	Name           string
}

// NewIPNSKey creates a new IPNSKey. It generates a new key pair
// and derives the IPNS name from the public key.
// This function does not require an IPFS node to be running.
func NewIPNSKey(name string) (IPNSKey, error) {

	if !internal.IsValidNanoID(name) {
		return IPNSKey{}, fmt.Errorf("key/new: invalid name: %v", name)
	}

	privKey, pubKey, err := generateEd25519KeyPair()
	if err != nil {
		return IPNSKey{}, err
	}

	encodedPrivKey, err := EncodeIPNSPrivKey(privKey)
	if err != nil {
		return IPNSKey{}, err
	}

	ipnsName, err := ipnsNameFromPublicKey(pubKey)
	if err != nil {
		return IPNSKey{}, err
	}

	return IPNSKey{
		PrivKey:        privKey,
		PublicKey:      pubKey,
		EncodedPrivKey: encodedPrivKey,
		Id:             ipnsName,
		Name:           name,
	}, nil
}

// We want to be able to create a key from an stored secret key.
// This function does not require an IPFS node to be running.
func NewIPNSKeyFromEncodedPrivKey(encodedPrivKey string, name string) (IPNSKey, error) {
	privKey, pubKey, err := keyPairFromEncodedPrivkey(encodedPrivKey)
	if err != nil {
		return IPNSKey{}, err
	}

	ipnsName, err := ipnsNameFromPublicKey(pubKey)
	if err != nil {
		return IPNSKey{}, err
	}

	return IPNSKey{
		PrivKey:        privKey,
		PublicKey:      pubKey,
		EncodedPrivKey: encodedPrivKey,
		Id:             ipnsName,
		Name:           name,
	}, nil
}

func PrintEncodedIPNSKeyAndExit() {

	encodedPrivKey, err := GenerateEncodedIPNSKey()
	if err != nil {
		log.Fatalf("Failed to generate encoded private key: %v", err)
	}

	fmt.Println(encodedPrivKey)

	os.Exit(0)

}

func GenerateEncodedIPNSKey() (string, error) {

	pk, _, err := generateEd25519KeyPair()
	if err != nil {
		return "", err
	}
	encodedPrivKey, err := EncodeIPNSPrivKey(pk)
	if err != nil {
		return "", err

	}

	return encodedPrivKey, nil
}

func DecodeIPNSPrivKey(privKey string) (crypto.PrivKey, error) {

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

func EncodeIPNSPrivKey(privKey crypto.PrivKey) (string, error) {

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

	privKey, err := DecodeIPNSPrivKey(encodedPrivKey)
	if err != nil {
		return nil, nil, err
	}

	return crypto.KeyPairFromStdKey(privKey)

}

// Import the key into IPFS under it's IPNS name.
// Doesn't try to be clever. If the same is already
// there - do nothing. If a key with the same name exist
// then fail. User will have to delete it manually or choose
// a different name.
func (i *IPNSKey) ExportToIPFS(name string) error {

	keyReader := bytes.NewReader([]byte(i.EncodedPrivKey))

	// Get the key from IPFS
	shell := internal.GetShell()
	err := shell.KeyImport(internal.GetContext(), name, keyReader)
	if err != nil {
		return fmt.Errorf("failed to import key: %v", err)
	}

	return nil

}
