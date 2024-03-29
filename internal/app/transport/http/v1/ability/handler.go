package ability

import (
	"errors"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/http/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/dto"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

type Handler struct {
	useCase service.Ability
}

// @Summary      Create Ability
// @Description  Create Ability
// @Tags         Abilities
// @Accept       json
// @Produce      json
// @Param ability body dto.AbilityRequest true "ability"
// @Success      201  {object}  dto.AbilityResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /abilities [post]
func (h *Handler) postAbility(ctx *gin.Context) {
	c, span := otel.Tracer("ability").Start(ctx.Request.Context(), "postAbility")
	defer span.End()

	var req dto.AbilityRequest
	if err := ctx.BindJSON(&req); err != nil {
		log.GetLoggerFromContext(ctx).Error(ctx, err, nil)
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	if err := req.Validator(); err != nil {
		log.GetLoggerFromContext(ctx).Error(ctx, err, nil)
		ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidRequest))
		return
	}

	resp, err := h.useCase.CreateAbility(c, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @Summary      Get Ability By ID
// @Description  Get Ability By ID
// @Tags         Abilities
// @Accept       json
// @Produce      json
// @Param id query string true "ability"
// @Success      201  {object}  dto.AbilityResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /abilities [get]
func (h *Handler) getAbilityByID(ctx *gin.Context) {
	c, span := otel.Tracer("ability").Start(ctx.Request.Context(), "getAbilityByID")
	defer span.End()
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			log.GetLoggerFromContext(ctx).Error(ctx, errors.New("invalid uuid"), id)
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}

	resp, err := h.useCase.GetAbilityByID(c, id)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary      Get Ability By Hero ID
// @Description  Get Ability By Hero ID
// @Tags         Abilities
// @Accept       json
// @Produce      json
// @Param heroId query string true "heroId"
// @Success      200  {array}  dto.AbilityResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /abilities/heroes [get]
func (h *Handler) getAbilitiesByHeroID(ctx *gin.Context) {
	c, span := otel.Tracer("ability").Start(ctx.Request.Context(), "getAbilitiesByHeroID")
	defer span.End()
	var id, ok = ctx.GetQuery("heroId")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}

	resp, err := h.useCase.GetAbilitiesByHeroID(c, id)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary      Delete Ability By ID
// @Description  Delete Ability By ID
// @Tags         Abilities
// @Accept       json
// @Produce      json
// @Param id query string true "ability id"
// @Success      204
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /abilities [delete]
func (h *Handler) deleteAbility(ctx *gin.Context) {
	c, span := otel.Tracer("ability").Start(ctx.Request.Context(), "deleteAbility")
	defer span.End()
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}
	if err := h.useCase.DeleteAbility(c, id); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
