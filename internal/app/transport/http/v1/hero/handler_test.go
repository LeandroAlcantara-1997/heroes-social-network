package hero

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/http/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

var (
	batmanRequest = &dto.HeroRequest{
		HeroName:  "Batman",
		CivilName: "Bruce Wayne",
		Hero:      true,
		Universe:  "DC",
		Team:      nil,
	}
	batmanResponse = &dto.HeroResponse{
		ID:        "4d67708f-f5fc-49c5-8ed3-90e5e078917c",
		HeroName:  "Batman",
		CivilName: "Bruce Wayne",
		Hero:      true,
		Universe:  "DC",
		Team:      nil,
		CreatedAt: time.Now().UTC(),
	}
)

// Post Hero
func TestHandlerPostHeroSuccess(t *testing.T) {
	var (
		response, _ = json.Marshal(&batmanResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/heroes",
		strings.NewReader(`{
					"heroName": "Batman",
					"civilName": "Bruce Wayne",
					"hero":  true,
					"universe": "DC"
				}`),
	)
	useCase.EXPECT().CreateHero(gomock.Any(), batmanRequest).Return(batmanResponse, nil)
	h := Handler{
		useCase,
	}
	h.postHero(ctx)
}

func TestHandlerPostHeroFailInvalidField(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrInvalidFields))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusBadRequest,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/heroes",
		strings.NewReader(`{
					"heroName": "Batman",
					"civilName": "Bruce Wayne",
					"hero":  true,
					"universe": "DCI"
				}`),
	)
	h := Handler{
		useCase,
	}
	h.postHero(ctx)
}

func TestHandlerPostHeroFailInternalServerError(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrInternalServer))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/heroes",
		strings.NewReader(`{
					"heroName": "Batman",
					"civilName": "Bruce Wayne",
					"hero":  true,
					"universe": "DC"
				}`),
	)
	useCase.EXPECT().CreateHero(gomock.Any(), batmanRequest).Return(nil, exception.ErrInternalServer)
	h := Handler{
		useCase,
	}
	h.postHero(ctx)
}

// PUT Hero
func TestHandlerPutHeroSuccess(t *testing.T) {
	var (
		response, _ = json.Marshal(&batmanRequest)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().UpdateHero(gomock.Any(), batmanResponse.ID, batmanRequest).
		Return(nil)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{
					"heroName": "Batman",
					"civilName": "Bruce Wayne",
					"hero":  true,
					"universe": "DC"
				}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.putHero(ctx)
}

func TestHandlerPutHeroFailInvalidField(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrInvalidFields))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusBadRequest,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{
					"heroName": "Batman",
					"civilName": "Bruce Wayne",
					"hero":  true,
					"universe": "DCI"
				}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.putHero(ctx)
}

func TestHandlerPutHeroFailInternalServerError(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrInternalServer))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().UpdateHero(gomock.Any(), batmanResponse.ID, batmanRequest).
		Return(exception.ErrInternalServer)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{
					"heroName": "Batman",
					"civilName": "Bruce Wayne",
					"hero":  true,
					"universe": "DC"
				}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.putHero(ctx)
}

// GET
func TestHandlerGetHeroByIDSucess(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrInternalServer))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().GetHeroByID(gomock.Any(), batmanResponse.ID).
		Return(nil, exception.ErrInternalServer)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.getHeroByID(ctx)
}

func TestHandlerGetHeroByIDFailHeroNotFound(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrHeroNotFound))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().GetHeroByID(gomock.Any(), batmanResponse.ID).
		Return(nil, exception.ErrHeroNotFound)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.getHeroByID(ctx)
}

func TestHandlerGetHeroByIDFailHInternalServerError(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrInternalServer))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().GetHeroByID(gomock.Any(), batmanResponse.ID).
		Return(nil, exception.ErrInternalServer)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.getHeroByID(ctx)
}

// Delete
func TestHandlerDeleteHeroByIDSuccess(t *testing.T) {
	var (
		response, _ = json.Marshal(nil)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().DeleteHeroByID(gomock.Any(), batmanResponse.ID).
		Return(nil)
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.deleteHeroByID(ctx)
}

func TestHandlerDeleteHeroByIDFailHeroNotFound(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrHeroNotFound))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().DeleteHeroByID(gomock.Any(), batmanResponse.ID).
		Return(exception.ErrHeroNotFound)
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.deleteHeroByID(ctx)
}

func TestHandlerDeleteHeroByIDFailInternalServerError(t *testing.T) {
	var (
		response, _ = json.Marshal(response.New(exception.ErrInternalServer))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().DeleteHeroByID(gomock.Any(), batmanResponse.ID).
		Return(exception.ErrInternalServer)
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.ID),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.deleteHeroByID(ctx)
}
