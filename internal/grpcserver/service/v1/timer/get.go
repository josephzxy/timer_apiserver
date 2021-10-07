package timer

import (
	"context"
	"time"

	pb "github.com/josephzxy/timer_apiserver/api/grpc"
)

// GetAllPendingTimers gets all pending timers.
// A timer is "pending" if it is not deleted, and not triggerred yet.
func (s *timerServer) GetAllPendingTimers(context.Context, *pb.GetAllPendingTimersReq) (*pb.GetAllPendingTimersResp, error) {
	timers, err := s.serviceRouter.Timer().GetAllPending()
	if err != nil {
		return nil, err
	}

	data := make([]*pb.TimerInfo, 0, len(timers))
	for _, timer := range timers {
		data = append(data, &pb.TimerInfo{
			Name:      timer.Name,
			TriggerAt: timer.TriggerAt.Format(time.RFC3339),
		})
	}
	resp := &pb.GetAllPendingTimersResp{
		Items: data,
	}
	return resp, nil
}
