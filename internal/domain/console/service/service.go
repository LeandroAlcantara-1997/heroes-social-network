package console

import (
	"context"
	"fmt"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/repository"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"go.opentelemetry.io/otel"
)

//go:generate mockgen -destination ../../../mock/console_mock.go -package=mock -source=service.go
type Console interface {
	CreateConsoles(ctx context.Context, req *dto.ConsoleRequest) (*dto.ConsoleResponse, error)
	GetConsoles(ctx context.Context) (*dto.ConsoleResponse, error)
}

type service struct {
	repository repository.ConsoleRepository
	logger     log.Log
}

func New(repository repository.ConsoleRepository, log log.Log) *service {
	return &service{
		repository: repository,
		logger:     log,
	}
}

func (s *service) CreateConsoles(ctx context.Context, req *dto.ConsoleRequest) (*dto.ConsoleResponse, error) {
	ctx, span := otel.Tracer("console").Start(ctx, "createConsoles")
	defer span.End()
	resp, err := s.createConsoles(ctx, req)
	if err != nil {
		s.logger.SendErrorLog(ctx, err)
		return nil, err
	}
	return resp, nil
}

func (s *service) createConsoles(ctx context.Context, req *dto.ConsoleRequest) (*dto.ConsoleResponse, error) {
	if err := s.repository.CreateConsoles(ctx, req.Names); err != nil {
		return nil, exception.New(fmt.Sprintf("createConsoles\n%s", err.Error()), err)
	}

	return &dto.ConsoleResponse{
		Names: req.Names,
	}, nil
}

func (s *service) GetConsoles(ctx context.Context) (*dto.ConsoleResponse, error) {
	ctx, span := otel.Tracer("console").Start(ctx, "getConsoles")
	defer span.End()
	resp, err := s.getConsoles(ctx)
	if err != nil {
		return nil, exception.New(fmt.Sprintf("getConsoles\n%s", err.Error()), err)
	}
	return resp, nil
}

func (s *service) getConsoles(ctx context.Context) (*dto.ConsoleResponse, error) {
	consoles, err := s.repository.GetConsoles(ctx)
	if err != nil {
		return nil, err
	}

	return &dto.ConsoleResponse{
		Names: consoles,
	}, nil
}
