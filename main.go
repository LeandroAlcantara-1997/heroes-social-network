package main

import (
	_ "github.com/LeandroAlcantara-1997/heroes-social-network/docs"

	"github.com/LeandroAlcantara-1997/heroes-social-network/config/env"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/container"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest"
)

// @title           Heroes Social Network
// @version         1.0
// @description     Heroes social network is a project created to make life easier for superhero fans.
// @termsOfService  http://swagger.io/terms/

// @contact.url    https://www.linkedin.com/in/leandro-alcantara-pro

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	ctx, cont, err := container.New()
	if err != nil {
		panic(err)
	}
	rest.New(env.Env.APIPort, env.Env.AllowOrigins, cont).NewServer(ctx)
}
