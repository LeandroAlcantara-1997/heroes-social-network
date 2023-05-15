package teams

import (
	"encoding/json"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	input "github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase input.Team
}

func (h *Handler) PostTeam(ctx *gin.Context) {
	var request *input.TeamRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, err := h.UseCase.RegisterTeam(ctx, request)
	if err != nil {
		code, err := exception.RestError(err)
		ctx.JSON(code, err)
		return

	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *Handler) GetTeamByID(ctx *gin.Context) {
	id, _ := ctx.GetQuery("id")
	response, err := h.UseCase.GetTeamByID(ctx, id)
	if err != nil {
		code, err := exception.RestError(err)
		ctx.JSON(code, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
