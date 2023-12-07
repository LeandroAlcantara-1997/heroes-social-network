package team

import (
	"encoding/json"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/dto"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase service.Team
}

func (h *Handler) PostTeam(ctx *gin.Context) {
	var request dto.TeamRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := request.Validator(); err != nil {
		ctx.AbortWithError(response.RestError(exception.ErrInvalidFields))
		return
	}

	resp, err := h.UseCase.RegisterTeam(ctx, &request)
	if err != nil {
		ctx.JSON(response.RestError(err))
		return

	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetTeamByID(ctx *gin.Context) {
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}
	resp, err := h.UseCase.GetTeamByID(ctx, id)
	if err != nil {
		ctx.JSON(response.RestError(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteTeamByID(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
		if err := h.UseCase.DeleteTeamByID(ctx, id); err != nil {
			ctx.JSON(response.RestError(err))
			return
		}
	}
	ctx.AbortWithStatus(http.StatusOK)
}

func (h *Handler) GetTeamByName(ctx *gin.Context) {
	var request *dto.GetTeamByName
	if err := ctx.BindUri(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resp, err := h.UseCase.GetTeamByName(ctx, request)
	if err != nil {
		ctx.JSON(response.RestError(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateTeam(ctx *gin.Context) {
	var request *dto.TeamRequest
	id, ok := ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}

		if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := h.UseCase.UpdateTeam(ctx, id, request); err != nil {
			ctx.AbortWithError(response.RestError(err))
			return
		}
		ctx.Status(http.StatusOK)
	}
}
