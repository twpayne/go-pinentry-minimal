// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/twpayne/go-pinentry (interfaces: Process)

// Package pinentry is a generated GoMock package.
package pinentry

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProcess is a mock of Process interface.
type MockProcess struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMockRecorder
}

// MockProcessMockRecorder is the mock recorder for MockProcess.
type MockProcessMockRecorder struct {
	mock *MockProcess
}

// NewMockProcess creates a new mock instance.
func NewMockProcess(ctrl *gomock.Controller) *MockProcess {
	mock := &MockProcess{ctrl: ctrl}
	mock.recorder = &MockProcessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProcess) EXPECT() *MockProcessMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockProcess) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockProcessMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProcess)(nil).Close))
}

// ReadLine mocks base method.
func (m *MockProcess) ReadLine() ([]byte, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadLine")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadLine indicates an expected call of ReadLine.
func (mr *MockProcessMockRecorder) ReadLine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadLine", reflect.TypeOf((*MockProcess)(nil).ReadLine))
}

// Start mocks base method.
func (m *MockProcess) Start(arg0 string, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockProcessMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockProcess)(nil).Start), arg0, arg1)
}

// Write mocks base method.
func (m *MockProcess) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockProcessMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockProcess)(nil).Write), arg0)
}