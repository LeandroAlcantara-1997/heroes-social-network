// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/model"
	model0 "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
	model1 "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
	model2 "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	model3 "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/model"
	gomock "github.com/golang/mock/gomock"
)

// RepositoryMock is a mock of Repository interface.
type RepositoryMock struct {
	ctrl     *gomock.Controller
	recorder *RepositoryMockMockRecorder
}

// RepositoryMockMockRecorder is the mock recorder for RepositoryMock.
type RepositoryMockMockRecorder struct {
	mock *RepositoryMock
}

// NewRepositoryMock creates a new mock instance.
func NewRepositoryMock(ctrl *gomock.Controller) *RepositoryMock {
	mock := &RepositoryMock{ctrl: ctrl}
	mock.recorder = &RepositoryMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *RepositoryMock) EXPECT() *RepositoryMockMockRecorder {
	return m.recorder
}

// AddAbilityToHero mocks base method.
func (m *RepositoryMock) AddAbilityToHero(ctx context.Context, abilityID, heroID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAbilityToHero", ctx, abilityID, heroID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAbilityToHero indicates an expected call of AddAbilityToHero.
func (mr *RepositoryMockMockRecorder) AddAbilityToHero(ctx, abilityID, heroID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAbilityToHero", reflect.TypeOf((*RepositoryMock)(nil).AddAbilityToHero), ctx, abilityID, heroID)
}

// CreateGame mocks base method.
func (m *RepositoryMock) CreateGame(ctx context.Context, game *model1.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGame", ctx, game)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGame indicates an expected call of CreateGame.
func (mr *RepositoryMockMockRecorder) CreateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGame", reflect.TypeOf((*RepositoryMock)(nil).CreateGame), ctx, game)
}

// CreateHero mocks base method.
func (m *RepositoryMock) CreateHero(ctx context.Context, hero *model2.Hero) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateHero", ctx, hero)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateHero indicates an expected call of CreateHero.
func (mr *RepositoryMockMockRecorder) CreateHero(ctx, hero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateHero", reflect.TypeOf((*RepositoryMock)(nil).CreateHero), ctx, hero)
}

// CreateTeam mocks base method.
func (m *RepositoryMock) CreateTeam(ctx context.Context, team *model3.Team) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeam", ctx, team)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTeam indicates an expected call of CreateTeam.
func (mr *RepositoryMockMockRecorder) CreateTeam(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeam", reflect.TypeOf((*RepositoryMock)(nil).CreateTeam), ctx, team)
}

// DeleteGameByID mocks base method.
func (m *RepositoryMock) DeleteGameByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGameByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGameByID indicates an expected call of DeleteGameByID.
func (mr *RepositoryMockMockRecorder) DeleteGameByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGameByID", reflect.TypeOf((*RepositoryMock)(nil).DeleteGameByID), ctx, id)
}

// DeleteHeroByID mocks base method.
func (m *RepositoryMock) DeleteHeroByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHeroByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHeroByID indicates an expected call of DeleteHeroByID.
func (mr *RepositoryMockMockRecorder) DeleteHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHeroByID", reflect.TypeOf((*RepositoryMock)(nil).DeleteHeroByID), ctx, id)
}

// DeleteTeamByID mocks base method.
func (m *RepositoryMock) DeleteTeamByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeamByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeamByID indicates an expected call of DeleteTeamByID.
func (mr *RepositoryMockMockRecorder) DeleteTeamByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeamByID", reflect.TypeOf((*RepositoryMock)(nil).DeleteTeamByID), ctx, id)
}

