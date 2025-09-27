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

func TestSignup(t *testing.T) {
	testServer := testutils.NewTestServer()
	defer testServer.Close()

	entClient := testServer.EntClient
	router := testServer.Router
	ctx := context.Background()

	params := controllers.SignupParams{
		UserName: "Foo",
	}
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer(data)
	request := httptest.NewRequest(http.MethodPost, "/signup", body)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	testutils.AssertResponseCode(t, response.Code, http.StatusCreated)

	testutils.AssertEqual(t, entClient.User.Query().CountX(ctx), 1)
	testutils.AssertEqual(t, entClient.ClientSession.Query().CountX(ctx), 1)

	testutils.AssertEqual(t, entClient.User.Query().FirstX(ctx).Name, params.UserName)

	var res serializers.ClientSessionSerializer
	err = json.Unmarshal(response.Body.Bytes(), &res)
	testutils.AssertNoError(t, err)
}

func TestSignup_BadRequest(t *testing.T) {
	testServer := testutils.NewTestServer()
	defer testServer.Close()

	entClient := testServer.EntClient
	router := testServer.Router
	ctx := context.Background()

	body := bytes.NewBuffer([]byte("invalid json"))
	request := httptest.NewRequest(http.MethodPost, "/signup", body)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	testutils.AssertResponseCode(t, response.Code, http.StatusBadRequest)

	testutils.AssertEqual(t, entClient.User.Query().CountX(ctx), 0)
}
