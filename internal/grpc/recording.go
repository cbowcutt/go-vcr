package grpc

import (
	"github.com/cbowcutt/go-vcr/internal/utils"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
	"os"
)

type GrpcRecording struct {
	method      string
	rawRequest  interface{}
	rawResponse interface{}
	err         error
}

type GrpcYaml struct {
	Method   string
	Request  string
	Response string
}

func NewGrpcRecording(method string, request interface{}, response interface{}, err error) *GrpcRecording {
	return &GrpcRecording{
		method:      method,
		rawRequest:  request,
		rawResponse: response,
		err:         err,
	}
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
		Method:   g.method,
		Request:  reqJsonString,
		Response: respJsonString,
	}
	return yaml.Marshal(asYaml)
}

// filename will be <recording-path>/<g.Method>_n.yml,
// if filename exists, filename will be <recording-path>/<g.Method>_n.yml, starting with 0
func (g *GrpcRecording) WriteToFile() error {
	os.WriteFile(g.method + ".yml")
}

func FindMatchingRecordingByRequestData(method, rawRequest interface{}) *GrpcRecording {

}
