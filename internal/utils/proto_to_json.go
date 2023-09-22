package utils

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtoToJson(message proto.Message) (string, error) {
	jsonBytes, err := protojson.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), err
}

//func ProtoToJson(message proto.Message) (map[string]interface{}, error) {
//	result := make(map[string]interface{})
//	jsonBytes, err := protojson.Marshal(message)
//	if err != nil {
//		return result, err
//	}
//	err = json.Unmarshal(jsonBytes, &result)
//	if err != nil {
//		return result, err
//	}
//	return result, err
//}
