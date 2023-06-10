package validator

import (
	"log"

	"github.com/LeandroAlcantara-1997/heroes-social-network/model"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CustomValidator struct {
	TagName    string
	CustomFunc func(fl validator.FieldLevel) bool
}

func RegisterValidateFunc(cv []CustomValidator) *validator.Validate {
	vl := validator.New()
	for c := range cv {
		if err := vl.RegisterValidation(cv[c].TagName, cv[c].CustomFunc); err != nil {
			log.Fatal(err)
		}
	}
	return vl
}

func CheckUniverse(fl validator.FieldLevel) bool {
	return model.CheckUniverse(model.Universe(fl.Field().String()))
}

func UUIDValidator(id string) bool {
	if id != "" {
		_, err := uuid.Parse(id)
		return err == nil
	}

	return false
}
