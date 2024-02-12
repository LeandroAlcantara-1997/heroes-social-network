package ability

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

const id = "e36b3582-f936-47b7-8832-47da045ea4e9"

var (
	abilityResponse = &dto.AbilityResponse{
		ID:          "1234",
		Description: "fly",
		CreatedAt:   time.Now().UTC(),
	}

	abilityRequest = &dto.AbilityRequest{
		Description: "fly",
	}
)

func TestHandlerPostAbility(t *testing.T) {
	var (
		response, _ = json.Marshal(&abilityResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusCreated,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockAbility(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/abilities",
		strings.NewReader(`{
					"description": "fly"
				}`),
	)
	useCase.EXPECT().CreateAbility(ctx, abilityRequest).Return(abilityResponse, nil)
	h := Handler{
		useCase,
	}
	h.postAbility(ctx)
}

func TestHandlerGetAbilityByID(t *testing.T) {
	var (
		response, _ = json.Marshal(&abilityResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockAbility(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodGet,
		"/v1/abilities?id="+id,
		nil,
	)
	useCase.EXPECT().GetAbilityByID(ctx, id).Return(abilityResponse, nil)
	h := Handler{
		useCase,
	}
	h.getAbilityByID(ctx)
}

func TestHandlerGetAbilityByHeroID(t *testing.T) {
	var (
		response, _ = json.Marshal(&abilityResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockAbility(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodGet,
		"/v1/abilities?heroId="+id,
		nil,
	)
	useCase.EXPECT().GetAbilitiesByHeroID(ctx, id).Return([]dto.AbilityResponse{
		*abilityResponse,
	}, nil)
	h := Handler{
		useCase,
	}
	h.getAbilitiesByHeroID(ctx)
}

func TestHandlerDeleteAbility(t *testing.T) {
	var (
		ctx, _ = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: nil,
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockAbility(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodGet,
		"/v1/abilities?id="+id,
		nil,
	)
	useCase.EXPECT().DeleteAbility(ctx, id).Return(nil)
	h := Handler{
		useCase,
	}
	h.deleteAbility(ctx)
}
