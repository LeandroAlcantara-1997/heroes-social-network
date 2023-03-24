package heroes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase input.Hero
}

func (h *Handler) PostHero(ctx *gin.Context) {
	var dto *input.HeroRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	dto.Universe = strings.ToUpper(dto.Universe)

	dtoResponse, err := h.UseCase.RegisterHero(ctx, dto)
	if err != nil {
		code, err := exception.RestError(err.Error())
		ctx.JSON(code, err)
		return
	}

	ctx.JSON(http.StatusOK, dtoResponse)

}
