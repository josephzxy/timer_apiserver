package err

import (
	"sync"

	"go.uber.org/zap"

	_ "github.com/josephzxy/timer_apiserver/internal/pkg/log"
)

// RESTAgent handles HTTP status and user-facing message
// for the response if an error of a certain AppErrCode occurs
// during RESTful API request.
type RESTAgent interface {
	// HTTP status for the response
	HTTPStatus() int
	// user-facing message for the response
	Msg() string
}

// SimpleRESTAgent is a simple internal implementation of RESTAgent interface
type SimpleRESTAgent struct {
	http int
	msg  string
}

func newSimpleRESTAgent(httpStatus int, msg string) *SimpleRESTAgent {
	return &SimpleRESTAgent{http: httpStatus, msg: msg}
}

func (s *SimpleRESTAgent) HTTPStatus() int { return s.http }
func (s *SimpleRESTAgent) Msg() string     { return s.msg }

var (
	restAgents map[AppErrCode]RESTAgent
	rwmtx      sync.RWMutex
)

func registerRESTAgent(code AppErrCode, httpStatus int, msg string) {
	allowedHttpStatus := []int{400, 404, 500}

	found := func() bool {
		for _, v := range allowedHttpStatus {
			if v == httpStatus {
				return true
			}
		}
		return false
	}()
	if !found {
		zap.S().Warnw("http status not allowed, will skip", "allowedHttpStatus", allowedHttpStatus, "got", httpStatus)
		return
	}

	rwmtx.Lock()
	defer rwmtx.Unlock()

	if _, ok := restAgents[code]; ok {
		zap.S().Warnw("error code already registered, will skip", "code", code)
	}
	restAgents[code] = newSimpleRESTAgent(httpStatus, msg)
}

func GetRESTAgent(code AppErrCode) RESTAgent {
	rwmtx.RLock()
	defer rwmtx.RUnlock()
	if a, ok := restAgents[code]; ok {
		return a
	} else {
		return restAgents[ErrUnknown]
	}
}

func init() {
	registerRESTAgent(ErrUnknown, 500, "Internal server error")
	registerRESTAgent(ErrValidation, 400, "Request validation failed")
	registerRESTAgent(ErrDatabase, 500, "Database error")
	registerRESTAgent(ErrTimerNotFound, 404, "Timer not found")
	registerRESTAgent(ErrTimerAlreadyExists, 400, "Timer already exists")
}
