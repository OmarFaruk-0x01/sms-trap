package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type AppValidator struct {
	validator *validator.Validate
}

type Error struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e *Error) Error() string {
	return e.Message
}

func (cv *AppValidator) Validate(i interface{}) error {

	if err := cv.validator.Struct(i); err != nil {
		var parsedError validator.ValidationErrors
		if errors.As(err.(validator.ValidationErrors), &parsedError) {
			out := &Error{
				Message: mapErrorMessageWithTag(parsedError[0].Field(), parsedError[0].Tag()),
				Errors:  map[string]string{},
			}

			for _, e := range parsedError {
				out.Errors[strings.ToLower(e.Field())] = mapErrorMessageWithTag(e.Field(), e.Tag())
			}

			return out
		}
	}

	return nil
}

func mapErrorMessageWithTag(field string, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	default:
		return "Unknown error"
	}
}

func NewAppValidator() *AppValidator {
	return &AppValidator{
		validator: validator.New(),
	}
}
