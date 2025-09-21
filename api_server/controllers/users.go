package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/api_server/utils"
)

type CreateUserParams struct {
	Name string `json:"name"`
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var params CreateUserParams
	err := utils.ParseJsonBody(r.Body, &params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = c.entClient.User.Create().SetName(params.Name).Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
