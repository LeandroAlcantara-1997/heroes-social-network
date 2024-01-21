package game

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/pkg/universe"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const id = "b4606b93-15a2-4314-9ffd-e84c9b5fe8b8"

var (
	spiderManGame = &dto.GameRequest{
		Name:        "Spider Man",
		ReleaseYear: 2023,
		HeroID:      util.GerPointer("b4606b93-15a2-4314-9ffd-e84c9b5fe8b8"),
		Universe:    universe.Marvel,
	}
)

func TestServiceCreateSuccess(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewMockGameRepository(ctrl)
		cacheMock      = mock.NewMockGameCache(ctrl)
	)

	repositoryMock.EXPECT().CreateGame(ctx, gomock.Any()).Return(nil)
	cacheMock.EXPECT().SetGame(ctx, gomock.Any()).Return(nil)
	s := &service{
		repository: repositoryMock,
		cache:      cacheMock,
		log:        nil,
	}
	out, err := s.CreateGame(ctx, spiderManGame)

	assert.NoError(t, err)
	assert.Equal(t, dto.NewGameResponse(out.ID, spiderManGame.Name, spiderManGame.ReleaseYear,
		spiderManGame.TeamID, spiderManGame.HeroID,
		spiderManGame.Universe, out.CreatedAt, out.UpdatedAt, out.Consoles), out)
}

func TestServiceCreateFail(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewMockGameRepository(ctrl)
		cacheMock      = mock.NewMockGameCache(ctrl)
	)

	repositoryMock.EXPECT().CreateGame(ctx, gomock.Any()).Return(nil)
	cacheMock.EXPECT().SetGame(ctx, gomock.Any()).Return(nil)
	s := &service{
		repository: repositoryMock,
		cache:      cacheMock,
		log:        nil,
	}
	out, err := s.CreateGame(ctx, spiderManGame)

	assert.NoError(t, err)
	assert.Equal(t, dto.NewGameResponse(out.ID, spiderManGame.Name, spiderManGame.ReleaseYear,
		spiderManGame.TeamID, spiderManGame.HeroID,
		spiderManGame.Universe, out.CreatedAt, out.UpdatedAt, out.Consoles), out)
}

func TestServiceUpdateGameSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		rep  = mock.NewMockGameRepository(ctrl)
		c    = mock.NewMockGameCache(ctrl)
	)

	rep.EXPECT().UpdateGame(ctx, gomock.Any()).Return(nil)
	c.EXPECT().SetGame(ctx, gomock.Any())
	s := &service{
		repository: rep,
		cache:      c,
	}
	err := s.UpdateGame(ctx, id, &dto.GameRequest{
		Name:        "Spider Man 3",
		ReleaseYear: 2023,
		HeroID:      util.GerPointer(id),
		Universe:    universe.Marvel,
	})
	assert.NoError(t, err)
}

func TestServiceGetByIDSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		rep  = mock.NewMockGameRepository(ctrl)
		c    = mock.NewMockGameCache(ctrl)
	)
	c.EXPECT().GetGame(ctx, id).Return(model.NewGame(id, spiderManGame), nil)
	s := &service{
		repository: rep,
		cache:      c,
	}
	out, err := s.GetByID(ctx, id)
	assert.Equal(t, spiderManGame.Name, out.Name)
	assert.NoError(t, err)
}

func TestServiceDeleteSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		rep  = mock.NewMockGameRepository(ctrl)
		c    = mock.NewMockGameCache(ctrl)
	)

	rep.EXPECT().DeleteGameByID(ctx, id).Return(nil)
	c.EXPECT().DeleteGame(ctx, id)
	s := &service{
		repository: rep,
		cache:      c,
	}
	err := s.Delete(ctx, id)
	assert.NoError(t, err)

}
