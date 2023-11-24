package msg

// Parses a received message in form of a multibase encoded JSON string
func Parse(data string) (*Message, error) {

	p, err := Unpack(data)
	if err != nil {
		return nil, err
	}

	err = p.Validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}
