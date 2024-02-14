package team

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// const id = "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0"

var (
	teenTitans = &model.Team{
		ID:        "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0",
		Name:      "Teen Titans",
		Universe:  "DC",
		CreatedAt: time.Date(2020, 10, 15, 14, 30, 30, 30, time.UTC),
	}
	teenTitansResponse = &dto.TeamResponse{
		ID:        "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0",
		Name:      "Teen Titans",
		Universe:  "DC",
		CreatedAt: time.Date(2020, 10, 15, 14, 30, 30, 30, time.UTC),
	}
)

func TestServiceRegisterTeamSuccess(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewRepositoryMock(ctrl)
		cacheMock      = mock.NewMockCache(ctrl)
		logMock        = mock.NewMockLog(ctrl)
	)

	repositoryMock.EXPECT().CreateTeam(gomock.Any(), gomock.Any()).Return(nil)
	cacheMock.EXPECT().SetTeam(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	s := New(repositoryMock, cacheMock, logMock)
	out, err := s.CreateTeam(ctx, &dto.TeamRequest{
		Name:     "the avengers",
		Universe: "MARVEL",
	})
	assert.Equal(t, dto.NewTeamResponse(out.ID, "the avengers", "MARVEL", out.CreatedAt, nil), out)
	assert.NoError(t, err)
}

func TestServiceRegisterTeamFail(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewRepositoryMock(ctrl)
		logMock        = mock.NewMockLog(ctrl)
		expected       *dto.TeamResponse
	)

	repositoryMock.EXPECT().CreateTeam(gomock.Any(), gomock.Any()).Return(exception.ErrTeamAlredyExists)
	logMock.EXPECT().SendErrorLog(gomock.Any(), gomock.Any())
	s := New(repositoryMock, nil, logMock)
	out, err := s.CreateTeam(ctx, &dto.TeamRequest{
		Name:     "The Avengers",
		Universe: "MARVEL",
	})
	var errorWithTrace *exception.ErrorWithTrace
	assert.ErrorAs(t, err, &errorWithTrace)
	assert.Equal(t, expected, out)
	assert.ErrorIs(t, errorWithTrace.GetError(), exception.ErrTeamAlredyExists)
}

func TestServiceGetTeamByIDSuccessByCache(t *testing.T) {
	var (
		ctx       = context.Background()
		ctrl      = gomock.NewController(t)
		cacheMock = mock.NewMockCache(ctrl)
	)
	cacheMock.EXPECT().GetTeam(gomock.Any(), "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0").Return(teenTitans, nil)
	s := New(nil, cacheMock, nil)
	out, err := s.GetTeamByID(ctx, "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0")
	assert.Equal(t, teenTitansResponse, out)
	assert.ErrorIs(t, err, nil)
}

func TestServiceGetTeamByIDSuccessByRepository(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewRepositoryMock(ctrl)
		cacheMock      = mock.NewMockCache(ctrl)
		logMock        = mock.NewMockLog(ctrl)
	)
	cacheCall := cacheMock.EXPECT().GetTeam(gomock.Any(), "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0").Return(nil, errors.New("unexpected error"))
	logMock.EXPECT().SendErrorLog(gomock.Any(), gomock.Any()).After(cacheCall)
	repositoryMock.EXPECT().GetTeamByID(gomock.Any(), "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0").Return(teenTitans, nil)
	s := New(repositoryMock, cacheMock, logMock)
	out, err := s.GetTeamByID(ctx, "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0")
	assert.Equal(t, teenTitansResponse, out)
	assert.ErrorIs(t, err, nil)
}

func TestServiceDeleteTeamByIDSuccessTeamDeleted(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewRepositoryMock(ctrl)
		cacheMock      = mock.NewMockCache(ctrl)
	)

	cacheMock.EXPECT().DeleteTeam(gomock.Any(), teenTitans.ID).Return(nil)
	repositoryMock.EXPECT().DeleteTeamByID(gomock.Any(), teenTitans.ID).Return(nil)
	s := &service{
		repository: repositoryMock,
		cache:      cacheMock,
		log:        nil,
	}
	err := s.DeleteTeamByID(ctx, teenTitans.ID)
	assert.ErrorIs(t, err, nil)

}

func TestServiceDeleteTeamByIDFailTeamNotDeletedCache(t *testing.T) {
	var (
		ctx       = context.Background()
		ctrl      = gomock.NewController(t)
		cacheMock = mock.NewMockCache(ctrl)
		logMock   = mock.NewMockLog(ctrl)
	)

	cacheMock.EXPECT().DeleteTeam(gomock.Any(), teenTitans.ID).Return(exception.ErrInternalServer)
	logMock.EXPECT().SendErrorLog(gomock.Any(), gomock.Any())
	s := &service{
		repository: nil,
		cache:      cacheMock,
		log:        logMock,
	}
	err := s.DeleteTeamByID(ctx, teenTitans.ID)
	var errorWithTrace *exception.ErrorWithTrace
	assert.ErrorAs(t, err, &errorWithTrace)
	assert.ErrorIs(t, errorWithTrace.GetError(), exception.ErrInternalServer)
}

func TestServiceDeleteTeamByIDFailTeamNotDeletedByDatabase(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		cacheMock      = mock.NewMockCache(ctrl)
		repositoryMock = mock.NewRepositoryMock(ctrl)
		logMock        = mock.NewMockLog(ctrl)
	)

	cacheMock.EXPECT().DeleteTeam(gomock.Any(), teenTitans.ID).Return(nil)
	repositoryMock.EXPECT().DeleteTeamByID(gomock.Any(), teenTitans.ID).Return(exception.ErrInternalServer)
	logMock.EXPECT().SendErrorLog(gomock.Any(), gomock.Any())
	s := &service{
		repository: repositoryMock,
		cache:      cacheMock,
		log:        logMock,
	}
	err := s.DeleteTeamByID(ctx, teenTitans.ID)
	var errWithTrace *exception.ErrorWithTrace
	assert.ErrorAs(t, err, &errWithTrace)
	assert.ErrorIs(t, errWithTrace.GetError(), exception.ErrInternalServer)
}
