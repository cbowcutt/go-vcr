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
	str := string(toYaml)
	assert.Nil(t, err)
	assert.Equal(t, "intro:\n    kind: grpc\n    data:\n        method: this.is.method\n        request: '{\"key\":\"Hello\",\"value\":\"Friend\",\"googleString\":\"someString\",\"someNumber\":42,\"googleNumber\":42}'\n        response: '{\"key\":\"Hello\",\"value\":\"Friend\",\"googleString\":\"someString\",\"someNumber\":42,\"googleNumber\":42}'\n", str)
}
