package main

import (
	"fmt"
	"github.com/cbowcutt/go-vcr/example/internal"
	"github.com/cbowcutt/go-vcr/example/time_server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	serverOptions := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 15 * time.Minute,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             30 * time.Second,
			PermitWithoutStream: true,
		}),
	}

	s := grpc.NewServer(serverOptions...)
	api.RegisterTimeServiceServer(s, &internal.TimeHandler{})
	shutdownGracePeriod := time.Minute * 60
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
		<-sigs
		//log.Infow("stopping res-inventory service", "shutdownGracePeriod", shutdownGracePeriod)

		// wait for some time to let kubernetes update service endpoints to avoid 503 errors
		if shutdownGracePeriod > 0 {
			<-time.After(shutdownGracePeriod)
		}

		timer := time.AfterFunc(shutdownGracePeriod, func() {
			fmt.Print("force stop gRPC server")
			s.Stop()
		})
		defer timer.Stop()

		s.GracefulStop()
	}()

	if err := s.Serve(lis); err != nil {
		fmt.Errorf("failed to serve: %w", err)
		return
	}

}
