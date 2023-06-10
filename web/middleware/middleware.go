package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Middleware struct {
	RequiredToken bool
	Admin         bool
	Origin        Origin
	Validator     *validator.Validate
}

type Origin struct {
	Cors *cors.Config
}

func (m *Middleware) Init(ctx *gin.Context) {

}
