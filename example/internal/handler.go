package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cbowcutt/go-vcr/example/time_server/api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"time"
)

type TimeHandler struct {
}

func (t *TimeHandler) GetTime(ctx context.Context, request *api.GetTimeRequest) (*api.GetTimeResponse, error) {
	target := make(map[string]interface{})
	//http.Handle()
	httpResponse, err := http.Get("http://worldtimeapi.org/api/timezone/utc")
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer httpResponse.Body.Close()
	decoder := json.NewDecoder(httpResponse.Body)
	err = decoder.Decode(&target)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	unixTime := target["unixtime"].(float64)

	currentTime := time.Unix(int64(unixTime), 0)

	resp := &api.GetTimeResponse{
		Time: timestamppb.New(currentTime),
	}
	locationName := request.GetTimezone().GetValue()
	if locationName != "" {
		loc, _ := time.LoadLocation(locationName)
		resp.LocalTime = &wrapperspb.StringValue{Value: fmt.Sprintf("%s\n", currentTime.In(loc))}
	}
	return resp, nil
}
