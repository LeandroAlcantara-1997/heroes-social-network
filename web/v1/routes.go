package v1

import (
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/hero"
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input/team"
	"github.com/LeandroAlcantara-1997/heroes-social-network/web/v1/heroes"
	"github.com/LeandroAlcantara-1997/heroes-social-network/web/v1/teams"
	"github.com/gin-gonic/gin"
)

func ConfigureHeroRoutes(r *gin.Engine, heroUseCase hero.Hero) {
	hero := heroes.Handler{
		UseCase: heroUseCase,
	}

	heroesRoute := r.Group("/v1/heroes")
	heroesRoute.POST("", hero.PostHero)
	heroesRoute.PUT("", hero.PutHero)
	heroesRoute.GET("", hero.GetHeroByID)
	heroesRoute.DELETE("", hero.DeleteHeroByID)
}

func ConfigureTeamRoutes(r *gin.Engine, teamUseCase team.Team) {
	team := teams.Handler{
		UseCase: teamUseCase,
	}

	teamsRoute := r.Group("/v1/teams")
	teamsRoute.POST("", team.PostTeam)
	teamsRoute.GET("", team.GetTeamByID)

}
