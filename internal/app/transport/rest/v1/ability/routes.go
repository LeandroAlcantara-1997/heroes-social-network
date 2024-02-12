package ability

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/config/env"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/middleware"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/ability/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureGameRoutes(r *gin.Engine, abilityUseCase service.Ability) {
	ability := Handler{
		useCase: abilityUseCase,
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

	abilityRoute := r.Group("/v1/abilities").Use(m.Init)
	abilityRoute.POST("", ability.postAbility)
	abilityRoute.GET("", ability.getAbilityByID)
	abilityRoute.GET("/heroes", ability.getAbilitiesByHeroID)
	abilityRoute.DELETE("", ability.deleteAbility)
}
