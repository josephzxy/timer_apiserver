// Package grpcserver implements a pluggable library for gRPC servers.
package grpcserver

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/pkg/errors"

	pb "github.com/josephzxy/timer_apiserver/api/grpc"
	"github.com/josephzxy/timer_apiserver/internal/grpcserver/service/v1/timer"
	"github.com/josephzxy/timer_apiserver/internal/pkg/util"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

// GRPCServer defines the interface of a gRPC server.
type GRPCServer interface {
	Start() error
	Stop()
}

type grpcServer struct {
	cfg            Config
	serviceRouter  service.ServiceRouter
	insecureServer *grpc.Server
}

// New returns the value of the implementation of interface GRPCServer
func New(cfg *Config, serviceRouter service.ServiceRouter) GRPCServer {
	s := &grpcServer{
		cfg:           *cfg,
		serviceRouter: serviceRouter,
	}
	s.init()
	return s
}

func (s *grpcServer) init() {
	s.insecureServer = grpc.NewServer(s.cfg.InsecureServing.Options...)
	pb.RegisterTimerServer(s.insecureServer, timer.NewTimerServer(s.serviceRouter))
	reflection.Register(s.insecureServer)
}

// startInsecureServing starts the insecure serving of grpcServer
func (s *grpcServer) startInsecureServing() error {
	addr := s.cfg.InsecureServing.Addr()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	zap.S().Infow("grpc server insecure serving starts", "addr", addr)
	if err := s.insecureServer.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		return err
	}
	return nil
}

// Start starts the gRPC server
func (s *grpcServer) Start() error {
	return util.BatchGoOrErr(
		s.startInsecureServing,
	)
}

// Stop gracefully stops the gRPC server
func (s *grpcServer) Stop() {
	zap.L().Info("grpc server starts shutting down gracefully")
	s.insecureServer.GracefulStop()
}
