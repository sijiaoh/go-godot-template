package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/serializers"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

func (c *Controller) ShowTransferCode(w http.ResponseWriter, r *http.Request, d *utils.Deps) {
	user, _, err := c.authenticate(d, w, r, true)
	if err != nil {
		return
	}

	err = user.LoadTransferCode(d)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serializer := serializers.NewTransferCodeSerializer(user.TransferCode)
	c.renderJSON(w, serializer)
}

func (c *Controller) RotateTransferCode(w http.ResponseWriter, r *http.Request, d *utils.Deps) {
	user, _, err := c.authenticate(d, w, r, true)
	if err != nil {
		return
	}

	err = user.LoadTransferCode(d)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = user.TransferCode.Rotate(d)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serializer := serializers.NewTransferCodeSerializer(user.TransferCode)
	c.renderJSON(w, serializer)
}
