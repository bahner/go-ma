package internal

import (
	"fmt"

	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	mh "github.com/multiformats/go-multihash"

	"gitee.com/rupy/go_utils/convert"
	log "github.com/sirupsen/logrus"
)

type DagPutResponse struct {
	Cid struct {
		CidString string `json:"/"`
	} `json:"Cid"`
}

// IPFSDagAddCBOR takes a CBOR encoded byte array and adds it to IPFS.
func IPFSDagAddCBOR(data []byte) (cid.Cid, error) {

	n, err := cbor.WrapObject(data, mh.SHA2_256, -1)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("ipfs/dag/add: failed to wrap object: %w", err)
	}

	return n.Cid(), GetIPFSAPI().Dag().Add(GetContext(), n)
}

// Publish publishes a simple CBOR struct to IPFS and returns the CID.
// This is a kludge, as the kubo client/rpc is not working.
// The parameters are consistent with the correspconding IPFS API call.
// Ref. https://docs.ipfs.io/reference/http/api/#api-v0-dag-put
func IPFSDagPutWithOptions(data []byte, inputCodec string, storeCodec string, pin bool, hash string, allowBigBlock bool) (string, error) {

	// Create a buffer to write our multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form field writer for field 'file'
	part, err := writer.CreateFormField("file")
	if err != nil {
		return "", err
	}

	// Write CBOR data into the multipart form field
	_, err = part.Write(data)
	if err != nil {
		return "", err
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		return "", err
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
		return "", err
	}

	// Set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Debugf("doc/publish: response.Body: %s", string(respBody))

	// Unmarshal the JSON data into the struct
	var ipfsResp DagPutResponse
	err = json.Unmarshal(respBody, &ipfsResp)
	if err != nil {
		return "", err
	}

	return ipfsResp.Cid.CidString, nil
}

func IPFSDagPutCBOR(data []byte) (string, error) {
	return IPFSDagPutWithOptions(data, "dag-cbor", "dag-cbor", false, "sha2-256", false)
}

func IPFSDagPutCBORAndPin(data []byte) (string, error) {
	return IPFSDagPutWithOptions(data, "dag-cbor", "dag-cbor", true, "sha2-256", false)
}

func IPFSDagPutJSON(data []byte) (string, error) {
	return IPFSDagPutWithOptions(data, "dag-json", "dag-cbor", true, "sha2-256", false)
}
