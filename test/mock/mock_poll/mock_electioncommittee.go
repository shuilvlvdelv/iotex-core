// Code generated by MockGen. DO NOT EDIT.
// Source: ./action/protocol/poll/electioncommittee.go

// Package mock_poll is a generated GoMock package.
package mock_poll

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "github.com/iotexproject/iotex-election/types"
	reflect "reflect"
)

// MockElectionCommittee is a mock of ElectionCommittee interface
type MockElectionCommittee struct {
	ctrl     *gomock.Controller
	recorder *MockElectionCommitteeMockRecorder
}

// MockElectionCommitteeMockRecorder is the mock recorder for MockElectionCommittee
type MockElectionCommitteeMockRecorder struct {
	mock *MockElectionCommittee
}

// NewMockElectionCommittee creates a new mock instance
func NewMockElectionCommittee(ctrl *gomock.Controller) *MockElectionCommittee {
	mock := &MockElectionCommittee{ctrl: ctrl}
	mock.recorder = &MockElectionCommitteeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockElectionCommittee) EXPECT() *MockElectionCommitteeMockRecorder {
	return m.recorder
}

// ResultByHeight mocks base method
func (m *MockElectionCommittee) ResultByHeight(arg0 uint64) (*types.ElectionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResultByHeight", arg0)
	ret0, _ := ret[0].(*types.ElectionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResultByHeight indicates an expected call of ResultByHeight
func (mr *MockElectionCommitteeMockRecorder) ResultByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResultByHeight", reflect.TypeOf((*MockElectionCommittee)(nil).ResultByHeight), arg0)
}

// Start mocks base method
func (m *MockElectionCommittee) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockElectionCommitteeMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockElectionCommittee)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockElectionCommittee) Stop(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockElectionCommitteeMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockElectionCommittee)(nil).Stop), arg0)
}