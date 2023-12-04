package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
	"main.go/models"
)

func Validation(data interface{}) (*[]models.Errors, error) {
	var afterErrorCorection []models.Errors
	var result models.Errors
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Tag() {
				case "required":
					err := fmt.Sprintf("%s is required", e.Field())
					result = models.Errors{Error: err}
				case "min":
					err := fmt.Sprintf("%s should be at least %s", e.Field(), e.Param())
					result = models.Errors{Error: err}
				case "max":
					err := fmt.Sprintf("%s should be at most %s", e.Field(), e.Param())
					result = models.Errors{Error: err}
				case "email":
					err := fmt.Sprintf("%s should be email structure %s ", e.Field(), e.Param())
					result = models.Errors{Error: err}
				case "eqfield":
					err := fmt.Sprintf("%s should be equal with %s ", e.Field(), e.Param())
					result = models.Errors{Error: err}
				case "len":
					err := fmt.Sprintf("%s should be have  %s ", e.Field(), e.Param())
					result = models.Errors{Error: err}
				case "alpha":
					err := fmt.Sprintf("%s should be Alphabet ", e.Field())
					result = models.Errors{Error: err}
				case "number":
					err := fmt.Sprintf("%s should be numeric %s ", e.Field(), e.Param())
					result = models.Errors{Error: err}
				case "numeric":
					err := fmt.Sprintf("%s should be  numeric %s ", e.Field(), e.Param())
					result = models.Errors{Error: err}
				case "uppercase":
					err := fmt.Sprintf("%s should be  %s %s ", e.Field(), e.Tag(), e.Param())
					result = models.Errors{Error: err}
				}

				afterErrorCorection = append(afterErrorCorection, result)
			}
		}
		return &afterErrorCorection, errors.New("doesn't fulfill the requirements")
	}
	return &afterErrorCorection, nil
}
