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

func (h *Handler) getConsoles(ctx *gin.Context) {
	resp, err := h.useCase.GetConsoles(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(response.RestError(err))
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}
