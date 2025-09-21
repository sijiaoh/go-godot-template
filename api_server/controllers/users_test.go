package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sijiaoh/go-godot-template/api_server/controllers"
	"github.com/sijiaoh/go-godot-template/api_server/repositories"
	"github.com/sijiaoh/go-godot-template/api_server/routes"
	"github.com/sijiaoh/go-godot-template/api_server/testutils"
)

func TestCreateUser(t *testing.T) {
	entClient := repositories.NewEntClient()
	defer entClient.Close()
	router := routes.NewRouter(entClient)

	params := controllers.CreateUserParams{
		Name: "Foo",
	}
	json, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer(json)
	request := httptest.NewRequest(http.MethodPost, "/users", body)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	testutils.AssertResponseCode(t, response.Code, http.StatusCreated)
	testutils.AssertRecordCount(t, entClient.User.Query(), context.Background(), 1)
}

func TestCreateUser_BadRequest(t *testing.T) {
	entClient := repositories.NewEntClient()
	defer entClient.Close()
	router := routes.NewRouter(entClient)

	body := bytes.NewBuffer([]byte("invalid json"))
	request := httptest.NewRequest(http.MethodPost, "/users", body)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	testutils.AssertResponseCode(t, response.Code, http.StatusBadRequest)
	testutils.AssertRecordCount(t, entClient.User.Query(), context.Background(), 0)
}
