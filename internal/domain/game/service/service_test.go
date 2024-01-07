package game

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/pkg/universe"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

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
		spiderManGame.Universe, out.CreatedAt, out.UpdatedAt), out)
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
		spiderManGame.Universe, out.CreatedAt, out.UpdatedAt), out)
}
