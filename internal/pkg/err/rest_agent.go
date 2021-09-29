package err

import (
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"

	_ "github.com/josephzxy/timer_apiserver/internal/pkg/log"
	"github.com/josephzxy/timer_apiserver/internal/pkg/util"
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
	restAgents = make(map[AppErrCode]RESTAgent)
	rwmtx      sync.RWMutex
)

func registerRESTAgent(code AppErrCode, httpStatus int, msg string) error {
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
		msg := fmt.Sprintf("http status not allowed, will skip. should be one of %v, got %d", allowedHttpStatus, httpStatus)
		zap.L().Error(msg)
		return errors.New(msg)
	}

	rwmtx.Lock()
	defer rwmtx.Unlock()

	if _, ok := restAgents[code]; ok {
		msg := fmt.Sprintf("error code already registered, will skip. got %d", code)
		zap.L().Error(msg)
		return errors.New(msg)
	}
	restAgents[code] = newSimpleRESTAgent(httpStatus, msg)
	return nil
}

func GetRESTAgentByError(err error) RESTAgent {
	rwmtx.RLock()
	defer rwmtx.RUnlock()

	w, ok := err.(*WithCode)
	if !ok {
		return restAgents[ErrUnknown]
	}

	agent, ok := restAgents[w.Code()]
	if !ok {
		return restAgents[ErrUnknown]
	}
	return agent
}

func init() {
	util.PanicIfErr(registerRESTAgent(ErrUnknown, 500, "Internal server error"))
	util.PanicIfErr(registerRESTAgent(ErrValidation, 400, "Request validation failed"))
	util.PanicIfErr(registerRESTAgent(ErrDatabase, 500, "Database error"))
	util.PanicIfErr(registerRESTAgent(ErrTimerNotFound, 404, "Timer not found"))
	util.PanicIfErr(registerRESTAgent(ErrTimerAlreadyExists, 400, "Timer already exists"))
}
