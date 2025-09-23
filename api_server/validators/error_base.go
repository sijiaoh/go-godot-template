package validators

import "github.com/nicksnyder/go-i18n/v2/i18n"

type ValidationError interface {
	error
	Message(localizer i18n.Localizer) (string, error)
}

type ErrorBase struct {
	FieldName string
}

func (e *ErrorBase) Error() string {
	return "validation error"
}
