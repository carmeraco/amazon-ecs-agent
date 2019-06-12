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
// Source: github.com/aws/amazon-ecs-agent/agent/app/oswrapper (interfaces: OS)

// Package mock_oswrapper is a generated GoMock package.
package mock_oswrapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOS is a mock of OS interface
type MockOS struct {
	ctrl     *gomock.Controller
	recorder *MockOSMockRecorder
}

// MockOSMockRecorder is the mock recorder for MockOS
type MockOSMockRecorder struct {
	mock *MockOS
}

// NewMockOS creates a new mock instance
func NewMockOS(ctrl *gomock.Controller) *MockOS {
	mock := &MockOS{ctrl: ctrl}
	mock.recorder = &MockOSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOS) EXPECT() *MockOSMockRecorder {
	return m.recorder
}

// Getpid mocks base method
func (m *MockOS) Getpid() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getpid")
	ret0, _ := ret[0].(int)
	return ret0
}

// Getpid indicates an expected call of Getpid
func (mr *MockOSMockRecorder) Getpid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getpid", reflect.TypeOf((*MockOS)(nil).Getpid))
}
