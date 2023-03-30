package v1

import (
	"github.com/LeandroAlcantara-1997/heroes-social-network/ports/input"
	"github.com/LeandroAlcantara-1997/heroes-social-network/web/v1/heroes"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, heroUseCase input.Hero) {
	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})

	hero := heroes.Handler{
		UseCase: heroUseCase,
	}

	heroesRoute := r.Group("/v1/heroes")
	heroesRoute.POST("", hero.PostHero)
	heroesRoute.PUT("", hero.PutHero)
	heroesRoute.GET("", hero.GetHeroByID)
	heroesRoute.DELETE("", hero.DeleteHeroByID)
}
