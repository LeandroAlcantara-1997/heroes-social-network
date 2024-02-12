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

// @Summary      Create Team
// @Description  Create Team
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Param team body dto.TeamRequest true "team"
// @Success      201  {object}  dto.TeamResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /teams [post]
func (h *Handler) postTeam(ctx *gin.Context) {
	var request dto.TeamRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := request.Validator(); err != nil {
		ctx.AbortWithError(response.RestError(exception.ErrInvalidFields))
		return
	}

	resp, err := h.UseCase.CreateTeam(ctx, &request)
	if err != nil {
		ctx.JSON(response.RestError(err))
		return

	}

	ctx.JSON(http.StatusCreated, resp)
}

// @Summary      Get Team By ID
// @Description  Get Team By ID
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Param teamId formData string true "team"
// @Success      200  {object}  dto.TeamResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /teams [get]
func (h *Handler) getTeamByID(ctx *gin.Context) {
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

// @Summary      Delete Team
// @Description  Delete Team
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Param teamId formData string true "team"
// @Success      204
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /teams [delete]
func (h *Handler) deleteTeamByID(ctx *gin.Context) {
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

// @Summary      Get Team By Name
// @Description  Get Team By Name
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Param name path string true "team"
// @Success      200  {object}  dto.TeamResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /teams [get]
func (h *Handler) getTeamByName(ctx *gin.Context) {
	var request *dto.GetTeamByName
	if err := ctx.BindUri(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	resp, err := h.UseCase.GetTeamByName(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// @Summary      Update Team
// @Description  Update Team
// @Tags         Teams
// @Accept       json
// @Produce      json
// @Param teamId formData string true "team"
// @Param team body dto.TeamRequest true "team"
// @Success      200  {object}  dto.TeamResponse
// @Failure      400  {object}  error
// @Failure      404  {object}  error
// @Failure      500  {object}  error
// @Router       /teams [post]
func (h *Handler) updateTeam(ctx *gin.Context) {
	var request *dto.TeamRequest
	id, ok := ctx.GetQuery("id")
	if ok {
		if !validator.UUIDValidator(id) {
			ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidFields))
			return
		}

		if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		if err := h.UseCase.UpdateTeam(ctx, id, request); err != nil {
			ctx.AbortWithStatusJSON(response.RestError(err))
			return
		}
		ctx.Status(http.StatusOK)
	}
}
