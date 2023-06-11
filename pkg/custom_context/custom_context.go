package customcontext

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	validatorKey string = "validator"
)

func AddValidator(ctx *gin.Context, validator *validator.Validate) {
	ctx.Set(validatorKey, validator)
}

func GetValidator(ctx context.Context) *validator.Validate {
	v := ctx.Value(validatorKey)
	if v != nil {
		return v.(*validator.Validate)
	}
	return nil
}
