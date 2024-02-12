package console

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/dto"
	console "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase console.Console
}

// @Summary      Create Console
// @Description  Create new Console
// @Tags         Consoles
// @Accept       json
// @Produce      json
// @Param console body dto.ConsoleRequest true "consoles"
// @Success      201  {object} dto.ConsoleResponse
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /consoles [post]
func (h *Handler) postConsoles(ctx *gin.Context) {
	var req dto.ConsoleRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	if err := req.Validator(); err != nil {
		ctx.AbortWithStatusJSON(response.RestError(exception.ErrInvalidRequest))
		return
	}

	resp, err := h.useCase.CreateConsoles(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

// @Summary      Get All Consoles
// @Description  Get All Consoles
// @Tags         Consoles
// @Accept       json
// @Produce      json
// @Success      200  {object} dto.ConsoleResponse
// @Failure      400  {object}  error
// @Failure      500  {object}  error
// @Router       /consoles [get]
func (h *Handler) getConsoles(ctx *gin.Context) {
	resp, err := h.useCase.GetConsoles(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
