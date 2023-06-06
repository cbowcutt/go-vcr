package grpc

import (
	"context"

	"google.golang.org/grpc"
)

func unaryClientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) (err error) {
	// get mode from Context
	vcrMode := ctx.Value("VCR_MODE")
	switch vcrMode {
	case "off":
		//err = invoker(ctx, method, req, reply, cc, opts)
	case "record":

	case "test":
	}
	return
}
