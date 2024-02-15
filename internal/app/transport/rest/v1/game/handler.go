package game

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

type Handler struct {
	useCase service.Game
}

// @Summary      Create Game
// @Description  Create Game
// @Tags         Games
// @Accept       json
// @Produce      json
// @Param game body dto.GameRequest true "game"
// @Success      201  {object}  dto.GameResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /games [post]
func (h *Handler) postGame(ctx *gin.Context) {
	c, span := otel.Tracer("game").Start(ctx.Request.Context(), "postGame")
	defer span.End()
	var req dto.GameRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	if err := req.Validator(); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidRequest))
		return
	}

	resp, err := h.useCase.CreateGame(c, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @Summary      Get Game By ID
// @Description  Get Game BY ID
// @Tags         Games
// @Accept       json
// @Produce      json
// @Param id query string true "game id"
// @Success      201  {object}  dto.GameResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /games [get]
func (h *Handler) getGame(ctx *gin.Context) {
	c, span := otel.Tracer("game").Start(ctx.Request.Context(), "getGame")
	defer span.End()
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}
	resp, err := h.useCase.GetByID(c, id)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// @Summary      Delete Game By ID
// @Description  Delete Game BY ID
// @Tags         Games
// @Accept       json
// @Produce      json
// @Param id query string true "game id"
// @Success      204
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /games [delete]
func (h *Handler) deleteGame(ctx *gin.Context) {
	c, span := otel.Tracer("game").Start(ctx.Request.Context(), "deleteGame")
	defer span.End()
	var id, ok = ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}
	}
	if err := h.useCase.Delete(c, id); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// @Summary      Update Game
// @Description  Update Game
// @Tags         Games
// @Accept       json
// @Produce      json
// @Param id query string true "game id"
// @Param game body dto.GameRequest true "body game"
// @Success      200
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /games [put]
func (h *Handler) putGame(ctx *gin.Context) {
	c, span := otel.Tracer("game").Start(ctx.Request.Context(), "putGame")
	defer span.End()
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

	if err := h.useCase.UpdateGame(c, id, &request); err != nil {
		ctx.JSON(response.RestError(err))
		return
	}

	ctx.Status(http.StatusOK)
}
