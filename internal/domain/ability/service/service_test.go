package ability

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const id = "68ed1b87-ce4c-4645-a88a-144398e65db2"

func TestServiceCreateAbilitySuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		r    = mock.NewMockAbilityRepository(ctrl)
		c    = mock.NewMockAbilityCache(ctrl)
	)

	r.EXPECT().CreateAbility(gomock.Any(), gomock.Any())
	c.EXPECT().SetAbility(gomock.Any(), gomock.Any()).Return(nil)
	s := &service{
		repository: r,
		cache:      c,
	}
	out, err := s.CreateAbility(ctx, &dto.AbilityRequest{
		Description: "fly",
	})

	assert.Equal(t, "fly", out.Description)
	assert.NoError(t, err)

}

func TestServiceGetAbilityByIDSuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		c    = mock.NewMockAbilityCache(ctrl)
	)
	c.EXPECT().GetAbility(gomock.Any(), id).Return(&model.Ability{
		ID:          id,
		Description: "fly",
	}, nil)

	s := &service{
		cache: c,
	}
	out, err := s.GetAbilityByID(ctx, id)

	assert.Equal(t, "fly", out.Description)
	assert.Equal(t, id, out.ID)
	assert.NoError(t, err)
}

func TestServiceGetAbilitiesByHeroID(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		r    = mock.NewMockAbilityRepository(ctrl)
	)
	r.EXPECT().GetAbilitiesByHeroID(gomock.Any(), id).Return([]model.Ability{
		{
			ID:          id,
			Description: "fly",
		},
	}, nil)

	s := &service{
		repository: r,
	}
	out, err := s.GetAbilitiesByHeroID(ctx, id)

	assert.Equal(t, "fly", out[0].Description)
	assert.Equal(t, id, out[0].ID)
	assert.NoError(t, err)
}

func TestServiceDeleteAbilitySuccess(t *testing.T) {
	var (
		ctx  = context.Background()
		ctrl = gomock.NewController(t)
		r    = mock.NewMockAbilityRepository(ctrl)
		c    = mock.NewMockAbilityCache(ctrl)
	)
	r.EXPECT().DeleteAbilityByID(gomock.Any(), id).Return(nil)
	c.EXPECT().DeleteAbility(gomock.Any(), id).Return(nil)
	s := &service{
		repository: r,
		cache:      c,
	}
	err := s.DeleteAbility(ctx, id)
	assert.NoError(t, err)
}
