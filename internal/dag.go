package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"gitee.com/rupy/go_utils/convert"
	"github.com/ipfs/go-cid"
	log "github.com/sirupsen/logrus"
)

type DagPutResponse struct {
	Cid struct {
		CidString string `json:"/"`
	} `json:"Cid"`
}

// Publish publishes a simple CBOR struct to IPFS and returns the CID.
// This is a kludge, as the kubo client/rpc is not working.
// The parameters are consistent with the correspconding IPFS API call.
// Ref. https://docs.ipfs.io/reference/http/api/#api-v0-dag-put
func IPFSDagPutWithOptions(data []byte, inputCodec string, storeCodec string, pin bool, hash string, allowBigBlock bool) (cid.Cid, error) {

	// Create a buffer to write our multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form field writer for field 'file'
	part, err := writer.CreateFormField("file")
	if err != nil {
		return cid.Cid{}, err
	}

	// Write CBOR data into the multipart form field
	_, err = part.Write(data)
	if err != nil {
		return cid.Cid{}, err
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return cid.Cid{}, err
	}

	// Build the URL with query parameters
	params := url.Values{}
	params.Add("store-codec", storeCodec)
	params.Add("input-codec", inputCodec)
	params.Add("pin", convert.Bool2Str(pin))
	params.Add("hash", hash)
	params.Add("allow-big-block", convert.Bool2Str(allowBigBlock))

	apiUrl := GetIPFSAPIUrl() + "/dag/put"
	apiUrl += "?" + params.Encode()

	// Prepare the HTTP request
	req, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		return cid.Cid{}, err
	}

	// Set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cid.Cid{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return cid.Cid{}, err
	}
	log.Debugf("doc/publish: response.Body: %s", string(respBody))

	// Unmarshal the JSON data into the struct
	var ipfsResp DagPutResponse
	err = json.Unmarshal(respBody, &ipfsResp)
	if err != nil {
		return cid.Cid{}, err
	}
	log.Debugf("doc/publish: unmarshalled response to: %s", ipfsResp)

	c, err := cid.Decode(ipfsResp.Cid.CidString)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("doc/publish: failed to decode CID: %w", err)
	}
	log.Debugf("doc/publish: decoded CID: %s", c)

	return c, nil
}

// Publish publishes a simple CBOR struct to IPFS and returns the CID.
// The parameters are the data, a  flag whether to pin the data and a flag whether to allow big blocks.
func IPFSDagPutCBOR(data []byte, pin bool, bb bool) (cid.Cid, error) {
	return IPFSDagPutWithOptions(data, "dag-cbor", "dag-cbor", pin, "sha2-256", bb)
}
