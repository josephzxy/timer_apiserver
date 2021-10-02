// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/josephzxy/work/lab/go/timer_apiserver/internal/resource/v1/service/timer.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/josephzxy/timer_apiserver/internal/resource/v1/model"
)

// MockTimerService is a mock of TimerService interface.
type MockTimerService struct {
	ctrl     *gomock.Controller
	recorder *MockTimerServiceMockRecorder
}

// MockTimerServiceMockRecorder is the mock recorder for MockTimerService.
type MockTimerServiceMockRecorder struct {
	mock *MockTimerService
}

// NewMockTimerService creates a new mock instance.
func NewMockTimerService(ctrl *gomock.Controller) *MockTimerService {
	mock := &MockTimerService{ctrl: ctrl}
	mock.recorder = &MockTimerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTimerService) EXPECT() *MockTimerServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTimerService) Create(arg0 *model.Timer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTimerServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTimerService)(nil).Create), arg0)
}

// DeleteByName mocks base method.
func (m *MockTimerService) DeleteByName(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByName", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByName indicates an expected call of DeleteByName.
func (mr *MockTimerServiceMockRecorder) DeleteByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByName", reflect.TypeOf((*MockTimerService)(nil).DeleteByName), name)
}

// GetByName mocks base method.
func (m *MockTimerService) GetByName(name string) (*model.Timer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", name)
	ret0, _ := ret[0].(*model.Timer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockTimerServiceMockRecorder) GetByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockTimerService)(nil).GetByName), name)
}

// UpdateByName mocks base method.
func (m *MockTimerService) UpdateByName(name string, want *model.TimerCore) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByName", name, want)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByName indicates an expected call of UpdateByName.
func (mr *MockTimerServiceMockRecorder) UpdateByName(name, want interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByName", reflect.TypeOf((*MockTimerService)(nil).UpdateByName), name, want)
}
