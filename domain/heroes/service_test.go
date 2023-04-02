package heroes

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/mock"
	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
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

	ironman = &model.Hero{
		Id:        "96f15886-6570-4469-8d9e-e694733000d1",
		HeroName:  "Iron Man",
		CivilName: "Tony Stark",
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
	c.EXPECT().Set(ctx, gomock.Any()).Return(nil)
	defer ctrl.Finish()
	s := &service{
		repository: rep,
		cache:      c,
	}
	out, err := s.RegisterHero(ctx, superMan)
	assert.Equal(t, &input.HeroResponse{
		Id:        out.Id,
		HeroName:  "Super-man",
		CivilName: "Clark Kent",
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
		expected *input.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().CreateHero(ctx, gomock.Any()).
		Return(exception.ErrInternalServer)
	c.EXPECT().Delete(ctx, gomock.Any()).Return(nil)

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

func TestServiceRegisterHeroFailInvalidField(t *testing.T) {
	var (
		ctx      = context.Background()
		ctrl     = gomock.NewController(t)
		l        = mock.NewMockLog(ctrl)
		expected *input.HeroResponse
	)
	defer ctrl.Finish()
	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	s := &service{
		repository: nil,
		log:        l,
	}

	out, err := s.RegisterHero(ctx, batman)
	assert.Equal(t, expected, out)
	assert.Equal(t, exception.ErrInvalidFields, err)
}

// Get
func TestServiceGetHeroByIDSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		c    = mock.NewMockCache(ctrl)
	)
	defer ctrl.Finish()
	c.EXPECT().Get(ctx, ironman.Id).Return(ironman, nil)

	s := &service{
		repository: nil,
		cache:      c,
	}
	out, err := s.GetHeroByID(ctx, ironman.Id)
	assert.Equal(t, &input.HeroResponse{
		Id:        ironman.Id,
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
		expected *input.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().GetHeroByID(ctx, ironman.Id).
		Return(nil, exception.ErrHeroNotFound)
	l.EXPECT().SendErrorLog(ctx, gomock.Any()).AnyTimes()
	c.EXPECT().Get(ctx, ironman.Id).Return(nil, exception.ErrHeroNotFound)

	s := &service{
		repository: rep,
		log:        l,
		cache:      c,
	}
	out, err := s.GetHeroByID(ctx, ironman.Id)
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
		expected *input.HeroResponse
	)
	defer ctrl.Finish()
	rep.EXPECT().GetHeroByID(ctx, ironman.Id).
		Return(nil, exception.ErrInternalServer)
	l.EXPECT().SendErrorLog(ctx, gomock.Any()).AnyTimes()
	c.EXPECT().Get(ctx, ironman.Id).Return(nil, exception.ErrInternalServer)

	s := &service{
		repository: rep,
		log:        l,
		cache:      c,
	}
	out, err := s.GetHeroByID(ctx, ironman.Id)
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
	c.EXPECT().Set(ctx, gomock.Any()).Return(nil)

	s := &service{
		repository: rep,
		log:        nil,
		cache:      c,
	}
	out, err := s.UpdateHero(ctx, ironman.Id, &input.HeroRequest{
		HeroName:  ironman.HeroName,
		CivilName: ironman.CivilName,
		Hero:      ironman.Hero,
		Universe:  ironman.Universe,
	})
	assert.Equal(t, &input.HeroResponse{
		Id:        ironman.Id,
		HeroName:  ironman.HeroName,
		CivilName: ironman.CivilName,
		Hero:      ironman.Hero,
		Universe:  ironman.Universe,
		CreatedAt: out.CreatedAt,
		UpdatedAt: out.UpdatedAt,
	}, out)

	assert.ErrorIs(t, nil, err)
}

func TestServiceUpdateHeroFailInvalidField(t *testing.T) {
	var (
		ctx      = context.Background()
		ctrl     = gomock.NewController(t)
		l        = mock.NewMockLog(ctrl)
		expected *input.HeroResponse
	)
	defer ctrl.Finish()
	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	s := &service{
		repository: nil,
		log:        l,
		cache:      nil,
	}
	out, err := s.UpdateHero(ctx, ironman.Id, &input.HeroRequest{
		HeroName:  ironman.HeroName,
		CivilName: ironman.CivilName,
		Hero:      ironman.Hero,
		Universe:  "Marrvel",
	})
	assert.Equal(t, expected, out)

	assert.ErrorIs(t, exception.ErrInvalidFields, err)
}

func TestServiceUpdateHeroFailInternalServerError(t *testing.T) {
	var (
		ctx      = context.Background()
		ctrl     = gomock.NewController(t)
		l        = mock.NewMockLog(ctrl)
		r        = mock.NewMockRepository(ctrl)
		c        = mock.NewMockCache(ctrl)
		expected *input.HeroResponse
	)
	defer ctrl.Finish()
	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	r.EXPECT().UpdateHero(ctx, gomock.Any()).Return(exception.ErrInternalServer)
	c.EXPECT().Delete(ctx, gomock.Any()).Return(nil)
	s := &service{
		repository: r,
		log:        l,
		cache:      c,
	}
	out, err := s.UpdateHero(ctx, ironman.Id, &input.HeroRequest{
		HeroName:  ironman.HeroName,
		CivilName: ironman.CivilName,
		Hero:      ironman.Hero,
		Universe:  ironman.Universe,
	})
	assert.Equal(t, expected, out)

	assert.ErrorIs(t, exception.ErrInternalServer, err)
}

// Delete

func TestServiceDeleteHeroByIDSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		c    = mock.NewMockCache(ctrl)
		r    = mock.NewMockHeroRepository(ctrl)
	)
	defer ctrl.Finish()
	c.EXPECT().Delete(ctx, ironman.Id).Return(nil)
	r.EXPECT().DeleteHeroByID(ctx, ironman.Id).Return(nil)
	s := New(r, c, nil)
	err := s.DeleteHeroByID(ctx, ironman.Id)
	assert.ErrorIs(t, nil, err)
}

func TestServiceDeleteHeroByIDFailInternalServerError(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		c    = mock.NewMockCache(ctrl)
		l    = mock.NewMockLog(ctrl)
	)
	c.EXPECT().Delete(ctx, ironman.Id).Return(exception.ErrInternalServer)
	l.EXPECT().SendErrorLog(ctx, gomock.Any())
	s := New(nil, c, l)
	err := s.DeleteHeroByID(ctx, ironman.Id)
	assert.ErrorIs(t, exception.ErrInternalServer, err)
}
