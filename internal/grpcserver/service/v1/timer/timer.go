package timer

import (
	pb "github.com/josephzxy/timer_apiserver/api/grpc"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
)

type timerServer struct {
	pb.UnimplementedTimerServer

	serviceRouter service.ServiceRouter
}

// NewTimerServer returns a value of the implementation of the interface TimerServer
func NewTimerServer(serviceRouter service.ServiceRouter) pb.TimerServer {
	return &timerServer{serviceRouter: serviceRouter}
}
