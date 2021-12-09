// Code generated by MockGen. DO NOT EDIT.
// Source: ./repositories.go

// Package groupservice is a generated GoMock package.
package groupservice

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSmartContract is a mock of SmartContract interface.
type MockSmartContract struct {
	ctrl     *gomock.Controller
	recorder *MockSmartContractMockRecorder
}

// MockSmartContractMockRecorder is the mock recorder for MockSmartContract.
type MockSmartContractMockRecorder struct {
	mock *MockSmartContract
}

// NewMockSmartContract creates a new mock instance.
func NewMockSmartContract(ctrl *gomock.Controller) *MockSmartContract {
	mock := &MockSmartContract{ctrl: ctrl}
	mock.recorder = &MockSmartContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSmartContract) EXPECT() *MockSmartContractMockRecorder {
	return m.recorder
}

// GetGroup mocks base method.
func (m *MockSmartContract) GetGroup(arg0 context.Context, arg1 int64) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockSmartContractMockRecorder) GetGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockSmartContract)(nil).GetGroup), arg0, arg1)
}

// GetGroupIDs mocks base method.
func (m *MockSmartContract) GetGroupIDs(arg0 context.Context) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupIDs", arg0)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupIDs indicates an expected call of GetGroupIDs.
func (mr *MockSmartContractMockRecorder) GetGroupIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupIDs", reflect.TypeOf((*MockSmartContract)(nil).GetGroupIDs), arg0)
}
