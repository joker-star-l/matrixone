// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/matrixorigin/matrixone/pkg/common/morpc (interfaces: ClientSession,Stream)

// Package mock_morpc is a generated GoMock package.
package mock_morpc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	morpc "github.com/matrixorigin/matrixone/pkg/common/morpc"
)

// MockClientSession is a mock of ClientSession interface.
type MockClientSession struct {
	ctrl     *gomock.Controller
	recorder *MockClientSessionMockRecorder
}

// MockClientSessionMockRecorder is the mock recorder for MockClientSession.
type MockClientSessionMockRecorder struct {
	mock *MockClientSession
}

// NewMockClientSession creates a new mock instance.
func NewMockClientSession(ctrl *gomock.Controller) *MockClientSession {
	mock := &MockClientSession{ctrl: ctrl}
	mock.recorder = &MockClientSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientSession) EXPECT() *MockClientSessionMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClientSession) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientSessionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientSession)(nil).Close))
}

// CreateCache mocks base method.
func (m *MockClientSession) CreateCache(arg0 context.Context, arg1 uint64) (morpc.MessageCache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCache", arg0, arg1)
	ret0, _ := ret[0].(morpc.MessageCache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCache indicates an expected call of CreateCache.
func (mr *MockClientSessionMockRecorder) CreateCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCache", reflect.TypeOf((*MockClientSession)(nil).CreateCache), arg0, arg1)
}

// DeleteCache mocks base method.
func (m *MockClientSession) DeleteCache(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCache", arg0)
}

// DeleteCache indicates an expected call of DeleteCache.
func (mr *MockClientSessionMockRecorder) DeleteCache(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCache", reflect.TypeOf((*MockClientSession)(nil).DeleteCache), arg0)
}

// GetCache mocks base method.
func (m *MockClientSession) GetCache(arg0 uint64) (morpc.MessageCache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache", arg0)
	ret0, _ := ret[0].(morpc.MessageCache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCache indicates an expected call of GetCache.
func (mr *MockClientSessionMockRecorder) GetCache(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockClientSession)(nil).GetCache), arg0)
}

// RemoteAddress mocks base method.
func (m *MockClientSession) RemoteAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// RemoteAddress indicates an expected call of RemoteAddress.
func (mr *MockClientSessionMockRecorder) RemoteAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddress", reflect.TypeOf((*MockClientSession)(nil).RemoteAddress))
}

// Write mocks base method.
func (m *MockClientSession) Write(arg0 context.Context, arg1 morpc.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockClientSessionMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockClientSession)(nil).Write), arg0, arg1)
}

// MockStream is a mock of Stream interface.
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamMockRecorder
}

// MockStreamMockRecorder is the mock recorder for MockStream.
type MockStreamMockRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance.
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStream) EXPECT() *MockStreamMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStream) Close(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStreamMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStream)(nil).Close), arg0)
}

// ID mocks base method.
func (m *MockStream) ID() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockStreamMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockStream)(nil).ID))
}

// Receive mocks base method.
func (m *MockStream) Receive() (chan morpc.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(chan morpc.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive.
func (mr *MockStreamMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockStream)(nil).Receive))
}

// Send mocks base method.
func (m *MockStream) Send(arg0 context.Context, arg1 morpc.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockStreamMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockStream)(nil).Send), arg0, arg1)
}
