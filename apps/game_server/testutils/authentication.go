package testutils

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/sijiaoh/go-godot-template/game_server/controllers"
	"github.com/sijiaoh/go-godot-template/game_server/serializers"
)

func Signup(t *testing.T, router *chi.Mux, userName string) string {
	params := controllers.SignupParams{
		UserName: userName,
	}
	response := JSONRequest(t, router, nil, http.MethodPost, "/signup", params)
	AssertResponseCode(t, response.Code, http.StatusCreated)

	var res serializers.ClientSessionSerializer
	err := json.Unmarshal(response.Body.Bytes(), &res)
	AssertNoError(t, err)

	return res.Token
}
