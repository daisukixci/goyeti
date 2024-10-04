// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -source=client.go -package=mocks -destination=mocks/client.gen.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/yeti-platform/goyeti/models"
	gomock "go.uber.org/mock/gomock"
)

// MockClientI is a mock of ClientI interface.
type MockClientI struct {
	ctrl     *gomock.Controller
	recorder *MockClientIMockRecorder
}

// MockClientIMockRecorder is the mock recorder for MockClientI.
type MockClientIMockRecorder struct {
	mock *MockClientI
}

// NewMockClientI creates a new mock instance.
func NewMockClientI(ctrl *gomock.Controller) *MockClientI {
	mock := &MockClientI{ctrl: ctrl}
	mock.recorder = &MockClientIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientI) EXPECT() *MockClientIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClientI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientI)(nil).Close))
}

// Init mocks base method.
func (m *MockClientI) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockClientIMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockClientI)(nil).Init))
}

// ObservablesSearch mocks base method.
func (m *MockClientI) ObservablesSearch(search, searchType string) (models.ServerResponseObservables, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObservablesSearch", search, searchType)
	ret0, _ := ret[0].(models.ServerResponseObservables)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObservablesSearch indicates an expected call of ObservablesSearch.
func (mr *MockClientIMockRecorder) ObservablesSearch(search, searchType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObservablesSearch", reflect.TypeOf((*MockClientI)(nil).ObservablesSearch), search, searchType)
}

// Query mocks base method.
func (m *MockClientI) Query(endpoint, method, data string) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", endpoint, method, data)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockClientIMockRecorder) Query(endpoint, method, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockClientI)(nil).Query), endpoint, method, data)
}