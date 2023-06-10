package heroes

import (
	"encoding/json"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	customContext "github.com/LeandroAlcantara-1997/heroes-social-network/pkg/custom_context"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	input "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/hero"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase input.Hero
}

func (h *Handler) PostHero(ctx *gin.Context) {
	var request *input.HeroRequest

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := customContext.GetValidator(ctx).Struct(request); err != nil {
		code, err := exception.RestError(exception.ErrInvalidFields)
		ctx.AbortWithError(code, err)
		return
	}

	response, err := h.UseCase.RegisterHero(ctx, request)
	if err != nil {
		code, err := exception.RestError(err)
		ctx.JSON(code, err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) PutHero(ctx *gin.Context) {
	var (
		id, ok  = ctx.GetQuery("id")
		request *input.HeroRequest
	)
	if ok {
		if !validator.UUIDValidator(id) {
			code, err := exception.RestError(exception.ErrInvalidFields)
			ctx.AbortWithStatusJSON(code, err)
			return
		}
	}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := customContext.GetValidator(ctx).Struct(request); err != nil {
		code, err := exception.RestError(exception.ErrInvalidFields)
		ctx.AbortWithError(code, err)
		return
	}
	response, err := h.UseCase.UpdateHero(ctx, id, request)
	if err != nil {
		code, err := exception.RestError(err)
		ctx.JSON(code, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetHeroByID(ctx *gin.Context) {
	var id, ok = ctx.GetQuery("id")
	if ok && !validator.UUIDValidator(id) {
		code, err := exception.RestError(exception.ErrInvalidFields)
		ctx.AbortWithStatusJSON(code, err)
		return
	}

	response, err := h.UseCase.GetHeroByID(ctx, id)
	if err != nil {
		code, err := exception.RestError(err)
		ctx.JSON(code, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteHeroByID(ctx *gin.Context) {
	if id, ok := ctx.GetQuery("id"); ok {
		if !validator.UUIDValidator(id) {
			code, err := exception.RestError(exception.ErrInvalidFields)
			ctx.AbortWithStatusJSON(code, err)
			return
		}
		if err := h.UseCase.DeleteHeroByID(ctx, id); err != nil {
			code, err := exception.RestError(err)
			ctx.JSON(code, err)
			return
		}
	}

	ctx.Status(http.StatusAccepted)
}
