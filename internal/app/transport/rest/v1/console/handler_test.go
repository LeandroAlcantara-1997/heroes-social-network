package console

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/model"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestHandlerPostConsoles(t *testing.T) {
	var (
		response, _ = json.Marshal(&dto.ConsoleResponse{
			Names: []model.Console{
				"Playstaton1",
			},
		})
		ctx, _ = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockConsole(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().CreateConsoles(ctx, &dto.ConsoleRequest{
		Names: []model.Console{
			"Playstaton1",
		},
	}).Return(&dto.ConsoleResponse{
		Names: []model.Console{
			"Playstaton1",
		},
	}, nil)
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/consoles",
		strings.NewReader(`{
			"consoles": [
				"Playstaton1"
			]
		}`),
	)
	h := &Handler{
		useCase: useCase,
	}
	h.postConsoles(ctx)
}

func TestHandlerGetConsoles(t *testing.T) {
	var (
		response, _ = json.Marshal(&dto.ConsoleResponse{
			Names: []model.Console{
				"Playstaton1",
			},
		})
		ctx, _ = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockConsole(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().GetConsoles(ctx).Return(&dto.ConsoleResponse{
		Names: []model.Console{
			"Playstaton1",
		},
	}, nil)
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/consoles",
		strings.NewReader(``),
	)
	h := &Handler{
		useCase: useCase,
	}
	h.getConsoles(ctx)
}
