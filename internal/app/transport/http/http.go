package http

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

	docs "github.com/LeandroAlcantara-1997/heroes-social-network/docs"
	logger "github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/container"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/http/v1/ability"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/http/v1/console"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/http/v1/game"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/http/v1/hero"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/http/v1/team"
	"github.com/LeandroAlcantara-1997/heroes-social-network/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type api struct {
	allowOrigins string
	port         string
	version      string
	apiName      string
	environment  string
	container    *container.Container
}

func New(port, apiName, version, allowOrigins, environment string, container *container.Container) *api {
	return &api{
		port:         port,
		apiName:      apiName,
		version:      version,
		allowOrigins: allowOrigins,
		environment:  environment,
		container:    container,
	}
}

func (a *api) NewServer(ctx context.Context) {
	r := gin.Default()
	r.ContextWithFallback = true

	r.Use(otelgin.Middleware("heroes-social-network"))

	r.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Use(cors.New(cors.Config{
		AllowOrigins:  util.ChunkTextByComma(a.allowOrigins),
		AllowMethods:  []string{http.MethodGet},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"Content-Length", "content-type"},
	})).Use(logger.NewLogger(a.environment, a.container.GetVendor(),
		a.container.GetZapLogger()).NewLoggerMiddleware)

	hero.ConfigureHeroRoutes(r, a.container.Domains.HeroUseCase)
	team.ConfigureTeamRoutes(r, a.container.Domains.TeamUseCase)
	game.ConfigureGameRoutes(r, a.container.Domains.GameUseCase)
	console.ConfigureConsoleRoutes(r, a.container.Domains.ConsoleUseCase)
	ability.ConfigureGameRoutes(r, a.container.Domains.AbilityUseCase)

	log.Printf("Server listening in :%s", a.port)

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

func (a *api) initDoc() {
	docs.SwaggerInfo.Title = a.apiName
	docs.SwaggerInfo.Version = a.version
}
