package docs

import "github.com/josephzxy/timer_apiserver/internal/resource/v1/model"

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

//nolint:deadcode,unused
type dataObjField struct {
	Data model.Timer `json:"data"`
}

//nolint:deadcode,unused
type dataColField struct {
	Data []model.Timer `json:"data"`
}

//nolint:deadcode,unused
type errField struct {
	// in:body
	Err errInfo `json:"err"`
}

//nolint:deadcode,unused
type errInfo struct {
	// in:body
	// Application error code. E.g. 100001
	Code int `json:"code"`
	// Application error message. E.g. "Req validation failed"
	Msg string `json:"msg"`
}

//nolint:deadcode,unused
// swagger:parameters createTimerReq
type createTimerReq struct {
	// in:body
	Body model.TimerCore
}

//nolint:deadcode,unused
// swagger:response createTimerResp
type createTimerResp struct {
	// in:body
	Body dataObjField
}

//nolint:deadcode,unused
// swagger:parameters getTimerReq
type getTimerReq struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
}

//nolint:deadcode,unused
// swagger:response getTimerResp
type getTimerResp struct {
	// in:body
	Body dataObjField
}

//nolint:deadcode,unused
// swagger:parameters getTimersReq
type getTimersReq struct {
}

//nolint:deadcode,unused
// swagger:response getTimersResp
type getTimersResp struct {
	Body dataColField
}

//nolint:deadcode,unused
// swagger:parameters updateTimerReq
type updateTimerReq struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
	// in:body
	Body model.TimerCore
}

//nolint:deadcode,unused
// swagger:response updateTimerResp
type updateTimerResp struct {
}

//nolint:deadcode,unused
// swagger:parameters deleteTimerReq
type deleteTimerReq struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
}

//nolint:deadcode,unused
// swagger:response deleteTimerResp
type deleteTimerResp struct {
}

//nolint:deadcode,unused
// swagger:response errResp
type errResp struct {
	// in:body
	Body errField
}
