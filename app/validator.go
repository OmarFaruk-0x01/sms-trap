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
				field := strings.ToLower(e.Field())
				tag := e.Tag()

				fieldParts := strings.Split(field, "[")
				baseName := fieldParts[0]
				indexPart := strings.Trim(strings.Join(fieldParts[1:], ""), "]")

				if indexPart != "" {
					field = baseName + "." + indexPart
				}

				out.Errors[field] = mapErrorMessageWithTag(field, tag)
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
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, tag)
	case "numeric":
		return fmt.Sprintf("%s must be a number", field)
	case "oneof":
		return fmt.Sprintf("%s is not valid", field)
	default:
		return "Unknown error"
	}
}

func NewAppValidator() *AppValidator {
	return &AppValidator{
		validator: validator.New(),
	}
}
