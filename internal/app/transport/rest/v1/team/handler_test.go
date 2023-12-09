package team

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/dto"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

var (
	xMenResponse = &dto.TeamResponse{
		ID:        uuid.NewString(),
		Name:      "X-Men",
		Universe:  "MARVEL",
		CreatedAt: time.Now().UTC(),
	}
	xMenRequest = &dto.TeamRequest{
		Name:     "X-Men",
		Universe: "MARVEL",
	}
)

func TestHandler_PostTeam(t *testing.T) {
	var (
		response, _ = json.Marshal(&xMenResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockTeam(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/teams",
		strings.NewReader(`{
					"name": "X-Men",
					"universe": "MARVEL"
				}`),
	)
	useCase.EXPECT().RegisterTeam(ctx, xMenRequest).Return(xMenResponse, nil)
	h := Handler{
		useCase,
	}
	h.PostTeam(ctx)
}

func TestHandlerPostTeamFailInvaliidUniverse(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.ErrInvalidFields.Error())
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusOK,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockTeam(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodPost,
		"/v1/teams",
		strings.NewReader(`{
					"name": "X-Men",
					"universe": "MARVE"
				}`),
	)
	h := Handler{
		useCase,
	}
	h.PostTeam(ctx)
}

func TestHandlerGetTeamByIDSuccess(t *testing.T) {
	var (
		response, _ = json.Marshal(&xMenResponse)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusFound,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockTeam(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodGet,
		"/v1/teamss?id=e36b3582-f936-47b7-8832-47da045ea4e9",
		nil,
	)
	useCase.EXPECT().GetTeamByID(ctx, "e36b3582-f936-47b7-8832-47da045ea4e9").Return(xMenResponse, nil)
	h := Handler{
		useCase,
	}
	h.GetTeamByID(ctx)
}

func TestHandlerGetTeamByIDFailTeamNotFound(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.ErrTeamNotFound)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusFound,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockTeam(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodGet,
		"/v1/teamss?id=e36b3582-f936-47b7-8832-47da045ea4e9",
		nil,
	)
	useCase.EXPECT().GetTeamByID(ctx, "e36b3582-f936-47b7-8832-47da045ea4e9").Return(nil, exception.ErrTeamNotFound)
	h := Handler{
		useCase,
	}
	h.GetTeamByID(ctx)
}

func TestHandlerDeleteTeamByIDSuccess(t *testing.T) {
	var (
		ctx, _ = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusFound,
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockTeam(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		"/v1/teamss?id=e36b3582-f936-47b7-8832-47da045ea4e9",
		nil,
	)
	useCase.EXPECT().DeleteTeamByID(ctx, "e36b3582-f936-47b7-8832-47da045ea4e9").Return(nil)
	h := Handler{
		useCase,
	}

	h.DeleteTeamByID(ctx)
}

func TestHandlerDeleteTeamByIDFail(t *testing.T) {
	var (
		response, _ = json.Marshal(exception.ErrTeamNotFound)
		ctx, _      = gin.CreateTestContext(&httptest.ResponseRecorder{
			Code: http.StatusFound,
			Body: bytes.NewBuffer(response),
		})
		ctrl    = gomock.NewController(t)
		useCase = mock.NewMockTeam(ctrl)
	)
	defer ctrl.Finish()
	ctx.Request = httptest.NewRequest(
		http.MethodDelete,
		"/v1/teamss?id=e36b3582-f936-47b7-8832-47da045ea4e9",
		nil,
	)
	useCase.EXPECT().DeleteTeamByID(ctx, "e36b3582-f936-47b7-8832-47da045ea4e9").Return(nil)
	h := Handler{
		useCase,
	}

	h.DeleteTeamByID(ctx)
}
