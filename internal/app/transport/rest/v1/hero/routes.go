package hero

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/config/env"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/middleware"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/hero/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureHeroRoutes(r *gin.Engine, heroUseCase service.Hero) {
	hero := Handler{
		UseCase: heroUseCase,
	}

	m := &middleware.Middleware{
		Admin: false,
		Origin: middleware.Origin{
			Cors: &cors.Config{
				AllowOrigins:  util.ChunkTextByComma(env.Env.AllowOrigins),
				AllowMethods:  []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete},
				AllowHeaders:  []string{"*"},
				ExposeHeaders: []string{"Content-Length", "content-type"},
			},
		},
	}

	heroesRoute := r.Group("/v1/heroes").Use(m.Init)
	heroesRoute.POST("", hero.PostHero)
	heroesRoute.PUT("", hero.PutHero)
	heroesRoute.GET("", hero.GetHeroByID)
	heroesRoute.DELETE("", hero.DeleteHeroByID)
}
