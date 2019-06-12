// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/wsclient/wsconn (interfaces: WebsocketConn)

// Package mock_wsconn is a generated GoMock package.
package mock_wsconn

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockWebsocketConn is a mock of WebsocketConn interface
type MockWebsocketConn struct {
	ctrl     *gomock.Controller
	recorder *MockWebsocketConnMockRecorder
}

// MockWebsocketConnMockRecorder is the mock recorder for MockWebsocketConn
type MockWebsocketConnMockRecorder struct {
	mock *MockWebsocketConn
}

// NewMockWebsocketConn creates a new mock instance
func NewMockWebsocketConn(ctrl *gomock.Controller) *MockWebsocketConn {
	mock := &MockWebsocketConn{ctrl: ctrl}
	mock.recorder = &MockWebsocketConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWebsocketConn) EXPECT() *MockWebsocketConnMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockWebsocketConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockWebsocketConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWebsocketConn)(nil).Close))
}

// ReadMessage mocks base method
func (m *MockWebsocketConn) ReadMessage() (int, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessage")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadMessage indicates an expected call of ReadMessage
func (mr *MockWebsocketConnMockRecorder) ReadMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockWebsocketConn)(nil).ReadMessage))
}

// SetReadDeadline mocks base method
func (m *MockWebsocketConn) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockWebsocketConnMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockWebsocketConn)(nil).SetReadDeadline), arg0)
}

// SetWriteDeadline mocks base method
func (m *MockWebsocketConn) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline
func (mr *MockWebsocketConnMockRecorder) SetWriteDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockWebsocketConn)(nil).SetWriteDeadline), arg0)
}

// WriteMessage mocks base method
func (m *MockWebsocketConn) WriteMessage(arg0 int, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMessage indicates an expected call of WriteMessage
func (mr *MockWebsocketConnMockRecorder) WriteMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMessage", reflect.TypeOf((*MockWebsocketConn)(nil).WriteMessage), arg0, arg1)
}
