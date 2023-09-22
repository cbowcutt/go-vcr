package mixtape

import (
	"github.com/cbowcutt/go-vcr/internal/vcr/mixtape/track"
	"github.com/cbowcutt/go-vcr/testing/fixtures"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

func Test_Mixtape_ToYaml(t *testing.T) {
	protoRequest := &fixtures.TestRecord{
		Key:          "Hello",
		Value:        "Friend",
		GoogleString: &wrapperspb.StringValue{Value: "someString"},
		SomeNumber:   42,
		GoogleNumber: &wrapperspb.Int32Value{Value: 42},
	}
	protoResponse := &fixtures.TestRecord{
		Key:          "Hello",
		Value:        "Friend",
		GoogleString: &wrapperspb.StringValue{Value: "someString"},
		SomeNumber:   42,
		GoogleNumber: &wrapperspb.Int32Value{Value: 42},
	}

	testTrack := track.NewGrpcTrack("this.is.method")
	testTrack.SetGrpcRequestData(protoRequest)
	testTrack.SetGrpcResponseData(protoResponse)

	mixtape := Mixtape{
		Intro: testTrack,
	}
	toYaml, err := mixtape.ToYaml()
	assert.Nil(t, err)
	actualMixtape := Mixtape{}
	actualMixtape.FromYaml(toYaml)
	assert.Equal(t, "grpc", actualMixtape.Intro.Kind)
	grpcData := actualMixtape.Intro.Data.(track.GrpcTrackData)
	assert.Equal(t, "this.is.method", grpcData.Method)

	requestData := grpcData.Request.(map[string]interface{})
	assert.Equal(t, "Hello", requestData["key"])
	assert.Equal(t, "Friend", requestData["value"])
	assert.Equal(t, float64(42), requestData["googleNumber"])
	assert.Equal(t, float64(42), requestData["someNumber"])
	assert.Equal(t, "someString", requestData["googleString"])

	responseData := grpcData.Response.(map[string]interface{})
	assert.Equal(t, "Hello", responseData["key"])
	assert.Equal(t, "Friend", responseData["value"])
	assert.Equal(t, float64(42), responseData["googleNumber"])
	assert.Equal(t, float64(42), responseData["someNumber"])
	assert.Equal(t, "someString", responseData["googleString"])
}
