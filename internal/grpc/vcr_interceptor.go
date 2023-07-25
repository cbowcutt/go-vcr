package grpc

import (
	"context"
	"github.com/cbowcutt/go-vcr/internal/vcr/mode"
	"google.golang.org/grpc"
)




func VcrUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	vcrMode := mode.GetVcrMode()
	switch *vcrMode {
	case mode.VCR_OFF:
		resp, err = handler(ctx, req)
	case mode.VCR_RECORD:
		resp, err = handler(ctx, req)
		recording := NewGrpcRecording(info.FullMethod, req, resp, err)
		recording.ToYaml()
	case mode.VCR_TEST:
		// find a recording that matches the request
	}
	return
}

func VcrClientUnaryInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	var err error
	vcrMode := mode.GetVcrMode()
	switch *vcrMode {
	case mode.VCR_OFF:
		err = invoker(ctx, method, req, reply, cc, opts...)
	case mode.VCR_RECORD:
		err = invoker(ctx, method, req, reply, cc, opts...)
		recording := NewGrpcRecording(method, req, reply, err)
		recording.ToYaml()
	case mode.VCR_TEST:
		// find recording
		err = invoker(ctx, method, req, reply, cc, opts...)
	}
	err = invoker(ctx, method, req, reply, cc, opts...)
	return err
}