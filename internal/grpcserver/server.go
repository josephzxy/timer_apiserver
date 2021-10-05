package grpcserver

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/josephzxy/timer_apiserver/api/grpc"
	"github.com/josephzxy/timer_apiserver/internal/grpcserver/service/v1/timer"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

type GRPCServer struct {
	cfg            Config
	serviceRouter  service.ServiceRouter
	insecureServer *grpc.Server
}

func New(cfg *Config, serviceRouter service.ServiceRouter) *GRPCServer {
	s := &GRPCServer{
		cfg:           *cfg,
		serviceRouter: serviceRouter,
	}
	s.init()
	return s
}

func (s *GRPCServer) init() {
	s.insecureServer = grpc.NewServer(s.cfg.InsecureServing.Options...)
	pb.RegisterTimerServer(s.insecureServer, timer.NewTimerServer(s.serviceRouter))
	reflection.Register(s.insecureServer)
}

func (s *GRPCServer) Start() error {
	host := s.cfg.InsecureServing.Addr()
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	zap.S().Infow("grpc server insecure serving is serving", "host", host)
	if err := s.insecureServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *GRPCServer) Stop() {
	s.insecureServer.GracefulStop()
}
