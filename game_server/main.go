package main

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/api_server/repositories"
	"github.com/sijiaoh/go-godot-template/api_server/routes"
)

func main() {
	entClient := repositories.NewEntClient()
	defer entClient.Close()

	router := routes.NewRouter(entClient)
	http.ListenAndServe(":3000", router)
}
