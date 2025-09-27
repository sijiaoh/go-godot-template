package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func StrPtrEq(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func ParseJsonBody[T any](w http.ResponseWriter, body io.ReadCloser) (*T, error) {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	var v T
	err := decoder.Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	return &v, err
}
