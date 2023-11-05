package message

import (
	"mime"

	log "github.com/sirupsen/logrus"
)

// For more short names, see
// https://www.iana.org/assignments/media-types/media-types.xhtml
var CommonMimeTypes = map[string]string{
	"car":         "application/vnd.ipld.car",
	"ipld_cbor":   "application/vnd.ipld.cbor",
	"ipld_json":   "application/vnd.ipld.json",
	"ipld":        "application/vnd.ipld.raw",
	"ipns-record": "application/vnd.ipfs.ipns-record",
	"json":        "application/json",
	"message":     MIME_TYPE,
	"text":        "text/plain",
}

// MimeTypes returns a list of the keys (shorthand names) from CommonMimeTypes.
func MimeTypeAliases() []string {
	keys := make([]string, 0, len(CommonMimeTypes))
	for k := range CommonMimeTypes {
		keys = append(keys, k)
	}
	return keys
}

// MimeTypeValues returns a list of the MIME type values from CommonMimeTypes.
func MimeTypes() []string {
	values := make([]string, 0, len(CommonMimeTypes))
	for _, v := range CommonMimeTypes {
		values = append(values, v)
	}
	return values
}

// MimeTypeTuples returns a list of key-value pairs (tuples) from CommonMimeTypes.
func MimeTypeTuples() [][2]string {
	tuples := make([][2]string, 0, len(CommonMimeTypes))
	for k, v := range CommonMimeTypes {
		tuple := [2]string{k, v}
		tuples = append(tuples, tuple)
	}
	return tuples
}

// MimeType returns the MIME type value for a given key.
func MimeType(mime_type string) string {
	return CommonMimeTypes[mime_type]
}

func IsValidMimeType(mimetype string) bool {
	_, _, err := mime.ParseMediaType(mimetype)
	if err != nil {
		log.Errorf("Invalid MIME type: %s", mimetype)
	}
	return err == nil
}
