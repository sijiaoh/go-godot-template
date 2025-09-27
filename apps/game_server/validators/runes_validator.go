package validators

import (
	"strconv"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

func minRunes(fl validator.FieldLevel) bool {
	min, err := strconv.Atoi(fl.Param())
	if err != nil {
		panic(err)
	}
	field := fl.Field().String()
	runesLen := utf8.RuneCountInString(field)
	return runesLen >= min
}

func maxRunes(fl validator.FieldLevel) bool {
	max, err := strconv.Atoi(fl.Param())
	if err != nil {
		panic(err)
	}
	field := fl.Field().String()
	runesLen := utf8.RuneCountInString(field)
	return runesLen <= max
}
