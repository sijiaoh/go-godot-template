package validators

import (
	"unicode/utf8"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type StringValidator func(s string) error

type StringLengthError struct {
	ErrorBase

	ActualRunes int
	MinRunes    int
	MaxRunes    int
}

func (e *StringLengthError) Message(localizer i18n.Localizer) (string, error) {
	return localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "validators.string_length",
			Other: "{{.FieldName}}的长度必须在{{.MinRunes}}到{{.MaxRunes}}之间",
		},
		TemplateData: map[string]interface{}{
			"FieldName":   e.FieldName,
			"ActualRunes": e.ActualRunes,
			"MinRunes":    e.MinRunes,
			"MaxRunes":    e.MaxRunes,
		},
	})
}

func ValidateStringLength(minRunes int, maxRunes int) StringValidator {
	return func(s string) error {
		actualRunes := utf8.RuneCountInString(s)
		if actualRunes < minRunes || actualRunes > maxRunes {
			return &StringLengthError{
				ActualRunes: actualRunes,
				MinRunes:    minRunes,
				MaxRunes:    maxRunes,
			}
		}
		return nil
	}
}
