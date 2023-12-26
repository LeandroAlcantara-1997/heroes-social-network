package game

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/config/env"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/middleware"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/game/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureTeamRoutes(r *gin.Engine, gameUseCase service.Game) {
	game := Handler{
		useCase: gameUseCase,
	}
	m := &middleware.Middleware{
		Admin: false,
		Origin: middleware.Origin{
			Cors: &cors.Config{
				AllowOrigins:  util.ChunkTextByComma(env.Env.AllowOrigins),
				AllowMethods:  []string{http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodPut},
				AllowHeaders:  []string{"*"},
				ExposeHeaders: []string{"Content-Length", "content-type"},
			},
		},
	}

	gameRoute := r.Group("/v1/games").Use(m.Init)
	gameRoute.POST("", game.postGame)
}
