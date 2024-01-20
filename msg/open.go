package msg

// While this is a very empty file, it's intended to be a starting point for
// validating an doing decision making on the message envelope.

func Open(data []byte) (*Envelope, error) {
	return UnmarshalEnvelopeFromCBOR(data)
}
