// Code generated by MockGen. DO NOT EDIT.
// Source: team.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	team "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	gomock "github.com/golang/mock/gomock"
)

// MockTeam is a mock of Team interface.
type MockTeam struct {
	ctrl     *gomock.Controller
	recorder *MockTeamMockRecorder
}

// MockTeamMockRecorder is the mock recorder for MockTeam.
type MockTeamMockRecorder struct {
	mock *MockTeam
}

// NewMockTeam creates a new mock instance.
func NewMockTeam(ctrl *gomock.Controller) *MockTeam {
	mock := &MockTeam{ctrl: ctrl}
	mock.recorder = &MockTeamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeam) EXPECT() *MockTeamMockRecorder {
	return m.recorder
}

// RegisterTeam mocks base method.
func (m *MockTeam) RegisterTeam(ctx context.Context, dto *team.TeamRequest) (*team.TeamResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTeam", ctx, dto)
	ret0, _ := ret[0].(*team.TeamResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTeam indicates an expected call of RegisterTeam.
func (mr *MockTeamMockRecorder) RegisterTeam(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTeam", reflect.TypeOf((*MockTeam)(nil).RegisterTeam), ctx, dto)
}