package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sijiaoh/go-godot-template/api_server/ent"
)

type Controller struct {
	entClient *ent.Client
}

func NewController(entClient *ent.Client) *Controller {
	return &Controller{
		entClient: entClient,
	}
}

func (c *Controller) renderJson(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		fmt.Println("failed to write response:", err)
	}
}
