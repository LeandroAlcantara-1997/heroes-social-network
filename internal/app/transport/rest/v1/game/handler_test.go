package game

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/pkg/universe"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

const id = "b4606b93-15a2-4314-9ffd-e84c9b5fe8b8"

var (
	ironManTwoGameRequest = &dto.GameRequest{
		Name:        "Iron Man 2",
		ReleaseYear: 2008,
		HeroID:      util.GerPointer(id),
		Universe:    universe.Marvel,
	}

	ironManTwoGameResponse = &dto.GameResponse{
		ID:          id,
		Name:        ironManTwoGameRequest.Name,
		ReleaseYear: ironManTwoGameRequest.ReleaseYear,
		Universe:    universe.Marvel,
		HeroID:      util.GerPointer(id),
		CreatedAt:   time.Now().UTC(),
	}
)

func TestHandlerPostGameSuccess(t *testing.T) {
	var (
		response, _ = json.Marshal(&ironManTwoGameResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockGame(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().CreateGame(ctx, ironManTwoGameRequest).
		Return(ironManTwoGameResponse, nil)
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/games",
		strings.NewReader(`{
					"name": "Iron Man 2",
					"heroId":  "b4606b93-15a2-4314-9ffd-e84c9b5fe8b8",
					"universe": "MARVEL",
					"releaseYear": 2008
				}`),
	)
	h := &Handler{
		useCase: useCase,
	}
	h.postGame(ctx)
}

func TestHandlerPostGameFailInvalidField(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrInvalidRequest))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusBadRequest,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockGame(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/games",
		strings.NewReader(`{
					"name": "",
					"heroId":  "b4606b93-15a2-4314-9ffd-e84c9b5fe8b8",
					"universe": "MARVEL",
					"releaseYear": 2008
				}`),
	)
	h := &Handler{
		useCase: useCase,
	}
	h.postGame(ctx)
}

func TestHandlerGetGameSuccess(t *testing.T) {
	var (
		response, _ = json.Marshal(&ironManTwoGameResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockGame(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().GetByID(ctx, id).
		Return(ironManTwoGameResponse, nil)
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/games?id=b4606b93-15a2-4314-9ffd-e84c9b5fe8b8",
		strings.NewReader(``),
	)
	h := &Handler{
		useCase: useCase,
	}
	h.getGame(ctx)
}

func TestHandlerPutGameSuccess(t *testing.T) {
	var (
		response, _ = json.Marshal(&ironManTwoGameResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockGame(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().UpdateGame(ctx, id, ironManTwoGameRequest).
		Return(nil)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		"/v1/games?id=b4606b93-15a2-4314-9ffd-e84c9b5fe8b8",
		strings.NewReader(`
		{
			"name": "Iron Man 2",
			"heroId":  "b4606b93-15a2-4314-9ffd-e84c9b5fe8b8",
			"universe": "MARVEL",
			"releaseYear": 2008
		}`),
	)
	h := &Handler{
		useCase: useCase,
	}
	h.putGame(ctx)
}

func TestHandlerDeleteGameSuccess(t *testing.T) {
	var (
		response, _ = json.Marshal(&ironManTwoGameResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockGame(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().Delete(ctx, id).
		Return(nil)
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		"/v1/games?id=b4606b93-15a2-4314-9ffd-e84c9b5fe8b8",
		strings.NewReader(``),
	)
	h := &Handler{
		useCase: useCase,
	}
	h.deleteGame(ctx)
}
