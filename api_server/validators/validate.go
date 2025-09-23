package validators

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func Validate() *validator.Validate {
	if validate == nil {
		validate = validator.New(validator.WithRequiredStructEnabled())
		validate.RegisterValidation("min_runes", minRunes)
		validate.RegisterValidation("max_runes", maxRunes)
	}
	return validate
}
