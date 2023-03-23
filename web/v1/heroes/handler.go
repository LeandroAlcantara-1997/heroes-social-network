package heroes

import (
	"encoding/json"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase input.Hero
}

func (h *Handler) PostHero(ctx *gin.Context) {
	var dto *input.HeroRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	dtoResponse, err := h.UseCase.RegisterHero(ctx, *dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, dtoResponse)

}
