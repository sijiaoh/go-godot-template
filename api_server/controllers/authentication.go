package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/api_server/models"
	"github.com/sijiaoh/go-godot-template/api_server/serializers"
	"github.com/sijiaoh/go-godot-template/api_server/utils"
)

type SignupParams struct {
	UserName string `json:"userName"`
}

type SignupResponse struct {
	User *serializers.UserSerializer `json:"user"`
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
	response := SignupResponse{
		User: serializers.NewUserSerializer(user),
	}
	c.renderJson(w, response)
}
