package helpers

import "github.com/go-playground/validator/v10"

func ValidationErrorsToSlice(err error) (e []string) {
	for _, valErr := range err.(validator.ValidationErrors) {
		e = append(e, valErr.Error())
	}
	return e
}
