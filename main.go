package main

import (
	"github.com/LeandroAlcantara-1997/heroes-social-network/config/env"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/container"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest"
)

func main() {
	ctx, cont, err := container.New()
	if err != nil {
		panic(err)
	}
	rest.New(env.Env.APIPort, env.Env.AllowOrigins, cont).NewServer(ctx)
}
