package grpc

import (
	"context"

	 "google.golang.org/grpc"
)




func VcrUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	vcrMode := ctx.Value("VCR_MODE")
	switch vcrMode {
	case "off":
		resp, err = handler(ctx, req)
	case "record":
		resp, err = handler(ctx, req)
		recording := NewGrpcRecording(info.FullMethod, req, resp, err)
		recording.ToYaml()
	case "test":
	}
	return
}