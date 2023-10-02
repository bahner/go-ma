package envelope

import (
	"encoding/json"
	"fmt"

	"github.com/bahner/go-ma/internal"
)

const IPLDMIMEType = "application/x-ma-envelope+ipld"

// IPLDLink represents an IPLD link with a CID.
type IPLDLink struct {
	CID string `json:"/"`
}

// IPLDEnvelope represents the encrypted message and the encrypted symmetric key in a JSON envelope.
type IPLDEnvelope struct {
	MIMEType     string   `json:"mime_type"`
	EncryptedKey string   `json:"encrypted_key"`
	IPLDLink     IPLDLink `json:"ipld_link"`
}

// New creates a new Envelope.
// IPLD envelope takes longer to generate, BUT the message is smaller.
// Also it's much faster, when sending the message sagain, because then
// the message is already on the network and possibly with the recipient.
// There are pros and cons to both approaches.
func NewIPLD(encodedMsg string, encodedEncryptedSymKey string) (*IPLDEnvelope, error) {

	// NB! This might take a while. We can't do it in a goroutine, because
	// we need the CID to return from this function.
	msgCID, err := internal.IPLDPutDag(encodedMsg)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("envelope: error publishing message DAG: %s\n", err))
	}
	return &IPLDEnvelope{
		MIMEType:     MIMEType,
		EncryptedKey: encodedEncryptedSymKey,
		IPLDLink:     IPLDLink{CID: msgCID},
	}, nil
}

func (e *IPLDEnvelope) MarshalToJSON() ([]byte, error) {
	return json.Marshal(e)
}

func UnmarshalIPLDFromJSON(data []byte) (*IPLDEnvelope, error) {

	e := &IPLDEnvelope{}

	err := json.Unmarshal(data, e)
	if err != nil {
		return nil, internal.LogError(fmt.Sprintf("envelope: error unmarshalling envelope: %s\n", err))
	}

	return e, nil
}

func (e *IPLDEnvelope) String() string {
	data, err := e.MarshalToJSON()
	if err != nil {
		return ""
	}
	return string(data)
}

func (e *IPLDEnvelope) GetCID() string {
	return e.IPLDLink.CID
}

func (e *IPLDEnvelope) GetMIMEType() string {
	return e.MIMEType
}

func (e *IPLDEnvelope) GetEncryptedKey() string {
	return e.EncryptedKey
}

func (e *IPLDEnvelope) GetIPLDLink() IPLDLink {
	return e.IPLDLink
}

func (e *IPLDEnvelope) GetEncryptedMsg() error {

	var ie IPLDLink
	shell := internal.GetShell()

	err := shell.DagGet(e.GetCID(), &ie)
	if err != nil {
		return internal.LogError(fmt.Sprintf("envelope: error opening Envelope message: %s\n", err))
	}

	return err
}
