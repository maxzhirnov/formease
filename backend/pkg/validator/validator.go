// pkg/validator/validator.go
package validator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("email", validateEmail)
}

func ValidateStruct(s interface{}) error {
	err := validate.Struct(s)
	if err != nil {
		return formatValidationErrors(err)
	}
	return nil
}

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func formatValidationErrors(err error) error {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	var errorMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		switch err.Tag() {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("%s is required", field))
		case "email":
			errorMessages = append(errorMessages, fmt.Sprintf("%s must be a valid email address", field))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", field))
		}
	}

	return fmt.Errorf(strings.Join(errorMessages, "; "))
}
