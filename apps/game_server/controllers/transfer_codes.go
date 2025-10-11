package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/serializers"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

func (c *Controller) ShowTransferCode(w http.ResponseWriter, r *http.Request) {
	deps := utils.NewDeps(c.entClient, r.Context())
	if err := c.authenticate(deps, w, r); err != nil {
		return
	}
	if !c.requireLogin(w) {
		return
	}

	err := c.currentUser.LoadTransferCode(deps)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serializer := serializers.NewTransferCodeSerializer(c.currentUser.TransferCode)
	c.renderJSON(w, serializer)
}

func (c *Controller) RotateTransferCode(w http.ResponseWriter, r *http.Request) {
	deps := utils.NewDeps(c.entClient, r.Context())
	if err := c.authenticate(deps, w, r); err != nil {
		return
	}
	if !c.requireLogin(w) {
		return
	}

	err := c.currentUser.LoadTransferCode(deps)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.currentUser.TransferCode.Rotate(deps)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serializer := serializers.NewTransferCodeSerializer(c.currentUser.TransferCode)
	c.renderJSON(w, serializer)
}
