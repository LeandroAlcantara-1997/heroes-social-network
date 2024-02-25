package console

import (
	"context"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func getMockContext(ctrl *gomock.Controller) context.Context {
	var (
		ctx = context.Background()
		l   = mock.NewMockLogger(ctrl)
	)
	l.EXPECT().Error(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	return log.AddLoggerInContext(ctx, l)
}

func TestServiceCreateConsolesSuccess(t *testing.T) {
	var (
		ctrl           = gomock.NewController(t)
		ctx            = getMockContext(ctrl)
		repositoryMock = mock.NewMockConsoleRepository(ctrl)
	)

	repositoryMock.EXPECT().CreateConsoles(gomock.Any(), []model.Console{
		"Playstation1",
	}).Return(nil)

	s := &service{
		repository: repositoryMock,
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
		ctrl           = gomock.NewController(t)
		ctx            = getMockContext(ctrl)
		repositoryMock = mock.NewMockConsoleRepository(ctrl)
		expected       *dto.ConsoleResponse
	)

	repositoryMock.EXPECT().CreateConsoles(gomock.Any(), []model.Console{
		"Playstation1",
	}).Return(exception.ErrInternalServer)

	s := &service{
		repository: repositoryMock,
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
