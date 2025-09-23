package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/api_server/models"
	"github.com/sijiaoh/go-godot-template/api_server/serializers"
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

	user := models.NewUser(params.Name)
	err = user.Upsert(deps)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	serializer := serializers.NewUserSerializer(user)
	c.renderJson(w, serializer)
}
