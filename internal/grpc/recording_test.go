package grpc

import (
	"github.com/cbowcutt/go-vcr/testing/fixtures"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gopkg.in/yaml.v3"
	"testing"
)

func Test_Recording_ToYaml(t *testing.T) {
	protoRequest := &fixtures.TestRecord{
		Key: "Hello",
		Value: "Friend",
		GoogleString: &wrapperspb.StringValue{Value: "someString"},
		SomeNumber: 42,
		GoogleNumber: &wrapperspb.Int32Value{Value: 42},
	}
	protoResponse := &fixtures.TestRecord{
		Key: "Hello",
		Value: "Friend",
		GoogleString: &wrapperspb.StringValue{Value: "someString"},
		SomeNumber: 42,
		GoogleNumber: &wrapperspb.Int32Value{Value: 42},
	}

	recording := GrpcRecording{
		method: "this.is.method",
		rawRequest: protoRequest,
		rawResponse: protoResponse,
	}
	yamlBytes, err := recording.ToYaml()
	assert.Nil(t, err)
	assert.NotNil(t, yamlBytes)
	var yamlMap map[interface{}]interface{}
	yaml.Unmarshal(yamlBytes, &yamlMap)
	assert.Equal(t, "this.is.method", yamlMap["method"])
	assert.Equal(t, `{"key":"Hello", "value":"Friend", "googleString":"someString", "someNumber":42, "googleNumber":42}`, yamlMap["request"])
	assert.Equal(t, `{"key":"Hello", "value":"Friend", "googleString":"someString", "someNumber":42, "googleNumber":42}`, yamlMap["response"])
}