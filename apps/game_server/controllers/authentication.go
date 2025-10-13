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

func (c *Controller) Signup(w http.ResponseWriter, r *http.Request, d *utils.Deps) {
	params, err := utils.ParseJsonBody[SignupParams](w, r.Body)
	if err != nil {
		return
	}

	user, err := models.CreateUser(d, params.UserName)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	cs, err := models.CreateClientSession(d, user)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	c.renderJSON(w, serializers.NewClientSessionSerializer(cs))
}

type LoginParams struct {
	TransferCode string `json:"transferCode"`
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request, d *utils.Deps) {
	params, err := utils.ParseJsonBody[LoginParams](w, r.Body)
	if err != nil {
		return
	}

	cs, err := models.NewAuthenticator().Login(d, params.TransferCode)
	if err != nil {
		// TODO: 区分错误类型
		utils.RenderJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	c.renderJSON(w, serializers.NewClientSessionSerializer(cs))
}
