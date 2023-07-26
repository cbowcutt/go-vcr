package grpc

import (
	"context"
	"github.com/cbowcutt/go-vcr/internal/vcr/mixtape"
	"github.com/cbowcutt/go-vcr/internal/vcr/mixtape/track"
	"github.com/cbowcutt/go-vcr/internal/vcr/mode"
	"google.golang.org/grpc"
)

func UnaryServerMixtapeInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		vcrMode := mode.GetVcrMode()
		switch *vcrMode {
		case mode.OFF:
			resp, err = handler(ctx, req)
		case mode.RECORD:
			tape := &mixtape.Mixtape{
				Intro: &track.Track{
					Kind: "grpc",
					Data: &track.GrpcTrackData{Method: info.FullMethod},
				},
			}
			newContext := context.WithValue(ctx, "mixtape", tape)

			tape.Intro.SetGrpcRequestData(req)
			resp, err := handler(newContext, req)
			if err != nil {
				tape.Intro.SetError(err)
			}
			tape.Intro.SetGrpcResponseData(resp)
			err = tape.ToFile()
			if err != nil {
				panic(err)
			}
			return resp, err
		case mode.PLAYBACK:
			// find a recording that matches the request
		}
		return resp, err
	}
}

//func VcrClientUnaryInterceptor(
//	ctx context.Context,
//	method string,
//	req interface{},
//	reply interface{},
//	cc *grpc.ClientConn,
//	invoker grpc.UnaryInvoker,
//	opts ...grpc.CallOption,
//) error {
//	var err error
//	vcrMode := mode.GetVcrMode()
//	switch *vcrMode {
//	case mode.OFF:
//		err = invoker(ctx, method, req, reply, cc, opts...)
//	case mode.RECORD:
//		tape := ctx.Value("mixtape")
//		if tape == nil {
//			return errors.New("could not find mixtape in given context")
//		}
//		err = invoker(ctx, method, req, reply, cc, opts...)
//		recording := NewGrpcRecording(method, req, reply, err)
//		recording.ToYaml()
//	case mode.PLAYBACK:
//		// find recording
//		err = invoker(ctx, method, req, reply, cc, opts...)
//	}
//	err = invoker(ctx, method, req, reply, cc, opts...)
//	return err
//}
