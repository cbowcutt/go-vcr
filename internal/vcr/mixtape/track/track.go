package track

import (
	"encoding/json"
	"errors"
	"github.com/cbowcutt/go-vcr/internal/utils"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
	"strings"
)

type Track struct {
	Kind string `yaml:"kind"`
	Data any    `yaml:"data"`
}

type GrpcTrackData struct {
	Method   string      `yaml:"method"`
	Request  interface{} `yaml:"request"`
	Response interface{} `yaml:"response"`
	Error    error       `yaml:"err"`
}

type HttpTrackData struct {
	Method     string              `yaml:"method"`
	Url        string              `yaml:"url"`
	Query      map[string][]string `yaml:"query"`
	Request    interface{}         `yaml:"request"`
	Response   interface{}         `yaml:"response"`
	StatusCode int                 `yaml:"status_code"`
}

func NewHttpTrack(method string, url string, query map[string][]string, request string, response string, statusCode int) (*Track, error) {
	return &Track{
		Kind: "http",
		Data: HttpTrackData{
			Method:     method,
			Url:        url,
			Request:    request,
			Query:      query,
			Response:   response,
			StatusCode: statusCode,
		},
	}, nil
}

func NewGrpcTrack(method string) *Track {
	return &Track{
		Kind: "grpc",
		Data: &GrpcTrackData{
			Method: method,
		},
	}
}

func (t *Track) SetGrpcRequestData(request interface{}) error {
	requestJson, err := utils.ProtoToJson(request.(proto.Message))
	if err != nil {
		return err
	}
	data := t.Data.(*GrpcTrackData)

	data.Request = requestJson
	t.Data = data
	return nil
}

func (t *Track) SetGrpcResponseData(response interface{}) error {
	responseJson, err := utils.ProtoToJson(response.(proto.Message))
	if err != nil {
		return err
	}
	data := t.Data.(*GrpcTrackData)
	data.Response = responseJson
	t.Data = data
	return nil
}

func (t *Track) SetError(err error) {
	t.Data.(*GrpcTrackData).Error = err
}

func (t *Track) Serialize() (string, error) {
	bytes, err := yaml.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func (t *Track) Decode() error {
	dataMap := t.Data.(map[string]interface{})
	if strings.ToLower(t.Kind) == "grpc" {
		grpcData := GrpcTrackData{}
		if dataMap["method"] != nil {
			grpcData.Method = dataMap["method"].(string)
		}
		if dataMap["request"] != nil {
			dataString := dataMap["request"].(string)
			request := make(map[string]interface{})
			err := json.Unmarshal([]byte(dataString), &request)
			if err != nil {
				return err
			}
			grpcData.Request = request
		}
		if dataMap["response"] != nil {
			dataString := dataMap["response"].(string)
			response := make(map[string]interface{})
			err := json.Unmarshal([]byte(dataString), &response)
			if err != nil {
				return err
			}
			grpcData.Response = response
		}
		if dataMap["error"] != nil {
			grpcData.Error = errors.New(dataMap["error"].(string))
		}
		t.Data = grpcData
	}
	if strings.ToLower(t.Kind) == "http" {
		httpData := HttpTrackData{}
		if dataMap["method"] != nil {
			httpData.Method = dataMap["method"].(string)
		}
		if dataMap["request"] != nil {
			httpData.Request = dataMap["request"].(map[string]interface{})
		}
		if dataMap["response"] != nil {
			httpData.Response = dataMap["response"].(map[string]interface{})
		}
		if dataMap["status_code"] != nil {
			httpData.StatusCode = dataMap["status_code"].(int)
		}
		if dataMap["url"] != nil {
			httpData.Url = dataMap["url"].(string)
		}
		if dataMap["query"] != nil {
			httpData.Query = dataMap["query"].(map[string][]string)
		}
		t.Data = httpData
	}
	return nil
}
