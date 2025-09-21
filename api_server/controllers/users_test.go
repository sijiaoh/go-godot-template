package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sijiaoh/go-godot-template/api_server/controllers"
	"github.com/sijiaoh/go-godot-template/api_server/serializers"
	"github.com/sijiaoh/go-godot-template/api_server/testutils"
)

func TestCreateUser(t *testing.T) {
	testServer := testutils.NewTestServer()
	defer testServer.Close()

	entClient := testServer.EntClient
	router := testServer.Router

	params := controllers.CreateUserParams{
		Name: "Foo",
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer(data)
	request := httptest.NewRequest(http.MethodPost, "/users", body)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	testutils.AssertResponseCode(t, response.Code, http.StatusCreated)
	testutils.AssertRecordCount(t, entClient.User.Query(), context.Background(), 1)

	var res serializers.UserSerializer
	json.Unmarshal(response.Body.Bytes(), &res)
	testutils.AssertStrPtrEqual(t, res.Name, &params.Name)
}

func TestCreateUser_BadRequest(t *testing.T) {
	testServer := testutils.NewTestServer()
	defer testServer.Close()

	entClient := testServer.EntClient
	router := testServer.Router

	body := bytes.NewBuffer([]byte("invalid json"))
	request := httptest.NewRequest(http.MethodPost, "/users", body)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	testutils.AssertResponseCode(t, response.Code, http.StatusBadRequest)
	testutils.AssertRecordCount(t, entClient.User.Query(), context.Background(), 0)
}
