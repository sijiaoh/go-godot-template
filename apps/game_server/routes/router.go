package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sijiaoh/go-godot-template/game_server/controllers"
	"github.com/sijiaoh/go-godot-template/game_server/ent"
)

func NewRouter(entClient *ent.Client) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	controller := controllers.NewController(entClient)

	router.Post("/signup", controller.ToHandler(controller.Signup))

	router.Get("/me", controller.ToHandler(controller.ShowMe))

	router.Get("/transfer-code", controller.ToHandler(controller.ShowTransferCode))
	router.Post("/transfer-code/rotate", controller.ToHandler(controller.RotateTransferCode))

	return router
}
