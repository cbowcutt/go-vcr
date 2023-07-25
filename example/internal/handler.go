package internal

import (
	"context"
	"fmt"
	"github.com/cbowcutt/go-vcr/example/time_server/api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

type TimeHandler struct {
}

func (t *TimeHandler) GetTime(ctx context.Context, request *api.GetTimeRequest) (*api.GetTimeResponse, error) {
	currentTime := time.Now()

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
