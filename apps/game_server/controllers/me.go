package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/serializers"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

func (c *Controller) ShowMe(w http.ResponseWriter, r *http.Request) {
	deps := utils.NewDeps(c.entClient, r.Context())
	if err := c.authenticate(deps, w, r); err != nil {
		return
	}
	if !c.requireLogin(w) {
		return
	}

	serializer := serializers.NewUserSerializer(c.currentUser)
	c.renderJson(w, serializer)
}