// GetGameByID mocks base method.
func (m *RepositoryMock) GetGameByID(ctx context.Context, id string) (*model1.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameByID", ctx, id)
	ret0, _ := ret[0].(*model1.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameByID indicates an expected call of GetGameByID.
func (mr *RepositoryMockMockRecorder) GetGameByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameByID", reflect.TypeOf((*RepositoryMock)(nil).GetGameByID), ctx, id)
}

// GetHeroByID mocks base method.
func (m *RepositoryMock) GetHeroByID(ctx context.Context, id string) (*model2.Hero, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeroByID", ctx, id)
	ret0, _ := ret[0].(*model2.Hero)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeroByID indicates an expected call of GetHeroByID.
func (mr *RepositoryMockMockRecorder) GetHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeroByID", reflect.TypeOf((*RepositoryMock)(nil).GetHeroByID), ctx, id)
}

// GetTeamByID mocks base method.
func (m *RepositoryMock) GetTeamByID(ctx context.Context, id string) (*model3.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamByID", ctx, id)
	ret0, _ := ret[0].(*model3.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamByID indicates an expected call of GetTeamByID.
func (mr *RepositoryMockMockRecorder) GetTeamByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamByID", reflect.TypeOf((*RepositoryMock)(nil).GetTeamByID), ctx, id)
}

// GetTeamByName mocks base method.
func (m *RepositoryMock) GetTeamByName(ctx context.Context, name string) (*model3.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamByName", ctx, name)
	ret0, _ := ret[0].(*model3.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamByName indicates an expected call of GetTeamByName.
func (mr *RepositoryMockMockRecorder) GetTeamByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamByName", reflect.TypeOf((*RepositoryMock)(nil).GetTeamByName), ctx, name)
}

// UpdateGame mocks base method.
func (m *RepositoryMock) UpdateGame(ctx context.Context, game *model1.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGame", ctx, game)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGame indicates an expected call of UpdateGame.
func (mr *RepositoryMockMockRecorder) UpdateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGame", reflect.TypeOf((*RepositoryMock)(nil).UpdateGame), ctx, game)
}

// UpdateHero mocks base method.
func (m *RepositoryMock) UpdateHero(ctx context.Context, hero *model2.Hero) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHero", ctx, hero)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHero indicates an expected call of UpdateHero.
func (mr *RepositoryMockMockRecorder) UpdateHero(ctx, hero interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHero", reflect.TypeOf((*RepositoryMock)(nil).UpdateHero), ctx, hero)
}

// UpdateTeam mocks base method.
func (m *RepositoryMock) UpdateTeam(ctx context.Context, team *model3.Team) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeam", ctx, team)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTeam indicates an expected call of UpdateTeam.
func (mr *RepositoryMockMockRecorder) UpdateTeam(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeam", reflect.TypeOf((*RepositoryMock)(nil).UpdateTeam), ctx, team)
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

// AddAbilityToHero mocks base method.
func (m *MockHeroRepository) AddAbilityToHero(ctx context.Context, abilityID, heroID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAbilityToHero", ctx, abilityID, heroID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAbilityToHero indicates an expected call of AddAbilityToHero.
func (mr *MockHeroRepositoryMockRecorder) AddAbilityToHero(ctx, abilityID, heroID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAbilityToHero", reflect.TypeOf((*MockHeroRepository)(nil).AddAbilityToHero), ctx, abilityID, heroID)
}

// CreateHero mocks base method.
func (m *MockHeroRepository) CreateHero(ctx context.Context, hero *model2.Hero) error {
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
func (m *MockHeroRepository) GetHeroByID(ctx context.Context, id string) (*model2.Hero, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeroByID", ctx, id)
	ret0, _ := ret[0].(*model2.Hero)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeroByID indicates an expected call of GetHeroByID.
func (mr *MockHeroRepositoryMockRecorder) GetHeroByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeroByID", reflect.TypeOf((*MockHeroRepository)(nil).GetHeroByID), ctx, id)
}

// UpdateHero mocks base method.
func (m *MockHeroRepository) UpdateHero(ctx context.Context, hero *model2.Hero) error {
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
func (m *MockTeamRepository) CreateTeam(ctx context.Context, team *model3.Team) error {
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

// DeleteTeamByID mocks base method.
func (m *MockTeamRepository) DeleteTeamByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeamByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeamByID indicates an expected call of DeleteTeamByID.
func (mr *MockTeamRepositoryMockRecorder) DeleteTeamByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeamByID", reflect.TypeOf((*MockTeamRepository)(nil).DeleteTeamByID), ctx, id)
}

// GetTeamByID mocks base method.
func (m *MockTeamRepository) GetTeamByID(ctx context.Context, id string) (*model3.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamByID", ctx, id)
	ret0, _ := ret[0].(*model3.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamByID indicates an expected call of GetTeamByID.
func (mr *MockTeamRepositoryMockRecorder) GetTeamByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamByID", reflect.TypeOf((*MockTeamRepository)(nil).GetTeamByID), ctx, id)
}

// GetTeamByName mocks base method.
func (m *MockTeamRepository) GetTeamByName(ctx context.Context, name string) (*model3.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamByName", ctx, name)
	ret0, _ := ret[0].(*model3.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamByName indicates an expected call of GetTeamByName.
func (mr *MockTeamRepositoryMockRecorder) GetTeamByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamByName", reflect.TypeOf((*MockTeamRepository)(nil).GetTeamByName), ctx, name)
}

// UpdateTeam mocks base method.
func (m *MockTeamRepository) UpdateTeam(ctx context.Context, team *model3.Team) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeam", ctx, team)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTeam indicates an expected call of UpdateTeam.
func (mr *MockTeamRepositoryMockRecorder) UpdateTeam(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeam", reflect.TypeOf((*MockTeamRepository)(nil).UpdateTeam), ctx, team)
}

// MockGameRepository is a mock of GameRepository interface.
type MockGameRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGameRepositoryMockRecorder
}

// MockGameRepositoryMockRecorder is the mock recorder for MockGameRepository.
type MockGameRepositoryMockRecorder struct {
	mock *MockGameRepository
}

// NewMockGameRepository creates a new mock instance.
func NewMockGameRepository(ctrl *gomock.Controller) *MockGameRepository {
	mock := &MockGameRepository{ctrl: ctrl}
	mock.recorder = &MockGameRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGameRepository) EXPECT() *MockGameRepositoryMockRecorder {
	return m.recorder
}

// CreateGame mocks base method.
func (m *MockGameRepository) CreateGame(ctx context.Context, game *model1.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGame", ctx, game)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateGame indicates an expected call of CreateGame.
func (mr *MockGameRepositoryMockRecorder) CreateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGame", reflect.TypeOf((*MockGameRepository)(nil).CreateGame), ctx, game)
}

// DeleteGameByID mocks base method.
func (m *MockGameRepository) DeleteGameByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGameByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGameByID indicates an expected call of DeleteGameByID.
func (mr *MockGameRepositoryMockRecorder) DeleteGameByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGameByID", reflect.TypeOf((*MockGameRepository)(nil).DeleteGameByID), ctx, id)
}

// GetGameByID mocks base method.
func (m *MockGameRepository) GetGameByID(ctx context.Context, id string) (*model1.Game, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGameByID", ctx, id)
	ret0, _ := ret[0].(*model1.Game)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGameByID indicates an expected call of GetGameByID.
func (mr *MockGameRepositoryMockRecorder) GetGameByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGameByID", reflect.TypeOf((*MockGameRepository)(nil).GetGameByID), ctx, id)
}

// UpdateGame mocks base method.
func (m *MockGameRepository) UpdateGame(ctx context.Context, game *model1.Game) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGame", ctx, game)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGame indicates an expected call of UpdateGame.
func (mr *MockGameRepositoryMockRecorder) UpdateGame(ctx, game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGame", reflect.TypeOf((*MockGameRepository)(nil).UpdateGame), ctx, game)
}

// MockConsoleRepository is a mock of ConsoleRepository interface.
type MockConsoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockConsoleRepositoryMockRecorder
}

// MockConsoleRepositoryMockRecorder is the mock recorder for MockConsoleRepository.
type MockConsoleRepositoryMockRecorder struct {
	mock *MockConsoleRepository
}

// NewMockConsoleRepository creates a new mock instance.
func NewMockConsoleRepository(ctrl *gomock.Controller) *MockConsoleRepository {
	mock := &MockConsoleRepository{ctrl: ctrl}
	mock.recorder = &MockConsoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsoleRepository) EXPECT() *MockConsoleRepositoryMockRecorder {
	return m.recorder
}

// CreateConsoles mocks base method.
func (m *MockConsoleRepository) CreateConsoles(ctx context.Context, consoles []model0.Console) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConsoles", ctx, consoles)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateConsoles indicates an expected call of CreateConsoles.
func (mr *MockConsoleRepositoryMockRecorder) CreateConsoles(ctx, consoles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConsoles", reflect.TypeOf((*MockConsoleRepository)(nil).CreateConsoles), ctx, consoles)
}

// GetConsoles mocks base method.
func (m *MockConsoleRepository) GetConsoles(ctx context.Context) ([]model0.Console, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsoles", ctx)
	ret0, _ := ret[0].([]model0.Console)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsoles indicates an expected call of GetConsoles.
func (mr *MockConsoleRepositoryMockRecorder) GetConsoles(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsoles", reflect.TypeOf((*MockConsoleRepository)(nil).GetConsoles), ctx)
}

// MockAbilityRepository is a mock of AbilityRepository interface.
type MockAbilityRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAbilityRepositoryMockRecorder
}

// MockAbilityRepositoryMockRecorder is the mock recorder for MockAbilityRepository.
type MockAbilityRepositoryMockRecorder struct {
	mock *MockAbilityRepository
}

// NewMockAbilityRepository creates a new mock instance.
func NewMockAbilityRepository(ctrl *gomock.Controller) *MockAbilityRepository {
	mock := &MockAbilityRepository{ctrl: ctrl}
	mock.recorder = &MockAbilityRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAbilityRepository) EXPECT() *MockAbilityRepositoryMockRecorder {
	return m.recorder
}

// CreateAbility mocks base method.
func (m *MockAbilityRepository) CreateAbility(ctx context.Context, ability *model.Ability) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAbility", ctx, ability)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAbility indicates an expected call of CreateAbility.
func (mr *MockAbilityRepositoryMockRecorder) CreateAbility(ctx, ability interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAbility", reflect.TypeOf((*MockAbilityRepository)(nil).CreateAbility), ctx, ability)
}

// DeleteAbilityByID mocks base method.
func (m *MockAbilityRepository) DeleteAbilityByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAbilityByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAbilityByID indicates an expected call of DeleteAbilityByID.
func (mr *MockAbilityRepositoryMockRecorder) DeleteAbilityByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAbilityByID", reflect.TypeOf((*MockAbilityRepository)(nil).DeleteAbilityByID), ctx, id)
}

// GetAbilitiesByHeroID mocks base method.
func (m *MockAbilityRepository) GetAbilitiesByHeroID(ctx context.Context, id string) ([]model.Ability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAbilitiesByHeroID", ctx, id)
	ret0, _ := ret[0].([]model.Ability)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAbilitiesByHeroID indicates an expected call of GetAbilitiesByHeroID.
func (mr *MockAbilityRepositoryMockRecorder) GetAbilitiesByHeroID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAbilitiesByHeroID", reflect.TypeOf((*MockAbilityRepository)(nil).GetAbilitiesByHeroID), ctx, id)
}

// GetAbilityByID mocks base method.
func (m *MockAbilityRepository) GetAbilityByID(ctx context.Context, id string) (*model.Ability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAbilityByID", ctx, id)
	ret0, _ := ret[0].(*model.Ability)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAbilityByID indicates an expected call of GetAbilityByID.
func (mr *MockAbilityRepositoryMockRecorder) GetAbilityByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAbilityByID", reflect.TypeOf((*MockAbilityRepository)(nil).GetAbilityByID), ctx, id)
}
