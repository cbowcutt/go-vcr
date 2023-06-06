func ProtoToJson(r interface{}) (string, error) {
	reqProto := g.RawRequest.(proto.Message)
	jsonBytes, err := protojson.Marshal(reqProto)
	if err != nil {
		return nil, err
	}
	return string(jsonBytes), err
}