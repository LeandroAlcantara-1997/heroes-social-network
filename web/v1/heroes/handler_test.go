package heroes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/mock"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

var (
	batmanRequest = &input.HeroRequest{
		HeroName:  "Batman",
		CivilName: "Bruce Wayne",
		Hero:      true,
		Universe:  "DC",
		Team:      nil,
	}
	batmanResponse = &input.HeroResponse{
		Id:        "4d67708f-f5fc-49c5-8ed3-90e5e078917c",
		HeroName:  "Batman",
		CivilName: "Bruce Wayne",
		Hero:      true,
		Universe:  "DC",
		Team:      nil,
		CreatedAt: time.Now().UTC(),
	}
	batmanRequestWithInvalidField = &input.HeroRequest{
		HeroName:  "Batman",
		CivilName: "Bruce Wayne",
		Hero:      true,
		Universe:  "DCI",
		Team:      nil,
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
	useCase.EXPECT().RegisterHero(ctx, batmanRequest).Return(batmanResponse, nil)
	h := Handler{
		useCase,
	}
	h.PostHero(ctx)
}

func TestHandlerPostHeroFailInvalidField(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrInvalidFields.Error()))
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
	useCase.EXPECT().RegisterHero(ctx, batmanRequestWithInvalidField).Return(nil, exception.ErrInvalidFields)
	h := Handler{
		useCase,
	}
	h.PostHero(ctx)
}

func TestHandlerPostHeroFailInternalServerError(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrInternalServer.Error()))
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
	useCase.EXPECT().RegisterHero(ctx, batmanRequest).Return(nil, exception.ErrInternalServer)
	h := Handler{
		useCase,
	}
	h.PostHero(ctx)
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
	useCase.EXPECT().UpdateHero(ctx, batmanResponse.Id, batmanRequest).
		Return(batmanResponse, nil)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
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
	h.PutHero(ctx)
}

func TestHandlerPutHeroFailInvalidField(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrInvalidFields.Error()))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusBadRequest,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().UpdateHero(ctx, batmanResponse.Id, batmanRequestWithInvalidField).
		Return(nil, exception.ErrInvalidFields)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
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
	h.PutHero(ctx)
}

func TestHandlerPutHeroFailInternalServerError(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrInternalServer.Error()))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().UpdateHero(ctx, batmanResponse.Id, batmanRequest).
		Return(nil, exception.ErrInternalServer)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
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
	h.PutHero(ctx)
}

// GET
func TestHandlerGetHeroByIDSucess(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrInternalServer.Error()))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().GetHeroByID(ctx, batmanResponse.Id).
		Return(nil, exception.ErrInternalServer)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.GetHeroByID(ctx)
}

func TestHandlerGetHeroByIDFailHeroNotFound(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrHeroNotFound.Error()))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().GetHeroByID(ctx, batmanResponse.Id).
		Return(nil, exception.ErrHeroNotFound)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.GetHeroByID(ctx)
}

func TestHandlerGetHeroByIDFailHInternalServerError(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrInternalServer.Error()))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().GetHeroByID(ctx, batmanResponse.Id).
		Return(nil, exception.ErrInternalServer)
	ctx.Request = httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.GetHeroByID(ctx)
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
	useCase.EXPECT().DeleteHeroByID(ctx, batmanResponse.Id).
		Return(nil)
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.DeleteHeroByID(ctx)
}

func TestHandlerDeleteHeroByIDFailHeroNotFound(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrHeroNotFound.Error()))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().DeleteHeroByID(ctx, batmanResponse.Id).
		Return(exception.ErrHeroNotFound)
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.DeleteHeroByID(ctx)
}

func TestHandlerDeleteHeroByIDFailInternalServerError(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.New(exception.ErrInternalServer.Error()))
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusInternalServerError,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockHero(ctrl)
	)
	defer ctrl.Finish()
	useCase.EXPECT().DeleteHeroByID(ctx, batmanResponse.Id).
		Return(exception.ErrInternalServer)
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/v1/heroes?id=%s", batmanResponse.Id),
		strings.NewReader(`{}`),
	)
	h := &Handler{
		UseCase: useCase,
	}
	h.DeleteHeroByID(ctx)
}
