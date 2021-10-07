//nolint:deadcode,unused
package docs

import "github.com/josephzxy/timer_apiserver/internal/resource/v1/model"

// swagger:route POST /timers timer createTimerRequest
// Create a timer
// responses:
//		200: createTimerResponse
//		default: errResp

// swagger:route GET /timers/{name} getTimerRequest
// Get a timer
// responses:
// 		200: getTimerResponse
//		default: errResp

// swagger:route GET /timers getTimersRequest
// Get all timers
// responses:
//		200: getTimersResponse
// 		default: errResp

// swagger:route PUT /timers/{name} updateTimerRequest
// Update a timer
// responses:
// 		200: updateTimerResponse
//		default: errResp

// swagger:route DELETE /timers/{name} deleteTimerRequest
// Delete a timer
// responses:
// 		200: deleteTimerResponse
//		default: errResp

type dataObjField struct {
	Data model.Timer `json:"data"`
}

type dataColField struct {
	Data []model.Timer `json:"data"`
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

// swagger:parameters createTimerRequest
type createTimerRequest struct {
	// in:body
	Body model.TimerCore
}

// swagger:response createTimerResponse
type createTimerResponse struct {
	// in:body
	Body dataObjField
}

// swagger:parameters getTimerRequest
type getTimerRequest struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
}

// swagger:response getTimerResponse
type getTimerResponse struct {
	// in:body
	Body dataObjField
}

// swagger:parameters getTimersRequest
type getTimersRequest struct {
}

// swagger:response getTimersResponse
type getTimersResponse struct {
	Body dataColField
}

// swagger:parameters updateTimerRequest
type updateTimerRequest struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
	// in:body
	Body model.TimerCore
}

// swagger:response updateTimerResponse
type updateTimerResponse struct {
}

// swagger:parameters deleteTimerRequest
type deleteTimerRequest struct {
	// The name of the timer
	// in:path
	Name string `json:"name"`
}

// swagger:response deleteTimerResponse
type deleteTimerResponse struct {
}

// swagger:response errResp
type errResp struct {
	// in:body
	Body errField
}
