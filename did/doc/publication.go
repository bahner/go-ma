package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did"
	"github.com/bahner/go-ma/internal"
	"github.com/ipfs/boxo/path"
	shell "github.com/ipfs/go-ipfs-api"
)

// Publish the document using a named key. The name is the name of the key in the IPFS node.

// The node is handled externally, so we don't need to worry about it here.

// If using a named key, the brave browser, this is the name of the key that matches your ID.

// The function also takes a string, which is the port the IPFS node is listening on.
// This is because Kubo listens on 5001, but Brave listens on 45005.

func (d *Document) Publish() error {

	api := internal.GetIPSAPI()

	// First we need to publish the document to IPFS and get the cid.

	// A document is not a DID :-) We need to parse the identifier out of the document.
	// in order to find the fragment, which *MUST* be the keyname.

	data, err := d.CBOR()
	if err != nil {
		return fmt.Errorf("doc/publish: failed to marshal document to JSON: %w", err)
	}

	// Publish the document to IPFS first. We need the CID to publish to IPNS.
	// So without that we ain't going nowhere.
	cid, err := internal.IPFSDagAddCBOR(data)
	if err != nil {
		return fmt.Errorf("doc: failed to publish document to IPFS: %w", err)
	}

	// Lookup short name of the identifier, ie. the fragment
	// The shortname is given to IPFS to lookup the actual key,
	// but that is transparent to use.
	// This gived us the possibility to change the key without
	// having to change the entity name within a given context.
	docdid, err := did.NewFromDID(d.ID)
	if err != nil {
		return fmt.Errorf("doc/publish: failed to parse DID: %w", err)
	}

	myPath, err := path.NewPath()
	if err != nil {
		return fmt.Errorf("doc/publish: failed to create path: %w", err)
	}

	return api.Name().Publish(internal.GetContext(), cid, shell.Name.Key(docdid.Fragment))

	return nil
}
