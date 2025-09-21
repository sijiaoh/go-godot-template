package utils

import (
	"encoding/json"
	"io"
)

func StrPtrEq(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil {
		return *a == *b
	}
	return *a == *b
}

func ParseJsonBody(body io.ReadCloser, v interface{}) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}
