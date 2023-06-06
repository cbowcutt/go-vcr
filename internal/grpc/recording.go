package grpc

import(
	"encoding/json"
	"gopkg.in/yaml.v3"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

type GrpcRecording struct{
	rawRequest interface{}
	RawResponse interface{}
}

func (g *GrpcRecording) ToYaml() error {
	reqJsonString, err := convertToJson(g.rawRequest)
	if err != nil {
		return err
	}
	respJsonString, err := convertToJson(g.rawREsponse)
	if err != nil {
		return err
	}
	// return yaml struct with request and response
}

func convertToJson(r interface{}) (string, error) {
	reqProto := g.RawRequest.(proto.Message)
	jsonBytes, err := protojson.Marshal(reqProto)
	if err != nil {
		return nil, err
	}
	return string(jsonBytes), err
}