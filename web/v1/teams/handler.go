package teams

import (
	"encoding/json"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	customContext "github.com/LeandroAlcantara-1997/heroes-social-network/pkg/custom_context"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	input "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase input.Team
}

func (h *Handler) PostTeam(ctx *gin.Context) {
	var request *input.TeamRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := customContext.GetValidator(ctx).Struct(request); err != nil {
		code, err := exception.RestError(exception.ErrInvalidFields)
		ctx.AbortWithError(code, err)
		return
	}
	response, err := h.UseCase.RegisterTeam(ctx, request)
	if err != nil {
		code, err := exception.RestError(err)
		ctx.JSON(code, err)
		return

	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) GetTeamByID(ctx *gin.Context) {
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			code, err := exception.RestError(exception.ErrInvalidFields)
			ctx.AbortWithStatusJSON(code, err)
			return
		}
	}
	response, err := h.UseCase.GetTeamByID(ctx, id)
	if err != nil {
		code, err := exception.RestError(err)
		ctx.JSON(code, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteTeamByID(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			code, err := exception.RestError(exception.ErrInvalidFields)
			ctx.AbortWithStatusJSON(code, err)
			return
		}
		if err := h.UseCase.DeleteTeamByID(ctx, id); err != nil {
			code, err := exception.RestError(err)
			ctx.JSON(code, err)
			return
		}
	}
	ctx.AbortWithStatus(http.StatusOK)
}
