package teams

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/mock"
	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const id = "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0"

var (
	teenTitans = &model.Team{
		ID:        "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0",
		Name:      "Teen Titans",
		Universe:  "DC",
		CreatedAt: time.Date(2020, 10, 15, 14, 30, 30, 30, time.UTC),
	}
	teenTitansResponse = &team.TeamResponse{
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
		repositoryMock = mock.NewMockRepository(ctrl)
		cacheMock      = mock.NewMockCache(ctrl)
		logMock        = mock.NewMockLog(ctrl)
	)

	repositoryMock.EXPECT().CreateTeam(ctx, gomock.Any()).Return(&model.Team{
		ID:       id,
		Name:     "The Avengers",
		Universe: "MARVEL",
	}, nil)
	cacheMock.EXPECT().SetTeam(ctx, gomock.Any()).Return(nil)
	s := New(repositoryMock, cacheMock, logMock)
	out, err := s.RegisterTeam(ctx, &team.TeamRequest{
		Name:     "The Avengers",
		Universe: "MARVEL",
	})
	assert.Equal(t, team.NewTeamResponse(out.ID, "The Avengers", "MARVEL", out.CreatedAt, nil), out)
	assert.ErrorIs(t, err, nil)
}

func TestServiceRegisterTeamFail(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewMockRepository(ctrl)
		logMock        = mock.NewMockLog(ctrl)
		expected       *team.TeamResponse
	)

	repositoryMock.EXPECT().CreateTeam(ctx, gomock.Any()).Return(nil, errors.New("unexpected error"))
	logMock.EXPECT().SendErrorLog(ctx, errors.New("unexpected error"))
	s := New(repositoryMock, nil, logMock)
	out, err := s.RegisterTeam(ctx, &team.TeamRequest{
		Name:     "The Avengers",
		Universe: "MARVEL",
	})
	assert.Equal(t, expected, out)
	assert.ErrorIs(t, err, exception.ErrInternalServer)
}

func TestServiceGetTeamByIDSuccessByCache(t *testing.T) {
	var (
		ctx       = context.Background()
		ctrl      = gomock.NewController(t)
		cacheMock = mock.NewMockCache(ctrl)
	)
	cacheMock.EXPECT().GetTeam(ctx, "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0").Return(teenTitans, nil)
	s := New(nil, cacheMock, nil)
	out, err := s.GetTeamByID(ctx, "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0")
	assert.Equal(t, teenTitansResponse, out)
	assert.ErrorIs(t, err, nil)
}

func TestServiceGetTeamByIDSuccessByRepository(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewMockRepository(ctrl)
		cacheMock      = mock.NewMockCache(ctrl)
		logMock        = mock.NewMockLog(ctrl)
	)
	cacheCall := cacheMock.EXPECT().GetTeam(ctx, "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0").Return(nil, errors.New("unexpected error"))
	logMock.EXPECT().SendErrorLog(ctx, errors.New("unexpected error")).After(cacheCall)
	repositoryMock.EXPECT().GetTeamByID(ctx, "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0").Return(teenTitans, nil)
	s := New(repositoryMock, cacheMock, logMock)
	out, err := s.GetTeamByID(ctx, "0c2ab516-d1b9-4ba4-bbf2-a7b77b21e8a0")
	assert.Equal(t, teenTitansResponse, out)
	assert.ErrorIs(t, err, nil)
}

func TestServiceDeleteTeamByIDSuccessTeamDeleted(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewMockRepository(ctrl)
		cacheMock      = mock.NewMockCache(ctrl)
	)

	cacheMock.EXPECT().DeleteTeam(ctx, teenTitans.ID).Return(nil)
	repositoryMock.EXPECT().DeleteTeamByID(ctx, teenTitans.ID).Return(nil)
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

	cacheMock.EXPECT().DeleteTeam(ctx, teenTitans.ID).Return(exception.ErrInternalServer)
	logMock.EXPECT().SendErrorLog(ctx, exception.ErrInternalServer)
	s := &service{
		repository: nil,
		cache:      cacheMock,
		log:        logMock,
	}
	err := s.DeleteTeamByID(ctx, teenTitans.ID)
	assert.ErrorIs(t, err, exception.ErrInternalServer)
}

func TestServiceDeleteTeamByIDFailTeamNotDeletedByDatabase(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		cacheMock      = mock.NewMockCache(ctrl)
		repositoryMock = mock.NewMockRepository(ctrl)
		logMock        = mock.NewMockLog(ctrl)
	)

	cacheMock.EXPECT().DeleteTeam(ctx, teenTitans.ID).Return(nil)
	repositoryMock.EXPECT().DeleteTeamByID(ctx, teenTitans.ID).Return(exception.ErrInternalServer)
	logMock.EXPECT().SendErrorLog(ctx, exception.ErrInternalServer)
	s := &service{
		repository: repositoryMock,
		cache:      cacheMock,
		log:        logMock,
	}
	err := s.DeleteTeamByID(ctx, teenTitans.ID)
	assert.ErrorIs(t, err, exception.ErrInternalServer)
}
