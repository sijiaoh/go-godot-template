package testutils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func JSONRequest(t *testing.T, router *chi.Mux, token *string, method string, url string, params any) *httptest.ResponseRecorder {
	var body io.Reader = http.NoBody
	if params != nil {
		data, err := json.Marshal(params)
		if err != nil {
			t.Fatal(err)
		}
		body = bytes.NewBuffer(data)
	}
	request := httptest.NewRequest(method, url, body)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json; charset=utf-8")
	if token != nil {
		request.Header.Set("Authorization", "Bearer "+*token)
	}

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	return response
}
