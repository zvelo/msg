package msg

// Bytes returns the proto bytes of the Override
func (r Override) Bytes() ([]byte, error) {
	return r.Marshal()
}

// OverrideFromBytes returns an URLOverride from its protobuf encoded bytes
func OverrideFromBytes(data []byte) (*Override, error) {
	var o Override
	if err := o.Unmarshal(data); err != nil {
		return nil, err
	}
	return &o, nil
}
