package did

import "lukechampine.com/blake3"

// Creates a deterministic nonce for the Keyset encryption
// This is implementation specific and not a MA thing.
// The parameter is the size in bytes for the nonce.
func (d DID) Nonce(size int) []byte {

	idHash := blake3.Sum256(d.IPNSName().Cid().Bytes())
	return idHash[:size]
}
