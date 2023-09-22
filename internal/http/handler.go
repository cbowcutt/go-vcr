package http

import (
	"fmt"
	"github.com/cbowcutt/go-vcr/internal/vcr/mixtape"
	"github.com/cbowcutt/go-vcr/internal/vcr/mixtape/track"
	"github.com/cbowcutt/go-vcr/internal/vcr/mode"
	"golang.org/x/net/context"
	"net/http"
	"strings"
)

type Recorder interface {
	writeToFile(*http.Request)
}

type HttpRecorder struct {
	handler  http.Handler
	recorder Recorder
}

type VcrHttpInterceptor struct {
	core http.RoundTripper
}

func (i *VcrHttpInterceptor) RoundTrip(req *http.Request) (*http.Response, error) {
	vcrMode := mode.GetVcrMode()
	if *vcrMode == mode.RECORD {

		resp, err := i.core.RoundTrip(req)
		defer resp.Body.Close()
		err = i.recordClientRequest(req.Context(), req, resp)
		if err != nil {
			fmt.Println(err)
			return req.Response, err
		}
		return req.Response, err
	}
	if *vcrMode == mode.PLAYBACK {
		// TODO: make and call method for Playback
	}
	return req.Response, nil
}

func (i *VcrHttpInterceptor) recordClientRequest(ctx context.Context, req *http.Request, resp *http.Response) error {
	tape := ctx.Value("mixtape").(*mixtape.Mixtape)
	var responseBuffer []byte
	resp.Body.Read(responseBuffer)
	responseBodyString := string(responseBuffer)
	var reqBodyString string
	if strings.ToLower(req.Method) == "get" || strings.ToLower(req.Method) == "delete" {
		reqBodyString = ""
	}
	var reqBodyBytes []byte
	reader, err := req.GetBody()
	if err != nil {
		fmt.Println(err)
	}
	_, err = reader.Read(reqBodyBytes)
	reqBodyString = string(reqBodyBytes)
	if err != nil {
		fmt.Println(err)
	}
	httpTrack, err := track.NewHttpTrack(req.Method, req.URL.Path, req.URL.Query(), reqBodyString, responseBodyString, resp.StatusCode)
	if err != nil {
		fmt.Println(err)
		return err
	}
	tape.Tracks = append(tape.Tracks, httpTrack)
	return nil
}
