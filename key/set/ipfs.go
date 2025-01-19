package set

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/bahner/go-ma/api"
	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
)

func (k Keyset) SaveToIPFS() (string, error) {

	ctx := context.Background()

	packed, err := k.Pack()
	if err != nil {
		return "", fmt.Errorf("KeysetSaveToIPFS: %w", err)
	}

	shell := api.GetIPFSAPI()

	keysetFile := files.NewBytesFile([]byte(packed))

	immutablePath, err := shell.Unixfs().Add(ctx, keysetFile)
	if err != nil {
		return "", fmt.Errorf("KeysetSaveToIPFS: %w", err)
	}

	return immutablePath.String(), nil
}

func LoadFromIPFS(pathString string) (Keyset, error) {

	ctx := context.Background()

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

	return Unpack(string(keysetNodeBytes))
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
