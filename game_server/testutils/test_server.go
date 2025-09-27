package testutils

import (
	"github.com/go-chi/chi/v5"
	"github.com/sijiaoh/go-godot-template/api_server/ent"
	"github.com/sijiaoh/go-godot-template/api_server/repositories"
	"github.com/sijiaoh/go-godot-template/api_server/routes"
)

type TestServer struct {
	EntClient *ent.Client
	Router    *chi.Mux
}

func NewTestServer() *TestServer {
	entClient := repositories.NewEntClient()
	router := routes.NewRouter(entClient)

	return &TestServer{
		EntClient: entClient,
		Router:    router,
	}
}

func (ts *TestServer) Close() {
	ts.EntClient.Close()
}
