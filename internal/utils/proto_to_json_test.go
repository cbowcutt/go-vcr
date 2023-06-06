package utils

import (
	"encoding/json"
	"github.com/cbowcutt/go-vcr/testing/fixtures"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)
func Test_ProtoToJson(t *testing.T) {
	protoRecord := &fixtures.TestRecord{
		Key: "Hello",
		Value: "Friend",
		GoogleString: &wrapperspb.StringValue{Value: "someString"}}
	jsonString, err := ProtoToJson(protoRecord)
	assert.Nil(t, err)
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonString), &jsonMap)
	assert.Nil(t, err)
	assert.Equal(t, "Hello", jsonMap["key"].(string))
	assert.Equal(t, "Friend", jsonMap["value"].(string))
	assert.Equal(t, "someString", jsonMap["googleString"].(string))
}
