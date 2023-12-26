package game

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/dto"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/service"
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
