package utils

import (
	"github.com/cbowcutt/go-vcr/testing/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)
func Test_ProtoToJson(t *testing.T) {
	jsonString, err := ProtoToJson(&fixtures.TestRecord{Key: "Hello", Value: "Friend"})
	assert.Nil(t, err)
	assert.Equal(t, `{"key":"Hello", "value":"Friend"}`, jsonString)
}
