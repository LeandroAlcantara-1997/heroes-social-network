package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/dependency"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	v1 "github.com/LeandroAlcantara-1997/heroes-social-network/web/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	ctx := context.Background()
	r := gin.Default()

	configureGlobalMiddleware(ctx, r)

	log.Default().Printf("Server listening in :%s", config.Env.ApiPort)
	r.Run(fmt.Sprintf(":%s", config.Env.ApiPort))
}

func configureGlobalMiddleware(ctx context.Context, r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:  util.ChunkTextByComma(config.Env.AllowOrigins),
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length", "content-type"},
		// AllowAllOrigins: true,
	}))

	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	dep, err := dependency.LoadDependency(ctx)
	if err != nil {
		panic(err)
	}

	v1.ConfigureHeroRoutes(r, dep.HeroUseCase)
	v1.ConfigureTeamRoutes(r, dep.TeamUseCase)
}
