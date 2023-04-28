// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/LeandroAlcantara-1997/heroes-social-network/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateHero mocks base method.
func (m *MockRepository) CreateHero(ctx context.Context, hero *model.Hero) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHero", ctx, hero)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHero indicates an expected call of CreateHero.
func (mr *MockRepositoryMockRecorder) CreateHero(ctx, hero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHero", reflect.TypeOf((*MockRepository)(nil).CreateHero), ctx, hero)
}

// CreateTeam mocks base method.
func (m *MockRepository) CreateTeam(ctx context.Context, team *model.Team) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", ctx, team)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *MockRepositoryMockRecorder) CreateTeam(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockRepository)(nil).CreateTeam), ctx, team)
}

// DeleteHeroByID mocks base method.
func (m *MockRepository) DeleteHeroByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHeroByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHeroByID indicates an expected call of DeleteHeroByID.
func (mr *MockRepositoryMockRecorder) DeleteHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHeroByID", reflect.TypeOf((*MockRepository)(nil).DeleteHeroByID), ctx, id)
}

// GetHeroByID mocks base method.
func (m *MockRepository) GetHeroByID(ctx context.Context, id string) (*model.Hero, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeroByID", ctx, id)
	ret0, _ := ret[0].(*model.Hero)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeroByID indicates an expected call of GetHeroByID.
func (mr *MockRepositoryMockRecorder) GetHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeroByID", reflect.TypeOf((*MockRepository)(nil).GetHeroByID), ctx, id)
}

// UpdateHero mocks base method.
func (m *MockRepository) UpdateHero(ctx context.Context, hero *model.Hero) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHero", ctx, hero)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHero indicates an expected call of UpdateHero.
func (mr *MockRepositoryMockRecorder) UpdateHero(ctx, hero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHero", reflect.TypeOf((*MockRepository)(nil).UpdateHero), ctx, hero)
}

// MockHeroRepository is a mock of HeroRepository interface.
type MockHeroRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHeroRepositoryMockRecorder
}

// MockHeroRepositoryMockRecorder is the mock recorder for MockHeroRepository.
type MockHeroRepositoryMockRecorder struct {
	mock *MockHeroRepository
}

// NewMockHeroRepository creates a new mock instance.
func NewMockHeroRepository(ctrl *gomock.Controller) *MockHeroRepository {
	mock := &MockHeroRepository{ctrl: ctrl}
	mock.recorder = &MockHeroRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHeroRepository) EXPECT() *MockHeroRepositoryMockRecorder {
	return m.recorder
}

// CreateHero mocks base method.
func (m *MockHeroRepository) CreateHero(ctx context.Context, hero *model.Hero) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHero", ctx, hero)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHero indicates an expected call of CreateHero.
func (mr *MockHeroRepositoryMockRecorder) CreateHero(ctx, hero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHero", reflect.TypeOf((*MockHeroRepository)(nil).CreateHero), ctx, hero)
}

// DeleteHeroByID mocks base method.
func (m *MockHeroRepository) DeleteHeroByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHeroByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHeroByID indicates an expected call of DeleteHeroByID.
func (mr *MockHeroRepositoryMockRecorder) DeleteHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHeroByID", reflect.TypeOf((*MockHeroRepository)(nil).DeleteHeroByID), ctx, id)
}

// GetHeroByID mocks base method.
func (m *MockHeroRepository) GetHeroByID(ctx context.Context, id string) (*model.Hero, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeroByID", ctx, id)
	ret0, _ := ret[0].(*model.Hero)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeroByID indicates an expected call of GetHeroByID.
func (mr *MockHeroRepositoryMockRecorder) GetHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeroByID", reflect.TypeOf((*MockHeroRepository)(nil).GetHeroByID), ctx, id)
}

// UpdateHero mocks base method.
func (m *MockHeroRepository) UpdateHero(ctx context.Context, hero *model.Hero) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHero", ctx, hero)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHero indicates an expected call of UpdateHero.
func (mr *MockHeroRepositoryMockRecorder) UpdateHero(ctx, hero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHero", reflect.TypeOf((*MockHeroRepository)(nil).UpdateHero), ctx, hero)
}

// MockTeamRepository is a mock of TeamRepository interface.
type MockTeamRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTeamRepositoryMockRecorder
}

// MockTeamRepositoryMockRecorder is the mock recorder for MockTeamRepository.
type MockTeamRepositoryMockRecorder struct {
	mock *MockTeamRepository
}

// NewMockTeamRepository creates a new mock instance.
func NewMockTeamRepository(ctrl *gomock.Controller) *MockTeamRepository {
	mock := &MockTeamRepository{ctrl: ctrl}
	mock.recorder = &MockTeamRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamRepository) EXPECT() *MockTeamRepositoryMockRecorder {
	return m.recorder
}

// CreateTeam mocks base method.
func (m *MockTeamRepository) CreateTeam(ctx context.Context, team *model.Team) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", ctx, team)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *MockTeamRepositoryMockRecorder) CreateTeam(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*MockTeamRepository)(nil).CreateTeam), ctx, team)
}
