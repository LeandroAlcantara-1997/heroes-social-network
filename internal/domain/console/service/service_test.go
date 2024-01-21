package console

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreateConsolesSuccess(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewMockConsoleRepository(ctrl)
	)

	repositoryMock.EXPECT().CreateConsoles(ctx, []model.Console{
		"Playstation1",
	}).Return(nil)

	s := &service{
		repository: repositoryMock,
		logger:     nil,
	}
	out, err := s.CreateConsoles(ctx, &dto.ConsoleRequest{
		Names: []model.Console{
			"Playstation1",
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, out.Names[0], model.Console("Playstation1"))
}

func TestServiceCreateConsolesFailRepositoryError(t *testing.T) {
	var (
		ctx            = context.Background()
		ctrl           = gomock.NewController(t)
		repositoryMock = mock.NewMockConsoleRepository(ctrl)
		loggerMock     = mock.NewMockLog(ctrl)
		expected       *dto.ConsoleResponse
	)

	repositoryMock.EXPECT().CreateConsoles(ctx, []model.Console{
		"Playstation1",
	}).Return(exception.ErrInternalServer)
	loggerMock.EXPECT().SendErrorLog(ctx, gomock.Any())

	s := &service{
		repository: repositoryMock,
		logger:     loggerMock,
	}
	out, err := s.CreateConsoles(ctx, &dto.ConsoleRequest{
		Names: []model.Console{
			"Playstation1",
		},
	})

	var domainError *exception.ErrorWithTrace
	assert.ErrorAs(t, err, &domainError)
	assert.ErrorIs(t, domainError.GetError(), exception.ErrInternalServer)
	assert.Equal(t, expected, out)
}
