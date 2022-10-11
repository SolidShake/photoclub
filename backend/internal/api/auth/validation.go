package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string
	Msg   string
}

func validate(ctx *gin.Context, obj any) []ApiError {
	err := ctx.ShouldBind(obj)

	var out []ApiError
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				out = append(out, ApiError{fe.Field(), msgForTag(fe.Tag())})
			}
		}

		return out
	}

	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return "Value less than minimum"
	case "max":
		return "Value more than maximum"
	}
	return ""
}
