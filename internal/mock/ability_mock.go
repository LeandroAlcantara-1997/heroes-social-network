// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	dto "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockAbility is a mock of Ability interface.
type MockAbility struct {
	ctrl     *gomock.Controller
	recorder *MockAbilityMockRecorder
}

// MockAbilityMockRecorder is the mock recorder for MockAbility.
type MockAbilityMockRecorder struct {
	mock *MockAbility
}

// NewMockAbility creates a new mock instance.
func NewMockAbility(ctrl *gomock.Controller) *MockAbility {
	mock := &MockAbility{ctrl: ctrl}
	mock.recorder = &MockAbilityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAbility) EXPECT() *MockAbilityMockRecorder {
	return m.recorder
}

// CreateAbility mocks base method.
func (m *MockAbility) CreateAbility(ctx context.Context, req *dto.AbilityRequest) (*dto.AbilityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAbility", ctx, req)
	ret0, _ := ret[0].(*dto.AbilityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAbility indicates an expected call of CreateAbility.
func (mr *MockAbilityMockRecorder) CreateAbility(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAbility", reflect.TypeOf((*MockAbility)(nil).CreateAbility), ctx, req)
}

// DeleteAbility mocks base method.
func (m *MockAbility) DeleteAbility(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAbility", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAbility indicates an expected call of DeleteAbility.
func (mr *MockAbilityMockRecorder) DeleteAbility(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAbility", reflect.TypeOf((*MockAbility)(nil).DeleteAbility), ctx, id)
}

// GetAbilitiesByHeroID mocks base method.
func (m *MockAbility) GetAbilitiesByHeroID(ctx context.Context, id string) ([]dto.AbilityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAbilitiesByHeroID", ctx, id)
	ret0, _ := ret[0].([]dto.AbilityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAbilitiesByHeroID indicates an expected call of GetAbilitiesByHeroID.
func (mr *MockAbilityMockRecorder) GetAbilitiesByHeroID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAbilitiesByHeroID", reflect.TypeOf((*MockAbility)(nil).GetAbilitiesByHeroID), ctx, id)
}

// GetAbilityByID mocks base method.
func (m *MockAbility) GetAbilityByID(ctx context.Context, id string) (*dto.AbilityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAbilityByID", ctx, id)
	ret0, _ := ret[0].(*dto.AbilityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAbilityByID indicates an expected call of GetAbilityByID.
func (mr *MockAbilityMockRecorder) GetAbilityByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAbilityByID", reflect.TypeOf((*MockAbility)(nil).GetAbilityByID), ctx, id)
}
