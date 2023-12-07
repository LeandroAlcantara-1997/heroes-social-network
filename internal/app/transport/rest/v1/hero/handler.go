package hero

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/dto"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase service.Hero
}

func (h *Handler) PostHero(ctx *gin.Context) {
	var request dto.HeroRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := request.Validator(); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidRequest))
		return
	}

	resp, err := h.UseCase.RegisterHero(ctx, &request)
	if err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) PutHero(ctx *gin.Context) {
	var (
		id, ok  = ctx.GetQuery("id")
		request dto.HeroRequest
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

	if err := h.UseCase.UpdateHero(ctx, id, &request); err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) GetHeroByID(ctx *gin.Context) {
	var id, ok = ctx.GetQuery("id")
	if ok && !validator.UUIDValidator(id) {
		ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
		return
	}

	resp, err := h.UseCase.GetHeroByID(ctx, id)
	if err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteHeroByID(ctx *gin.Context) {
	if id, ok := ctx.GetQuery("id"); ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}

		if err := h.UseCase.DeleteHeroByID(ctx, id); err != nil {
			ctx.JSON(response.RestError(err))
			return
		}
	}

	ctx.Status(http.StatusAccepted)
}
