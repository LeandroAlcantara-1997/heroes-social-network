package ability

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"go.opentelemetry.io/otel"
)

//go:generate mockgen -destination ../../../mock/ability_mock.go -package=mock -source=service.go
type Ability interface {
	CreateAbility(ctx context.Context, req *dto.AbilityRequest) (*dto.AbilityResponse, error)
	GetAbilityByID(ctx context.Context, id string) (*dto.AbilityResponse, error)
	GetAbilitiesByHeroID(ctx context.Context, id string) ([]dto.AbilityResponse, error)
	DeleteAbility(ctx context.Context, id string) error
}

type service struct {
	repository repository.AbilityRepository
	cache      cache.AbilityCache
	logger     log.Log
}

func New(repository repository.AbilityRepository, cache cache.AbilityCache, logger log.Log) *service {
	return &service{
		repository: repository,
		cache:      cache,
		logger:     logger,
	}
}

func (s *service) CreateAbility(ctx context.Context, req *dto.AbilityRequest) (*dto.AbilityResponse, error) {
	ctx, span := otel.Tracer("ability").Start(ctx, "createAbility")
	defer span.End()
	resp, err := s.createAbility(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *service) createAbility(ctx context.Context, req *dto.AbilityRequest) (*dto.AbilityResponse, error) {
	ability := model.NewAbility(req)
	if err := s.repository.CreateAbility(ctx, ability); err != nil {
		return nil, err
	}

	if err := s.cache.SetAbility(ctx, ability); err != nil {
		s.logger.SendErrorLog(ctx, err)
	}
	return dto.NewAbilityResponse(
		ability.ID,
		ability.Description,
		ability.CreatedAt,
		ability.UpdatedAt,
	), nil
}

func (s *service) GetAbilityByID(ctx context.Context, id string) (*dto.AbilityResponse, error) {
	ctx, span := otel.Tracer("ability").Start(ctx, "getAbilityByID")
	defer span.End()
	resp, err := s.getAbilityByID(ctx, id)
	if err != nil {
		s.logger.SendErrorLog(ctx, err)
		return nil, err
	}
	return resp, nil
}

func (s *service) getAbilityByID(ctx context.Context, id string) (*dto.AbilityResponse, error) {
	ability, err := s.cache.GetAbility(ctx, id)
	if err != nil {
		s.logger.SendErrorLog(ctx, exception.New(fmt.Sprintf("getAbility\n%s", err.Error()), err))
		ability, err = s.repository.GetAbilityByID(ctx, id)
		if err != nil {
			return nil, exception.New(fmt.Sprintf("GetAbilityByID\n%s", err.Error()), err)
		}
	}
	return dto.NewAbilityResponse(ability.ID, ability.Description, ability.CreatedAt, ability.UpdatedAt), nil
}

func (s *service) GetAbilitiesByHeroID(ctx context.Context, id string) ([]dto.AbilityResponse, error) {
	ctx, span := otel.Tracer("ability").Start(ctx, "getAbilitiesByHeroID")
	defer span.End()
	resp, err := s.getAbilitiesByHeroID(ctx, id)
	if err != nil {
		s.logger.SendErrorLog(ctx, err)
		return nil, err
	}
	return resp, nil
}

func (s *service) getAbilitiesByHeroID(ctx context.Context, id string) ([]dto.AbilityResponse, error) {
	abilities, err := s.repository.GetAbilitiesByHeroID(ctx, id)
	if err != nil {
		return nil, exception.New(fmt.Sprintf("getAbilityByHeroID\n%s", err.Error()), err)
	}

	var resp = make([]dto.AbilityResponse, len(abilities))
	for a := range abilities {
		resp[a] = dto.AbilityResponse{
			ID:          abilities[a].ID,
			Description: abilities[a].Description,
			CreatedAt:   abilities[a].CreatedAt,
			UpdatedAt:   abilities[a].UpdatedAt,
		}
	}
	return resp, nil
}

func (s *service) DeleteAbility(ctx context.Context, id string) error {
	ctx, span := otel.Tracer("ability").Start(ctx, "deleteAbility")
	defer span.End()
	if err := s.deleteAbility(ctx, id); err != nil {
		return err
	}

	return nil
}

func (s *service) deleteAbility(ctx context.Context, id string) error {
	if err := s.cache.DeleteAbility(ctx, id); err != nil {
		return exception.New(fmt.Sprintf("deleteAbility\n%s", err.Error()), err)
	}

	if err := s.repository.DeleteAbilityByID(ctx, id); err != nil {
		return exception.New(fmt.Sprintf("deleteAbilityByID\n%s", err.Error()), err)
	}

	return nil
}
