package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/api_server/utils"
)

type CreateUserParams struct {
	Name string `json:"name"`
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	deps := utils.NewDeps(c.entClient, r.Context())

	params, err := utils.ParseJsonBody[CreateUserParams](w, r.Body)
	if err != nil {
		return
	}

	_, err = deps.EntClient.User.Create().SetName(params.Name).Save(deps.Ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
