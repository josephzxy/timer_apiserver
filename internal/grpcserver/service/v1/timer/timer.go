package timer

import (
	pb "github.com/josephzxy/timer_apiserver/api/grpc"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

type timerServer struct {
	pb.UnimplementedTimerServer

	serviceRouter service.ServiceRouter
}

func NewTimerServer(serviceRouter service.ServiceRouter) *timerServer {
	return &timerServer{serviceRouter: serviceRouter}
}
