package set

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"

	"github.com/bahner/go-ma/api"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/tyler-smith/go-bip39"
	"lukechampine.com/blake3"
)

func (k Keyset) SaveToIPFS(ctx context.Context, mnemonic string) (path.ImmutablePath, error) {

	packed, err := k.Pack()
	if err != nil {
		return path.ImmutablePath{}, fmt.Errorf("KeysetSaveToIPFS: %w", err)
	}

	encrypted, err := k.encryptKeysetWithMnemonic(mnemonic, []byte(packed))
	if err != nil {
		return path.ImmutablePath{}, fmt.Errorf("KeysetSaveToIPFS: %w", err)
	}

	shell := api.GetIPFSAPI()

	keysetFile := files.NewBytesFile([]byte(encrypted))

	immutablePath, err := shell.Unixfs().Add(ctx, keysetFile)
	if err != nil {
		return path.ImmutablePath{}, fmt.Errorf("KeysetSaveToIPFS: %w", err)
	}

	return immutablePath, nil
}

func LoadFromIPFS(ctx context.Context, pathString, mnemonic string, nonce []byte) (Keyset, error) {

	shell := api.GetIPFSAPI()

	keysetPath, err := path.NewPath(pathString)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetLoadFromIPFS: %w", err)
	}

	keysetNode, err := shell.Unixfs().Get(ctx, keysetPath)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetLoadFromIPFS: %w", err)
	}

	keysetNodeBytes, err := nodeBytes(keysetNode)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetLoadFromIPFS: %w", err)
	}

	// Decrypt the data using the mnemonic and the nonce
	decrypted, err := decryptKeysetWithMnemonicAndNonce(mnemonic, nonce, keysetNodeBytes)
	if err != nil {
		return Keyset{}, fmt.Errorf("KeysetLoadFromIPFS (decryption): %w", err)
	}

	// Unpack the decrypted data into a Keyset
	return Unpack(string(decrypted))
}

func nodeBytes(node files.Node) ([]byte, error) {
	fileNode, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("node is not a file")
	}

	// Read the file contents
	var buf bytes.Buffer
	_, err := io.Copy(&buf, fileNode)
	if err != nil {
		return nil, fmt.Errorf("failed to read node: %w", err)
	}

	return buf.Bytes(), nil
}

func (k Keyset) encryptKeysetWithMnemonic(mnemonic string, data []byte) ([]byte, error) {
	seed := bip39.NewSeed(mnemonic, "")
	secretKey := blake3.Sum256(seed)

	block, err := aes.NewCipher(secretKey[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := Nonce(k.DID)

	encrypted := aesGCM.Seal(nil, nonce, data, nil)
	return encrypted, nil
}

// Decrypts the keyset for a DID using a mnemonic and a nonce.
func decryptKeysetWithMnemonicAndNonce(mnemonic string, nonce []byte, encryptedData []byte) ([]byte, error) {
	seed := bip39.NewSeed(mnemonic, "")
	secretKey := blake3.Sum256(seed)

	block, err := aes.NewCipher(secretKey[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	decrypted, err := aesGCM.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}

	return decrypted, nil
}
