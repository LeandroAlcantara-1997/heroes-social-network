package console

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/config/env"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/middleware"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/console/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureConsoleRoutes(r *gin.Engine, consoleUseCase service.Console) {
	console := Handler{
		useCase: consoleUseCase,
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

	consoleRoute := r.Group("/v1/consoles").Use(m.Init)
	consoleRoute.POST("", console.postConsoles)
	consoleRoute.GET("", console.getConsoles)
}
