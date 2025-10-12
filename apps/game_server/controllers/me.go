package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/serializers"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

func (c *Controller) ShowMe(w http.ResponseWriter, r *http.Request, d *utils.Deps) {
	user, _, err := c.authenticate(d, w, r, true)
	if err != nil {
		return
	}

	serializer := serializers.NewMeSerializer(user)
	c.renderJSON(w, serializer)
}
