package doc

import (
	"fmt"

	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
)

const (
	topicNumFields     = 2
	DEFAULT_TOPIC_TYPE = "p2p/gossipsub"
)

type Topic struct {
	ID   string `cbor:"id" json:"id"`
	Type string `cbor:"type" json:"type"`
}

func NewTopic(id string, t string) (Topic, error) {

	topic := Topic{
		ID:   id,
		Type: t,
	}

	err := validateTopic(topic)
	if err != nil {
		return Topic{}, fmt.Errorf("doc/NewTopic: %w", err)
	}

	return Topic{
		ID:   id,
		Type: t,
	}, nil
}

// Takes a libp2p PeerID and sets it as the PeerID of the document.
// This is the PeerID of the node to dial to communicate with the entity.
func (d *Document) SetTopic(topicId string, topicType string) error {

	topic, err := NewTopic(topicId, topicType)
	if err != nil {
		return fmt.Errorf("doc/SetP2PTopic: %w", err)
	}

	d.Topic = topic

	return nil
}

func validateTopic(topic Topic) error {

	if topic.ID == "" {
		return ErrTopicIdMissing
	}

	if topic.Type == "" {
		return ErrTopicTypeMissing
	}

	if topic.Type != DEFAULT_TOPIC_TYPE {
		return ErrInvalidTopicType
	}

	_, err := cid.Parse(topic.ID)
	if err != nil {
		return fmt.Errorf("doc/validatePeerID: %w", err)
	}

	return nil
}

func buildTopicNode(topic Topic) (ipld.Node, error) {
	nb := basicnode.Prototype.Map.NewBuilder()
	ma, err := nb.BeginMap(topicNumFields)
	if err != nil {
		return nil, err
	}

	ma.AssembleKey().AssignString("id")
	ma.AssembleValue().AssignString(topic.ID)

	ma.AssembleKey().AssignString("type")
	ma.AssembleValue().AssignString(topic.Type)

	ma.Finish()

	return nb.Build(), nil
}
