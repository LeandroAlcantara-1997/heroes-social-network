package ability

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/dto"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase service.Ability
}

func (h *Handler) postAbility(ctx *gin.Context) {
	var req dto.AbilityRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	if err := req.Validator(); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidRequest))
		return
	}

	resp, err := h.useCase.CreateAbility(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) getAbilityByID(ctx *gin.Context) {
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}

	resp, err := h.useCase.GetAbilityByID(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) getAbilitiesByHeroID(ctx *gin.Context) {
	var id, ok = ctx.GetQuery("heroId")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}

	resp, err := h.useCase.GetAbilitiesByHeroID(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
