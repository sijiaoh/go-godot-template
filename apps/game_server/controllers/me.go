package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/serializers"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

func (c *Controller) ShowMe(w http.ResponseWriter, r *http.Request) {
	deps := utils.NewDeps(c.entClient, r.Context())
	user, _, err := c.authenticate(deps, w, r)
	if err != nil {
		return
	}
	if !c.requireLogin(w, user) {
		return
	}

	serializer := serializers.NewMeSerializer(user)
	c.renderJSON(w, serializer)
}
