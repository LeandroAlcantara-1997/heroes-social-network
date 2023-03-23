package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/dependency"
	v1 "github.com/LeandroAlcantara-1997/heroes-social-network/web/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	ctx := context.Background()
	r := gin.Default()

	dep, err := dependency.LoadDependency(ctx)
	if err != nil {
		panic(err)
	}

	v1.ConfigureRoutes(r, dep.HeroUseCase)
	log.Default().Printf("Server listening in :%s", config.Env.ApiPort)
	r.Run(fmt.Sprintf(":%s", config.Env.ApiPort))

}
