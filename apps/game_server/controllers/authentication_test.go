package controllers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sijiaoh/go-godot-template/game_server/controllers"
	"github.com/sijiaoh/go-godot-template/game_server/serializers"
	"github.com/sijiaoh/go-godot-template/game_server/testutils"
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
	response := testutils.JSONRequest(t, router, nil, http.MethodPost, "/signup", params)
	testutils.AssertResponseCode(t, response.Code, http.StatusCreated)

	testutils.AssertEqual(t, entClient.User.Query().CountX(ctx), 1)
	testutils.AssertEqual(t, entClient.ClientSession.Query().CountX(ctx), 1)

	testutils.AssertEqual(t, entClient.User.Query().FirstX(ctx).Name, params.UserName)

	var res serializers.ClientSessionSerializer
	err := json.Unmarshal(response.Body.Bytes(), &res)
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

func TestLogin(t *testing.T) {
	testServer := testutils.NewTestServer()
	defer testServer.Close()

	entClient := testServer.EntClient
	router := testServer.Router
	ctx := context.Background()

	testutils.Signup(t, router, "Foo")

	params := controllers.LoginParams{
		TransferCode: entClient.TransferCode.Query().FirstX(ctx).Code,
	}
	response := testutils.JSONRequest(t, router, nil, http.MethodPost, "/login", params)
	testutils.AssertResponseCode(t, response.Code, http.StatusCreated)

	testutils.AssertEqual(t, entClient.User.Query().CountX(ctx), 1)
	testutils.AssertEqual(t, entClient.ClientSession.Query().CountX(ctx), 2)

	var res serializers.ClientSessionSerializer
	err := json.Unmarshal(response.Body.Bytes(), &res)
	testutils.AssertNoError(t, err)

	response = testutils.JSONRequest(t, router, &res.Token, http.MethodGet, "/me", nil)
	testutils.AssertResponseCode(t, response.Code, http.StatusOK)
}
