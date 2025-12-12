package helper

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func ValidateRequest(req any, c *gin.Context) map[string]string {
	v := validator.New()

	if err := v.Struct(req); err != nil {
		if verrs, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)

			for _, e := range verrs {
				switch e.Tag() {
				case "min":
					errors[strings.ToLower(e.Field())] = e.Tag() + " " + e.Param() + " characters"
				case "gt":
					errors[strings.ToLower(e.Field())] = "value must be greater than" + e.Param()
				default:
					errors[strings.ToLower(e.Field())] = e.Tag() + " validation failed"
				}
			}

			return errors
		}
	}

	return nil
}
