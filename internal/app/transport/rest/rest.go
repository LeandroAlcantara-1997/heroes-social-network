package rest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/container"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/v1/hero"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/v1/team"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type api struct {
	allowOrigins string
	port         string
	container    *container.Container
}

func New(port, allowOrigins string, container *container.Container) *api {
	return &api{
		port:         port,
		allowOrigins: allowOrigins,
		container:    container,
	}
}

func (a *api) NewServer(ctx context.Context) {
	r := gin.Default()

	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:  util.ChunkTextByComma(a.allowOrigins),
		AllowMethods:  []string{http.MethodGet},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length", "content-type"},
	}))

	hero.ConfigureHeroRoutes(r, a.container.HeroUseCase)
	team.ConfigureTeamRoutes(r, a.container.TeamUseCase)

	log.Default().Printf("Server listening in :%s", a.port)

	a.shutdown(&http.Server{
		Addr:    fmt.Sprintf(":%s", a.port),
		Handler: r,
	})
}

func (a *api) shutdown(server *http.Server) {
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}
