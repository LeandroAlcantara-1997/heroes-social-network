package main

import (
	"fmt"
	"log"

	"github.com/LeandroAlcantara-1997/heroes-social-network/domain/heroes"
	"github.com/LeandroAlcantara-1997/heroes-social-network/infrastructure/config"
	v1 "github.com/LeandroAlcantara-1997/heroes-social-network/web/v1"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	r := gin.Default()

	heroService := heroes.New()
	v1.ConfigureRoutes(r, heroService)
	log.Default().Printf("Server listening in :%s", config.Env.ApiPort)
	r.Run(fmt.Sprintf(":%s", config.Env.ApiPort))

}
