package hero

import (
	"context"
	"fmt"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/cache"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/model"
	"github.com/google/uuid"
)

//go:generate mockgen -destination ../../../mock/hero_mock.go -package=mock -source=service.go
type Hero interface {
	CreateHero(ctx context.Context, request *dto.HeroRequest) (*dto.HeroResponse, error)
	UpdateHero(ctx context.Context, id string, request *dto.HeroRequest) error
	GetHeroByID(ctx context.Context, id string) (*dto.HeroResponse, error)
	DeleteHeroByID(ctx context.Context, id string) (err error)
	AddAbilityToHero(ctx context.Context, abilityID, heroID string) error
}

type service struct {
	repository repository.HeroRepository
	cache      cache.HeroCache
	log        log.Log
}

func New(repository repository.Repository, cache cache.Cache,
	log log.Log) *service {
	return &service{
		repository: repository,
		cache:      cache,
		log:        log,
	}
}
func (s *service) CreateHero(ctx context.Context, req *dto.HeroRequest) (*dto.HeroResponse, error) {
	resp, err := s.createHero(ctx, req)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}
	return resp, nil
}

func (s *service) createHero(ctx context.Context, request *dto.HeroRequest) (*dto.HeroResponse, error) {
	hero := model.NewHero(uuid.NewString(), request)
	if err := s.repository.CreateHero(ctx, hero); err != nil {
		return nil, exception.New(fmt.Sprintf("createHero\n%s", err.Error()), err)
	}

	if err := s.cache.SetHero(ctx, hero); err != nil {
		return nil, exception.New(fmt.Sprintf("setHero\n%s", err.Error()), exception.ErrInternalServer)
	}

	return dto.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName, hero.Universe,
		hero.Hero, hero.CreatedAt, hero.UpdatedAt, hero.Team), nil
}

func (s *service) UpdateHero(ctx context.Context, id string, req *dto.HeroRequest) error {
	if err := s.updateHero(ctx, id, req); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}
	return nil
}
func (s *service) updateHero(ctx context.Context, id string, request *dto.HeroRequest) error {
	hero := model.NewHero(id, request)
	hero.UpdatedAt = util.GerPointer(time.Now().UTC())
	if err := s.repository.UpdateHero(ctx, hero); err != nil {
		return exception.New(fmt.Sprintf("updateHero\n%s", err.Error()), err)
	}

	if err := s.cache.SetHero(ctx, hero); err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("setHero\n%w", err))
		if err = s.cache.DeleteHero(ctx, hero.ID); err != nil {
			s.log.SendErrorLog(ctx, fmt.Errorf("deleteHero\n%w", err))
		}
	}

	return nil
}

func (s *service) GetHeroByID(ctx context.Context, id string) (*dto.HeroResponse, error) {
	resp, err := s.getHeroByID(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) getHeroByID(ctx context.Context, id string) (*dto.HeroResponse, error) {
	hero, err := s.cache.GetHero(ctx, id)
	if err != nil {
		s.log.SendErrorLog(ctx, fmt.Errorf("getHero\n%w", err))
		hero, err = s.repository.GetHeroByID(ctx, id)
		if err != nil {
			return nil, exception.New(fmt.Sprintf("getHeroByID\n%s", err.Error()), err)
		}
		return dto.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName,
			hero.Universe, hero.Hero, hero.CreatedAt, hero.UpdatedAt, nil), nil
	}

	return dto.NewHeroResponse(hero.ID, hero.HeroName, hero.CivilName,
		hero.Universe, hero.Hero, hero.CreatedAt, hero.UpdatedAt, hero.Team), nil
}

func (s *service) DeleteHeroByID(ctx context.Context, id string) (err error) {
	if err := s.deleteHeroByID(ctx, id); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}
	return nil
}

func (s *service) deleteHeroByID(ctx context.Context, id string) (err error) {
	if err = s.cache.DeleteHero(ctx, id); err != nil {
		return exception.New(fmt.Sprintf("deleteHero\n%s", err.Error()), err)
	}

	if err = s.repository.DeleteHeroByID(ctx, id); err != nil {
		return exception.New(fmt.Sprintf("deleteHeroByID\n%s", err.Error()), err)
	}
	return
}

func (s *service) AddAbilityToHero(ctx context.Context, abilityID, heroID string) error {
	if err := s.addAbilityToHero(ctx, abilityID, heroID); err != nil {
		s.log.SendErrorLog(ctx, err)
		return err
	}

	return nil
}

func (s *service) addAbilityToHero(ctx context.Context, abilityID, heroID string) error {
	if err := s.repository.AddAbilityToHero(ctx, abilityID, heroID); err != nil {
		return exception.New(fmt.Sprintf("AddAbilityToHero\n%s", err.Error()), err)
	}
	return nil
}
