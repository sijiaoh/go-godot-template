package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/serializers"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

func (c *Controller) ShowTransferCode(w http.ResponseWriter, r *http.Request) {
	deps := utils.NewDeps(c.entClient, r.Context())
	user, _, err := c.authenticate(deps, w, r)
	if err != nil {
		return
	}
	if !c.requireLogin(w, user) {
		return
	}

	err = user.LoadTransferCode(deps)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serializer := serializers.NewTransferCodeSerializer(user.TransferCode)
	c.renderJSON(w, serializer)
}

func (c *Controller) RotateTransferCode(w http.ResponseWriter, r *http.Request) {
	deps := utils.NewDeps(c.entClient, r.Context())
	user, _, err := c.authenticate(deps, w, r)
	if err != nil {
		return
	}
	if !c.requireLogin(w, user) {
		return
	}

	err = user.LoadTransferCode(deps)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = user.TransferCode.Rotate(deps)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serializer := serializers.NewTransferCodeSerializer(user.TransferCode)
	c.renderJSON(w, serializer)
}
