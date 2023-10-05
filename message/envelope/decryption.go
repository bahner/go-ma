package envelope

import (
	"github.com/bahner/go-ma/message"
)

func (e *Envelope) Open(privKey []byte) (*message.Message, error) {
	return kyberEd25519Decrypt(e, privKey)
}
