package set

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/bahner/go-ma/api"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/tyler-smith/go-bip39"
)

func (k Keyset) SaveToIPFS(ctx context.Context, mnemonic string) (path.ImmutablePath, error) {

	packed, err := k.Pack()
	if err != nil {
		return path.ImmutablePath{}, fmt.Errorf("KeysetSaveToIPFS: %w", err)
	}

	encrypted, err := encryptKeysetWithMnemonic(mnemonic, []byte(packed))
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

func LoadFromIPFS(ctx context.Context, pathString, mnemonic string) (Keyset, error) {

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

	// Decrypt the data using the mnemonic
	decrypted, err := decryptKeysetWithMnemonic(mnemonic, keysetNodeBytes)
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

func encryptKeysetWithMnemonic(mnemonic string, data []byte) ([]byte, error) {
	seed := bip39.NewSeed(mnemonic, "")
	secretKey := sha256.Sum256(seed[:32])

	block, err := aes.NewCipher(secretKey[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	// Generate a random nonce
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt the data
	encrypted := aesGCM.Seal(nil, nonce, data, nil)

	// Return nonce + encrypted data
	return append(nonce, encrypted...), nil
}

func decryptKeysetWithMnemonic(mnemonic string, encryptedData []byte) ([]byte, error) {
	// Derive the secret key from the BIP39 mnemonic
	seed := bip39.NewSeed(mnemonic, "")
	secretKey := sha256.Sum256(seed[:32])

	block, err := aes.NewCipher(secretKey[:])
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, fmt.Errorf("encrypted data is too short")
	}

	// Split the nonce and ciphertext
	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}

	return plaintext, nil
}
