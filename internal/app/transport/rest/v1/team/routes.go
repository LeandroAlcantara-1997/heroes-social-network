package team

import (
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/config/env"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/middleware"
	service "github.com/LeandroAlcantara-1997/heroes-social-network/internal/domain/team/service"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConfigureTeamRoutes(r *gin.Engine, teamUseCase service.Team) {
	team := Handler{
		UseCase: teamUseCase,
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

	teamsRoute := r.Group("/v1/teams").Use(m.Init)
	teamsRoute.POST("", team.PostTeam)
	teamsRoute.PUT("", team.UpdateTeam)
	teamsRoute.GET("", team.GetTeamByID)
	teamsRoute.GET(":name", team.GetTeamByName)
	teamsRoute.DELETE("", team.DeleteTeamByID)

}
