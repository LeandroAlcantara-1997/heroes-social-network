package hero

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// const id = "96f15886-6570-4469-8d9e-e694733000d1"

var (
	superMan = &dto.HeroRequest{
		HeroName:  "Super-man",
		CivilName: "Clark Kent",
		Hero:      true,
		Universe:  "DC",
		Team:      nil,
	}

	batman = &dto.HeroRequest{
		HeroName:  "Batman",
		CivilName: "Bruce Wayne",
		Hero:      true,
		Universe:  "DCI",
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

// Create
func TestServiceRegisterHeroSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		rep  = mock.NewMockRepository(ctrl)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	rep.EXPECT().CreateHero(ctx, gomock.Any()).
		Return(nil)
	c.EXPECT().SetHero(ctx, gomock.Any()).Return(nil)
	defer ctrl.Finish()
	s := &service{
		repository: rep,
		cache:      c,
	}
	out, err := s.RegisterHero(ctx, superMan)
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
		ctx      = context.Background()
		ctrl     = gomock.NewController(t)
		rep      = mock.NewMockRepository(ctrl)
		l        = mock.NewMockLog(ctrl)
		c        = mock.NewMockCache(ctrl)
		expected *dto.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().CreateHero(ctx, gomock.Any()).
		Return(exception.ErrInternalServer)

	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	s := &service{
		repository: rep,
		log:        l,
		cache:      c,
	}

	out, err := s.RegisterHero(ctx, superMan)
	assert.Equal(t, expected, out)
	assert.ErrorIs(t, exception.ErrInternalServer, err)
}

// Get
func TestServiceGetHeroByIDSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	c.EXPECT().GetHero(ctx, ironman.ID).Return(ironman, nil)

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
		ctx      = context.Background()
		ctrl     = gomock.NewController(t)
		rep      = mock.NewMockRepository(ctrl)
		l        = mock.NewMockLog(ctrl)
		c        = mock.NewMockCache(ctrl)
		expected *dto.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().GetHeroByID(ctx, ironman.ID).
		Return(nil, exception.ErrHeroNotFound)
	l.EXPECT().SendErrorLog(ctx, gomock.Any()).AnyTimes()
	c.EXPECT().GetHero(ctx, ironman.ID).Return(nil, exception.ErrHeroNotFound)

	s := &service{
		repository: rep,
		log:        l,
		cache:      c,
	}
	out, err := s.GetHeroByID(ctx, ironman.ID)
	assert.Equal(t, expected, out)

	assert.ErrorIs(t, exception.ErrHeroNotFound, err)
}

func TestServiceGetHeroByIDFailInternalServerError(t *testing.T) {
	var (
		ctx      = context.Background()
		ctrl     = gomock.NewController(t)
		rep      = mock.NewMockRepository(ctrl)
		l        = mock.NewMockLog(ctrl)
		c        = mock.NewMockCache(ctrl)
		expected *dto.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().GetHeroByID(ctx, ironman.ID).
		Return(nil, exception.ErrInternalServer)
	l.EXPECT().SendErrorLog(ctx, gomock.Any()).AnyTimes()
	c.EXPECT().GetHero(ctx, ironman.ID).Return(nil, exception.ErrInternalServer)

	s := &service{
		repository: rep,
		log:        l,
		cache:      c,
	}
	out, err := s.GetHeroByID(ctx, ironman.ID)
	assert.Equal(t, expected, out)

	assert.ErrorIs(t, exception.ErrInternalServer, err)
}

// Update

func TestServiceUpdateHeroSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		rep  = mock.NewMockRepository(ctrl)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	rep.EXPECT().UpdateHero(ctx, gomock.Any()).
		Return(nil)
	c.EXPECT().SetHero(ctx, gomock.Any()).Return(nil)

	s := &service{
		repository: rep,
		log:        nil,
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
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		l    = mock.NewMockLog(ctrl)
		r    = mock.NewMockRepository(ctrl)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	r.EXPECT().UpdateHero(ctx, gomock.Any()).Return(exception.ErrInternalServer)
	c.EXPECT().DeleteHero(ctx, gomock.Any()).Return(nil)
	s := &service{
		repository: r,
		log:        l,
		cache:      c,
	}
	err := s.UpdateHero(ctx, ironman.ID, &dto.HeroRequest{
		HeroName:  ironman.HeroName,
		CivilName: ironman.CivilName,
		Hero:      ironman.Hero,
		Universe:  ironman.Universe,
	})

	assert.ErrorIs(t, exception.ErrInternalServer, err)
}

// Delete

func TestServiceDeleteHeroByIDSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		c    = mock.NewMockCache(ctrl)
		r    = mock.NewMockRepository(ctrl)
	)
	defer ctrl.Finish()
	c.EXPECT().DeleteHero(ctx, ironman.ID).Return(nil)
	r.EXPECT().DeleteHeroByID(ctx, ironman.ID).Return(nil)
	s := New(r, c, nil)
	err := s.DeleteHeroByID(ctx, ironman.ID)
	assert.ErrorIs(t, nil, err)
}

func TestServiceDeleteHeroByIDFailInternalServerError(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		c    = mock.NewMockCache(ctrl)
		l    = mock.NewMockLog(ctrl)
	)
	c.EXPECT().DeleteHero(ctx, ironman.ID).Return(exception.ErrInternalServer)
	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	s := New(nil, c, l)
	err := s.DeleteHeroByID(ctx, ironman.ID)
	assert.ErrorIs(t, exception.ErrInternalServer, err)
}
