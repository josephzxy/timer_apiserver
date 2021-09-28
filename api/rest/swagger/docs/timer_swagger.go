package docs

import v1 "github.com/josephzxy/timer_apiserver/internal/resource/model/v1"

// swagger:route POST /timers timer createTimerReq
// Create a timer
// responses:
//		200: createTimerResp
//		default: errResp

// swagger:route GET /timers/{name} getTimerReq
// Get a timer
// responses:
// 		200: getTimerResp
//		default: errResp

// swagger:route GET /timers getTimersReq
// Get all timers
// responses:
//		200: getTimersResp
// 		default: errResp

// swagger:route PUT /timers/{name} updateTimerReq
// Update a timer
// responses:
// 		200: updateTimerResp
//		default: errResp

// swagger:route DELETE /timers/{name} deleteTimerReq
// Delete a timer
// responses:
// 		200: deleteTimerResp
//		default: errResp

type dataObjField struct {
	Data v1.Timer `json:"data"`
}

type dataColField struct {
	Data []v1.Timer `json:"data"`
}

type errField struct {
	// in:body
	Err errInfo `json:"err"`
}

type errInfo struct {
	// in:body
	// Application error code. E.g. 100001
	Code int `json:"code"`
	// Application error message. E.g. "Req validation failed"
	Msg string `json:"msg"`
}

// swagger:parameters createTimerReq
type createTimerReq struct {
	// in:body
	Body v1.TimerCore
}

// swagger:response createTimerResp
type createTimerResp struct {
	// in:body
	Body dataObjField
}

// swagger:parameters getTimerReq
type getTimerReq struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
}

// swagger:response getTimerResp
type getTimerResp struct {
	// in:body
	Body dataObjField
}

// swagger:parameters getTimersReq
type getTimersReq struct {
}

// swagger:response getTimersResp
type getTimersResp struct {
	Body dataColField
}

// swagger:parameters updateTimerReq
type updateTimerReq struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
	// in:body
	Body v1.TimerCore
}

// swagger:response updateTimerResp
type updateTimerResp struct {
	// in:body
	Body dataObjField
}

// swagger:parameters deleteTimerReq
type deleteTimerReq struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
}

// swagger:response deleteTimerResp
type deleteTimerResp struct {
	// in:body
	Body dataObjField
}

// swagger:response errResp
type errResp struct {
	// in:body
	Body errField
}
