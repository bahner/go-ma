package doc

import (
	"bytes"
	"fmt"

	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	ipldlegacy "github.com/ipfs/go-ipld-legacy"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	multihash "github.com/multiformats/go-multihash"
)

type documentNode struct {
	Node format.Node
	Cid  cid.Cid
}

func (d *Document) Node() (documentNode, error) {

	// Convert your struct to an IPLD node
	node, err := d.ipldStructure()
	if err != nil {
		panic(err)
	}

	var buf []byte
	buf, err = encodeIPLDNodeToDAGCBOR(node)
	if err != nil {
		return documentNode{}, fmt.Errorf("error encoding node to DAG-CBOR: %w", err)
	}
	// Create a CID for the block
	c, err := cid.V1Builder{Codec: cid.DagCBOR, MhType: multihash.SHA2_256}.Sum(buf)
	if err != nil {
		return documentNode{}, fmt.Errorf("error creating CID: %w", err)
	}

	// Create the block
	blk, err := blocks.NewBlockWithCid(buf, c)
	if err != nil {
		return documentNode{}, fmt.Errorf("error creating block: %w", err)
	}

	legacyNode := ipldlegacy.LegacyNode{Node: node, Block: blk}

	n := documentNode{Node: &legacyNode, Cid: c}

	return n, nil
}

func (d *Document) ipldStructure() (ipld.Node, error) {
	nb := basicnode.Prototype.Map.NewBuilder()
	ma, err := nb.BeginMap(7)
	if err != nil {
		return nil, err
	}

	// Context
	contextNode, err := buildStringListNode(d.Context)
	if err != nil {
		return nil, err
	}
	ma.AssembleKey().AssignString("context")
	ma.AssembleValue().AssignNode(contextNode)

	// ID
	ma.AssembleKey().AssignString("id")
	ma.AssembleValue().AssignString(d.ID)

	// Controllers
	controllerNode, err := buildStringListNode(d.Controller)
	if err != nil {
		return nil, err
	}
	ma.AssembleKey().AssignString("controller")
	ma.AssembleValue().AssignNode(controllerNode)

	// VerificationMethod
	ma.AssembleKey().AssignString("verificationMethod")
	verificationMethodsNode, err := buildVerificationMethodList(d.VerificationMethod)
	if err != nil {
		return nil, err
	}
	ma.AssembleValue().AssignNode(verificationMethodsNode)

	// AssertionMethod
	ma.AssembleKey().AssignString("assertionMethod")
	ma.AssembleValue().AssignString(d.AssertionMethod)

	// KeyAgreement
	ma.AssembleKey().AssignString("keyAgreement")
	ma.AssembleValue().AssignString(d.KeyAgreement)

	// Proof
	proofNode, err := buildProofNode(d.Proof)
	if err != nil {
		return nil, err
	}
	ma.AssembleKey().AssignString("proof")
	ma.AssembleValue().AssignNode(proofNode)

	ma.Finish()

	return nb.Build(), nil
}

func buildStringListNode(controllers []string) (ipld.Node, error) {
	nb := basicnode.Prototype.List.NewBuilder()
	la, err := nb.BeginList(-1)
	if err != nil {
		return nil, err
	}
	for _, controller := range controllers {
		la.AssembleValue().AssignString(controller)
	}
	la.Finish()

	return nb.Build(), nil
}

func encodeIPLDNodeToDAGCBOR(node ipld.Node) ([]byte, error) {
	var buf bytes.Buffer
	if err := dagcbor.Encode(node, &buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
