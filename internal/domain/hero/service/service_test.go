package hero

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const abilityID = "68ed1b87-ce4c-4645-a88a-144398e65db2"

var (
	superMan = &dto.HeroRequest{
		HeroName:  "Super-man",
		CivilName: "Clark Kent",
		Hero:      true,
		Universe:  "DC",
		Team:      nil,
	}

	ironman = &model.Hero{
		ID:        "96f15886-6570-4469-8d9e-e694733000d1",
		HeroName:  "iron man",
		CivilName: "tony stark",
		Hero:      true,
		Universe:  "Marvel",
		Team:      nil,
	}
)

func getMockContext(ctrl *gomock.Controller) context.Context {
	var (
		ctx = context.Background()
		l   = mock.NewMockLogger(ctrl)
	)
	l.EXPECT().Error(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	return log.AddLoggerInContext(ctx, l)
}

// Create
func TestServiceRegisterHeroSuccess(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		ctx  = getMockContext(ctrl)
		rep  = mock.NewRepositoryMock(ctrl)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	rep.EXPECT().CreateHero(gomock.Any(), gomock.Any()).
		Return(nil)
	c.EXPECT().SetHero(gomock.Any(), gomock.Any()).Return(nil)
	defer ctrl.Finish()
	s := &service{
		repository: rep,
		cache:      c,
	}
	out, err := s.CreateHero(ctx, superMan)
	assert.Equal(t, &dto.HeroResponse{
		ID:        out.ID,
		HeroName:  "super-man",
		CivilName: "clark kent",
		Hero:      true,
		Universe:  "DC",
		CreatedAt: out.CreatedAt,
		Team:      nil,
	}, out)

	assert.ErrorIs(t, nil, err)
}

func TestServiceRegisterHeroFailInternalServerError(t *testing.T) {
	var (
		ctrl     = gomock.NewController(t)
		ctx      = getMockContext(ctrl)
		rep      = mock.NewRepositoryMock(ctrl)
		c        = mock.NewMockCache(ctrl)
		expected *dto.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().CreateHero(gomock.Any(), gomock.Any()).
		Return(exception.ErrInternalServer)
	s := &service{
		repository: rep,
		cache:      c,
	}

	out, err := s.CreateHero(ctx, superMan)
	var errorWithTrace *exception.ErrorWithTrace
	assert.Equal(t, expected, out)
	assert.ErrorAs(t, err, &errorWithTrace)
	assert.ErrorIs(t, errorWithTrace.GetError(), exception.ErrInternalServer)
}

// Get
func TestServiceGetHeroByIDSuccess(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		ctx  = getMockContext(ctrl)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	c.EXPECT().GetHero(gomock.Any(), ironman.ID).Return(ironman, nil)

	s := &service{
		repository: nil,
		cache:      c,
	}
	out, err := s.GetHeroByID(ctx, ironman.ID)
	assert.Equal(t, &dto.HeroResponse{
		ID:        ironman.ID,
		HeroName:  ironman.HeroName,
		CivilName: ironman.CivilName,
		Hero:      true,
		Universe:  ironman.Universe,
		Team:      nil,
	}, out)

	assert.ErrorIs(t, nil, err)
}

func TestServiceGetHeroByIDFailHeroNotFoundError(t *testing.T) {
	var (
		ctrl     = gomock.NewController(t)
		ctx      = getMockContext(ctrl)
		rep      = mock.NewRepositoryMock(ctrl)
		c        = mock.NewMockCache(ctrl)
		expected *dto.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().GetHeroByID(gomock.Any(), ironman.ID).
		Return(nil, exception.ErrHeroNotFound)
	c.EXPECT().GetHero(gomock.Any(), ironman.ID).Return(nil, exception.ErrHeroNotFound)

	s := &service{
		repository: rep,
		cache:      c,
	}
	out, err := s.GetHeroByID(ctx, ironman.ID)
	var errorWithTrace *exception.ErrorWithTrace
	assert.Equal(t, expected, out)
	assert.ErrorAs(t, err, &errorWithTrace)
	assert.ErrorIs(t, errorWithTrace.GetError(), exception.ErrHeroNotFound)
}

func TestServiceGetHeroByIDFailInternalServerError(t *testing.T) {
	var (
		ctrl     = gomock.NewController(t)
		ctx      = getMockContext(ctrl)
		rep      = mock.NewRepositoryMock(ctrl)
		c        = mock.NewMockCache(ctrl)
		expected *dto.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().GetHeroByID(gomock.Any(), ironman.ID).
		Return(nil, exception.ErrInternalServer)
	c.EXPECT().GetHero(gomock.Any(), ironman.ID).Return(nil, exception.ErrInternalServer)

	s := &service{
		repository: rep,
		cache:      c,
	}
	out, err := s.GetHeroByID(ctx, ironman.ID)
	var errorWithTrace *exception.ErrorWithTrace

	assert.Equal(t, expected, out)
	assert.ErrorAs(t, err, &errorWithTrace)
	assert.ErrorIs(t, errorWithTrace.GetError(), exception.ErrInternalServer)
}

// Update

func TestServiceUpdateHeroSuccess(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		ctx  = getMockContext(ctrl)
		rep  = mock.NewRepositoryMock(ctrl)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	rep.EXPECT().UpdateHero(gomock.Any(), gomock.Any()).
		Return(nil)
	c.EXPECT().SetHero(gomock.Any(), gomock.Any()).Return(nil)

	s := &service{
		repository: rep,
		cache:      c,
	}
	err := s.UpdateHero(ctx, ironman.ID, &dto.HeroRequest{
		HeroName:  ironman.HeroName,
		CivilName: ironman.CivilName,
		Hero:      ironman.Hero,
		Universe:  ironman.Universe,
	})

	assert.ErrorIs(t, nil, err)
}

func TestServiceUpdateHeroFailInternalServerError(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		ctx  = getMockContext(ctrl)
		r    = mock.NewRepositoryMock(ctrl)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	r.EXPECT().UpdateHero(gomock.Any(), gomock.Any()).Return(exception.ErrInternalServer)
	s := &service{
		repository: r,
		cache:      c,
	}
	err := s.UpdateHero(ctx, ironman.ID, &dto.HeroRequest{
		HeroName:  ironman.HeroName,
		CivilName: ironman.CivilName,
		Hero:      ironman.Hero,
		Universe:  ironman.Universe,
	})

	var errWithTrace *exception.ErrorWithTrace
	assert.ErrorAs(t, err, &errWithTrace)
	assert.ErrorIs(t, errWithTrace.GetError(), exception.ErrInternalServer)
}

// Delete

func TestServiceDeleteHeroByIDSuccess(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		ctx  = getMockContext(ctrl)
		c    = mock.NewMockCache(ctrl)
		r    = mock.NewRepositoryMock(ctrl)
	)
	defer ctrl.Finish()
	c.EXPECT().DeleteHero(gomock.Any(), ironman.ID).Return(nil)
	r.EXPECT().DeleteHeroByID(gomock.Any(), ironman.ID).Return(nil)
	s := New(r, c)
	err := s.DeleteHeroByID(ctx, ironman.ID)
	assert.ErrorIs(t, nil, err)
}

func TestServiceDeleteHeroByIDFailInternalServerError(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		ctx  = getMockContext(ctrl)
		c    = mock.NewMockCache(ctrl)
	)
	c.EXPECT().DeleteHero(gomock.Any(), ironman.ID).Return(exception.ErrInternalServer)
	s := New(nil, c)
	err := s.DeleteHeroByID(ctx, ironman.ID)
	var errorWithTrace *exception.ErrorWithTrace
	assert.ErrorAs(t, err, &errorWithTrace)
	assert.ErrorIs(t, errorWithTrace.GetError(), exception.ErrInternalServer)
}

func TestServiceAddAbilityToHeroSuccess(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
		ctx  = getMockContext(ctrl)
		r    = mock.NewRepositoryMock(ctrl)
	)

	r.EXPECT().AddAbilityToHero(gomock.Any(), abilityID, ironman.ID).Return(nil)
	s := &service{
		repository: r,
	}
	err := s.AddAbilityToHero(ctx, abilityID, ironman.ID)
	assert.NoError(t, err)
}
