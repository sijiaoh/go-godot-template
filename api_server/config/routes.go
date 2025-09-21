package config

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sijiaoh/go-godot-template/api_server/controllers"
	"github.com/sijiaoh/go-godot-template/api_server/ent"
)

func NewRouter(entClient *ent.Client) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	controller := controllers.NewController(entClient)

	router.Post("/users", controller.CreateUser)

	return router
}
