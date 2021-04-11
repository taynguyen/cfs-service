package views

import "github.com/go-playground/validator/v10"

type ValidationErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(obj interface{}) []*ValidationErrorResponse {
	var errors []*ValidationErrorResponse
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
