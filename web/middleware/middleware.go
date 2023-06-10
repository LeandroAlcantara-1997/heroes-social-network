package middleware

import (
	"github.com/LeandroAlcantara-1997/heroes-social-network/exception"
	customContext "github.com/LeandroAlcantara-1997/heroes-social-network/pkg/custom_context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Middleware struct {
	Admin     bool
	Origin    Origin
	Validator *validator.Validate
}

type Origin struct {
	Cors *cors.Config
}

func (m *Middleware) Init(ctx *gin.Context) {
	customContext.AddValidator(ctx, m.Validator)
	if err := m.verifyMethod(ctx); err != nil {
		code, err := exception.RestError(err)
		ctx.AbortWithError(code, err)
		return
	}

}

func (m *Middleware) verifyMethod(ctx *gin.Context) error {
	for methods := range m.Origin.Cors.AllowMethods {
		if m.Origin.Cors.AllowMethods[methods] == ctx.Request.Method {
			return nil
		}
	}
	return exception.ErrInvalidRequest
}
