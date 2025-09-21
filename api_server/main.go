package main

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/api_server/config"
)

func main() {
	entClient := config.NewEntClient()
	defer entClient.Close()

	router := config.NewRouter(entClient)
	http.ListenAndServe(":3000", router)
}
