package http

import (
	"bytes"
	"encoding/json"
	"github.com/cbowcutt/go-vcr/internal/vcr/mixtape"
	track2 "github.com/cbowcutt/go-vcr/internal/vcr/mixtape/track"
	"github.com/cbowcutt/go-vcr/internal/vcr/mode"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockRoundTripper struct {
	response *http.Response
	error    error
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.response, m.error
}

func Test_Handler_RoundTrip_Record(t *testing.T) {
	vcrMode := mode.RECORD
	mode.SetVcrMode(&vcrMode)
	body := make(map[string]interface{})
	body["firstVal"] = "a"
	body["secondVal"] = 2.33
	bodyBytes, _ := json.Marshal(body)

	request, err := http.NewRequest("POST", "/search?a=1&b=2,3", bytes.NewReader(bodyBytes))
	if err != nil {
		assert.Fail(t, err.Error())
	}
	interceptor := &VcrHttpInterceptor{core: &MockRoundTripper{
		response: &http.Response{
			StatusCode: 200,
			Status:     "OK",
			Body:       io.NopCloser(strings.NewReader("{\"result\":\"OK\"}")),
		}}}

	tape := mixtape.Mixtape{}
	request = request.WithContext(context.WithValue(request.Context(), "mixtape", &tape))

	_, err = interceptor.RoundTrip(request)
	assert.Nil(t, err)
	assert.NotNil(t, tape.Tracks[0])
	track := tape.Tracks[0]
	trackData := track.Data.(track2.HttpTrackData)
	assert.Equal(t, "POST", trackData.Method)
	assert.Equal(t, "/search", trackData.Url)
	recordedBody := trackData.Request.(map[string]interface{})
	for k, v := range recordedBody {
		assert.Equal(t, body[k], v)
	}
	for k, v := range body {
		assert.Equal(t, recordedBody[k], v)
	}

}
