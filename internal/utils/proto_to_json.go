package utils

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtoToJson(message proto.Message) (string, error) {
	jsonBytes, err := protojson.Marshal(message)
	if err != nil {
		return "nil", err
	}
	return string(jsonBytes), err
}