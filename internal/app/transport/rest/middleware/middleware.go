package middleware

import (
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/adapter/log"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/app/transport/rest/response"
	"github.com/LeandroAlcantara-1997/heroes-social-network/internal/exception"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Admin  bool
	Origin Origin
}

type Origin struct {
	Cors *cors.Config
}

func (m *Middleware) Init(ctx *gin.Context) {
	if err := m.verifyMethod(ctx); err != nil {
		code, err := response.RestError(err)
		ctx.AbortWithError(code, err)
		return
	}
	log.AddLoggerInContext(ctx)

}

func (m *Middleware) verifyMethod(ctx *gin.Context) error {
	for methods := range m.Origin.Cors.AllowMethods {
		if m.Origin.Cors.AllowMethods[methods] == ctx.Request.Method {
			return nil
		}
	}
	return exception.ErrInvalidRequest
}
