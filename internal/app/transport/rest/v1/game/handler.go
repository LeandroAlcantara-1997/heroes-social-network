package game

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase service.Game
}

func (h *Handler) postGame(ctx *gin.Context) {
	var req *dto.GameRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.AbortWithError(response.RestError(err))
		return
	}
	resp, err := h.useCase.Create(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) getGame(ctx *gin.Context) {
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}
	resp, err := h.useCase.GetByID(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) deleteGame(ctx *gin.Context) {
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}
	if err := h.useCase.Delete(ctx, id); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *Handler) putGame(ctx *gin.Context) {
	var (
		id, ok  = ctx.GetQuery("id")
		request dto.GameRequest
	)
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := request.Validator(); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	if err := h.useCase.UpdateGame(ctx, id, &request); err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.Status(http.StatusOK)
}
