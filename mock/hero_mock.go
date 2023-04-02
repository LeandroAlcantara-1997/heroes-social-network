// Code generated by MockGen. DO NOT EDIT.
// Source: hero.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	input "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	gomock "github.com/golang/mock/gomock"
)

// MockHero is a mock of Hero interface.
type MockHero struct {
	ctrl     *gomock.Controller
	recorder *MockHeroMockRecorder
}

// MockHeroMockRecorder is the mock recorder for MockHero.
type MockHeroMockRecorder struct {
	mock *MockHero
}

// NewMockHero creates a new mock instance.
func NewMockHero(ctrl *gomock.Controller) *MockHero {
	mock := &MockHero{ctrl: ctrl}
	mock.recorder = &MockHeroMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHero) EXPECT() *MockHeroMockRecorder {
	return m.recorder
}

// DeleteHeroByID mocks base method.
func (m *MockHero) DeleteHeroByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHeroByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHeroByID indicates an expected call of DeleteHeroByID.
func (mr *MockHeroMockRecorder) DeleteHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHeroByID", reflect.TypeOf((*MockHero)(nil).DeleteHeroByID), ctx, id)
}

// GetHeroByID mocks base method.
func (m *MockHero) GetHeroByID(ctx context.Context, id string) (*input.HeroResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeroByID", ctx, id)
	ret0, _ := ret[0].(*input.HeroResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeroByID indicates an expected call of GetHeroByID.
func (mr *MockHeroMockRecorder) GetHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeroByID", reflect.TypeOf((*MockHero)(nil).GetHeroByID), ctx, id)
}

// RegisterHero mocks base method.
func (m *MockHero) RegisterHero(ctx context.Context, dto *input.HeroRequest) (*input.HeroResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterHero", ctx, dto)
	ret0, _ := ret[0].(*input.HeroResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterHero indicates an expected call of RegisterHero.
func (mr *MockHeroMockRecorder) RegisterHero(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHero", reflect.TypeOf((*MockHero)(nil).RegisterHero), ctx, dto)
}

// UpdateHero mocks base method.
func (m *MockHero) UpdateHero(ctx context.Context, id string, dto *input.HeroRequest) (*input.HeroResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHero", ctx, id, dto)
	ret0, _ := ret[0].(*input.HeroResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateHero indicates an expected call of UpdateHero.
func (mr *MockHeroMockRecorder) UpdateHero(ctx, id, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHero", reflect.TypeOf((*MockHero)(nil).UpdateHero), ctx, id, dto)
}
