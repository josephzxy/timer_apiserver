package grpcserver

import (
	"context"
	"errors"
	"net"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
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

func (s *GRPCServer) startInsecureServing() error {
	host := s.cfg.InsecureServing.Addr()
	lis, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	zap.S().Infow("grpc server insecure serving is serving", "host", host)
	return s.insecureServer.Serve(lis)
}

func (s *GRPCServer) Start() error {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		failed := make(chan struct{}, 1)
		go func() {
			if err := s.startInsecureServing(); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
				zap.S().Errorf("error occured during grpc server insecure serving", "err", err)
				failed <- struct{}{}
			}
		}()

		select {
		case <-ctx.Done():
			return nil
		case <-failed:
			return errors.New("grpc server insecure serving failed")
		}
	})
	return eg.Wait()
}

func (s *GRPCServer) Stop() {
	zap.L().Info("gracefully shutting down grpc server")
	s.insecureServer.GracefulStop()
}
