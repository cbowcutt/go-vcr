package utils

import (
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
	assert.Equal(t, `{"key":"Hello","value":"Friend","googleString":"someString"}`, jsonString)
}
