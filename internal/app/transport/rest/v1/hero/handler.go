package hero

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/dto"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

type Handler struct {
	UseCase service.Hero
}

// @Summary      Create Hero
// @Description  Create hero
// @Tags         Heroes
// @Accept       json
// @Produce      json
// @Param hero body dto.HeroRequest true "hero"
// @Success      200  {object}  dto.HeroResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /heroes [post]
func (h *Handler) postHero(ctx *gin.Context) {
	c, span := otel.Tracer("hero").Start(ctx.Request.Context(), "postHero")
	defer span.End()
	var request dto.HeroRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := request.Validator(); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidRequest))
		return
	}

	resp, err := h.UseCase.CreateHero(c, &request)
	if err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @Summary      Update Hero
// @Description  Update hero
// @Tags         Heroes
// @Accept       json
// @Produce      json
// @Param id query string true "hero id"
// @Param hero body dto.HeroRequest true "body hero"
// @Success      200  {object}  dto.HeroResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /heroes [put]
func (h *Handler) putHero(ctx *gin.Context) {
	c, span := otel.Tracer("hero").Start(ctx.Request.Context(), "putHero")
	defer span.End()
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

	if err := h.UseCase.UpdateHero(c, id, &request); err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary      Get Hero By ID
// @Description  Get Hero By ID
// @Tags         Heroes
// @Accept       json
// @Produce      json
// @Param id query string true "hero id"
// @Success      200  {object}  dto.HeroResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /heroes [get]
func (h *Handler) getHeroByID(ctx *gin.Context) {
	c, span := otel.Tracer("hero").Start(ctx.Request.Context(), "getHeroByID")
	defer span.End()
	var id, ok = ctx.GetQuery("id")
	if ok && !validator.UUIDValidator(id) {
		ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
		return
	}

	resp, err := h.UseCase.GetHeroByID(c, id)
	if err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary      Delete Hero By ID
// @Description  Delete Hero By ID
// @Tags         Heroes
// @Accept       json
// @Produce      json
// @Param id query string true "hero id"
// @Success      204
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /heroes [delete]
func (h *Handler) deleteHeroByID(ctx *gin.Context) {
	c, span := otel.Tracer("hero").Start(ctx.Request.Context(), "deleteHeroByID")
	defer span.End()
	if id, ok := ctx.GetQuery("id"); ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}

		if err := h.UseCase.DeleteHeroByID(c, id); err != nil {
			ctx.JSON(response.RestError(err))
			return
		}
	}

	ctx.Status(http.StatusNoContent)
}

// @Summary      Add Ability for Hero
// @Description  Add Ability for Hero
// @Tags         Heroes
// @Accept       json
// @Produce      json
// @Param ability query string true "ability id"
// @Param hero query string true "hero id"
// @Success      201
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /heroes/abilities [post]
func (h *Handler) postAddAbilityToHero(ctx *gin.Context) {
	c, span := otel.Tracer("hero").Start(ctx.Request.Context(), "postAddAbilityToHero")
	defer span.End()
	var request dto.AddAbilityToHeroRequest

	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.UseCase.AddAbilityToHero(c, request.AbilityID, request.HeroID); err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.Status(http.StatusCreated)
}
