package doc

import (
	"fmt"

	"github.com/bahner/go-ma/internal"
	api "github.com/ipfs/go-ipfs-api"
)

// Publish the document using a named key. The name is the name of the key in the IPFS node.

// The node is handled externally, so we don't need to worry about it here.

// If using a named key, the brave browser, this is the name of the key that matches your ID.

// The function also takes a string, which is the port the IPFS node is listening on.
// This is because Kubo listens on 5001, but Brave listens on 45005.

func (d *Document) Publish() (*api.PublishResponse, error) {

	// First we need to publish the document to IPFS and get the cid.

	// A document is not a DI :_) We need to parse the identifier out of the document.
	// in order to find the fragment, which *MUST* the keyname.

	data, err := d.String()
	if err != nil {
		return &api.PublishResponse{},
			internal.LogError(fmt.Sprintf("doc: failed to marshal document to JSON: %v", err))
	}

	// Lookup short name of the key to publish to in IPFS.

	name, err := internal.IPNSFindKeyName(d.ID)
	if err != nil {
		return &api.PublishResponse{},
			internal.LogError(fmt.Sprintf("doc: failed to find key name: %v", err))
	}

	cid, err := internal.IPFSPublishString(data)
	if err != nil {
		internal.LogError(fmt.Sprintf("doc: failed to publish document to IPFS: %v", err))
		return &api.PublishResponse{}, err
	}

	return internal.IPNSPublishCID(cid, name, true)

}
