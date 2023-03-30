package heroes

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/mock"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	superMan = &input.HeroRequest{
		HeroName:  "Super-man",
		CivilName: "Clark Kent",
		Hero:      true,
		Universe:  "DC",
		Team:      nil,
	}
	batman = &input.HeroRequest{
		HeroName:  "Batman",
		CivilName: "Bruce Wayne",
		Hero:      true,
		Universe:  "DCI",
		Team:      nil,
	}
)

func TestServiceRegisterHeroSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		rep  = mock.NewMockRepository(ctrl)
	)

	rep.EXPECT().CreateHero(ctx, gomock.Any()).
		Return(nil)

	s := &service{
		repository: rep,
	}
	out, err := s.RegisterHero(ctx, superMan)
	assert.Equal(t, &input.HeroResponse{
		Id:        out.Id,
		HeroName:  "Super-man",
		CivilName: "Clark Kent",
		Hero:      true,
		Universe:  "DC",
		Team:      nil,
	}, out)

	assert.ErrorIs(t, nil, err)
}

func TestServiceRegisterHeroFailInternalServerError(t *testing.T) {
	var (
		ctx       = context.Background()
		ctrl      = gomock.NewController(t)
		rep       = mock.NewMockRepository(ctrl)
		l         = mock.NewMockLog(ctrl)
		expected  *input.HeroResponse
		wantError = exception.New(exception.InternalServerError)
	)

	rep.EXPECT().CreateHero(ctx, gomock.Any()).
		Return(wantError)

	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	s := &service{
		repository: rep,
		log:        l,
	}

	out, err := s.RegisterHero(ctx, superMan)
	assert.Equal(t, expected, out)
	assert.Equal(t, wantError, err)
}

func TestServiceRegisterHeroFailInvalidField(t *testing.T) {
	var (
		ctx       = context.Background()
		ctrl      = gomock.NewController(t)
		l         = mock.NewMockLog(ctrl)
		expected  *input.HeroResponse
		wantError = exception.New(exception.InvalidFieldsError)
	)

	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	s := &service{
		repository: nil,
		log:        l,
	}

	out, err := s.RegisterHero(ctx, batman)
	assert.Equal(t, expected, out)
	assert.Equal(t, wantError, err)
}
