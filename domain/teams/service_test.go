package teams

import (
	"context"
	"errors"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/mock"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestServiceRegisterTeamSuccess(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewMockRepository(ctrl)
		cacheMock      = mock.NewMockCache(ctrl)
		logMock        = mock.NewMockLog(ctrl)
	)

	repositoryMock.EXPECT().CreateTeam(ctx, gomock.Any()).Return(nil)
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

	repositoryMock.EXPECT().CreateTeam(ctx, gomock.Any()).Return(errors.New("unexpected error"))
	logMock.EXPECT().SendErrorLog(ctx, "unexpected error")
	s := New(repositoryMock, nil, logMock)
	out, err := s.RegisterTeam(ctx, &team.TeamRequest{
		Name:     "The Avengers",
		Universe: "MARVEL",
	})
	assert.Equal(t, expected, out)
	assert.ErrorIs(t, err, exception.ErrInternalServer)
}
