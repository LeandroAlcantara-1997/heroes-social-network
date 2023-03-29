package heroes

import (
	"encoding/json"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/exception"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UseCase input.Hero
}

func (h *Handler) PostHero(ctx *gin.Context) {
	var request *input.HeroRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.UseCase.RegisterHero(ctx, request)
	if err != nil {
		code, err := exception.RestError(err.Error())
		ctx.JSON(code, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) PutHero(ctx *gin.Context) {
	var (
		id, _   = ctx.GetQuery("id")
		request *input.HeroRequest
	)

	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.UseCase.UpdateHero(ctx, id, request)
	if err != nil {
		code, err := exception.RestError(err.Error())
		ctx.JSON(code, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetHeroByID(ctx *gin.Context) {
	var id, _ = ctx.GetQuery("id")

	response, err := h.UseCase.GetHeroByID(ctx, id)
	if err != nil {
		code, err := exception.RestError(err.Error())
		ctx.JSON(code, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) DeleteHeroByID(ctx *gin.Context) {
	var id, _ = ctx.GetQuery("id")

	if err := h.UseCase.DeleteHeroByID(ctx, id); err != nil {
		code, err := exception.RestError(err.Error())
		ctx.JSON(code, err)
		return
	}

	ctx.JSON(http.StatusAccepted, nil)
}
