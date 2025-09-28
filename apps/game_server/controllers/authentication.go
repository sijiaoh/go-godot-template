package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/models"
	"github.com/sijiaoh/go-godot-template/game_server/serializers"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type SignupParams struct {
	UserName string `json:"userName"`
}

func (c *Controller) Signup(w http.ResponseWriter, r *http.Request) {
	deps := utils.NewDeps(c.entClient, r.Context())

	params, err := utils.ParseJsonBody[SignupParams](w, r.Body)
	if err != nil {
		return
	}

	user := models.NewUser(params.UserName)
	err = user.Save(deps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	cs, err := models.CreateClientSession(deps, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.renderJSON(w, serializers.NewClientSessionSerializer(cs))
}
