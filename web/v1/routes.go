package v1

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/validator"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/hero"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	"github.com/LeandroAlcantara-1997/heroes-social-network/web/middleware"
	"github.com/LeandroAlcantara-1997/heroes-social-network/web/v1/heroes"
	"github.com/LeandroAlcantara-1997/heroes-social-network/web/v1/teams"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureHeroRoutes(r *gin.Engine, heroUseCase hero.Hero) {
	hero := heroes.Handler{
		UseCase: heroUseCase,
	}

	m := &middleware.Middleware{
		Admin: false,
		Validator: validator.RegisterValidateFunc([]validator.CustomValidator{
			{
				TagName:    "universe",
				CustomFunc: validator.CheckUniverse,
			},
		}),
		Origin: middleware.Origin{
			Cors: &cors.Config{
				AllowOrigins:  util.ChunkTextByComma(config.Env.AllowOrigins),
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

func ConfigureTeamRoutes(r *gin.Engine, teamUseCase team.Team) {
	team := teams.Handler{
		UseCase: teamUseCase,
	}
	m := &middleware.Middleware{
		Admin: false,
		Validator: validator.RegisterValidateFunc([]validator.CustomValidator{
			{
				TagName:    "universe",
				CustomFunc: validator.CheckUniverse,
			},
		}),
		Origin: middleware.Origin{
			Cors: &cors.Config{
				AllowOrigins:  util.ChunkTextByComma(config.Env.AllowOrigins),
				AllowMethods:  []string{http.MethodPost, http.MethodGet, http.MethodDelete},
				AllowHeaders:  []string{"*"},
				ExposeHeaders: []string{"Content-Length", "content-type"},
			},
		},
	}

	teamsRoute := r.Group("/v1/teams").Use(m.Init)
	teamsRoute.POST("", team.PostTeam)
	teamsRoute.GET("", team.GetTeamByID)
	teamsRoute.DELETE("", team.DeleteTeamByID)

}
