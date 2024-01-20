package headers

// Parses a received message in form of a multibase encoded JSON string
func Parse(data []byte) (*Headers, error) {

	p, err := UnmarshalFromCBOR(data)
	if err != nil {
		return nil, err
	}

	err = p.Validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}
