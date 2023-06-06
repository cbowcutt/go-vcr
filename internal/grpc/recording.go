package grpc

import (
	"github.com/cbowcutt/go-vcr/internal/utils"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

type GrpcRecording struct {
	method	string
	rawRequest  interface{}
	rawResponse interface{}
}

type GrpcYaml struct {
	Method string
	Request string
	Response string
}

func (g *GrpcRecording) ToYaml() ([]byte, error) {

	reqJsonString, err := utils.ProtoToJson(g.rawRequest.(proto.Message))
	if err != nil {
		return nil, err
	}
	respJsonString, err := utils.ProtoToJson(g.rawResponse.(proto.Message))
	if err != nil {
		return nil, err
	}
	// return yaml struct with request and response
	asYaml := GrpcYaml{
		Method: g.method,
		Request: reqJsonString,
		Response: respJsonString,
	}
	return yaml.Marshal(asYaml)
}
