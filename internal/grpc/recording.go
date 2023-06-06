package grpc

import (
	"github.com/cbowcutt/go-vcr/utils"
)

type GrpcRecording struct {
	rawRequest  interface{}
	RawResponse interface{}
}

func (g *GrpcRecording) ToYaml() error {
	reqJsonString, err := utils.ProtoToJson(g.rawRequest)
	if err != nil {
		return err
	}
	respJsonString, err := utils.ProtoToJson(g.rawResponse)
	if err != nil {
		return err
	}
	// return yaml struct with request and response
}
