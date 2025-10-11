package controllers_test

import (
	"net/http"
	"testing"

	"github.com/sijiaoh/go-godot-template/game_server/testutils"
)

func TestShowTransferCode(t *testing.T) {
	testServer := testutils.NewTestServer()
	defer testServer.Close()

	router := testServer.Router
	token := testutils.Signup(t, router, "Foo")

	response := testutils.JSONRequest(t, router, &token, http.MethodGet, "/transfer-code", nil)
	testutils.AssertResponseCode(t, response.Code, http.StatusOK)
}

func TestRotateTransferCode(t *testing.T) {
	testServer := testutils.NewTestServer()
	defer testServer.Close()

	router := testServer.Router
	token := testutils.Signup(t, router, "Foo")

	response := testutils.JSONRequest(t, router, &token, http.MethodPost, "/transfer-code/rotate", nil)
	testutils.AssertResponseCode(t, response.Code, http.StatusOK)
}
