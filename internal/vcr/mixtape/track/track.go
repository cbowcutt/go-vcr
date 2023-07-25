package track

import (
	"github.com/cbowcutt/go-vcr/internal/utils"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

type Track struct {
	Kind string      `yaml:"kind"`
	Data interface{} `yaml:"data"`
}

type GrpcTrackData struct {
	Method   string      `yaml:"method"`
	Request  interface{} `yaml:"request"`
	Response interface{} `yaml:"response"`
	err      error       `yaml:"err"`
}

type HttpTrackData struct {
	Method     string      `yaml:"method"`
	Url        string      `yaml:"url"`
	Request    interface{} `yaml:"request"`
	Response   interface{} `yaml:"response"`
	StatusCode int         `yaml:"status_code"`
}

func NewHttpTrack(method string, url string, request string, response string, statusCode int) (*Track, error) {
	return &Track{
		Kind: "http",
		Data: HttpTrackData{
			Method:     method,
			Url:        url,
			Request:    request,
			Response:   response,
			StatusCode: statusCode,
		},
	}, nil
}

func NewGrpcTrack(method string, request interface{}, response interface{}, err error) (*Track, error) {
	requestJson, err := utils.ProtoToJson(request.(proto.Message))
	if err != nil {
		return nil, err
	}
	responseJson, err := utils.ProtoToJson(response.(proto.Message))
	if err != nil {
		return nil, err
	}
	return &Track{
		Kind: "grpc",
		Data: GrpcTrackData{
			Method:   method,
			Request:  requestJson,
			Response: responseJson,
			err:      err,
		},
	}, nil
}

func (g *Track) Serialize() (string, error) {
	bytes, err := yaml.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

//// filename will be <recording-path>/<g.Method>_n.yml,
//// if filename exists, filename will be <recording-path>/<g.Method>_n.yml, starting with 0
//func (g *Track) WriteToFile() error {
//	os.WriteFile(g.method + ".yml")
//}
