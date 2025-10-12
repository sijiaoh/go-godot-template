package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type (
	Handler = func(w http.ResponseWriter, r *http.Request)
	Action  = func(w http.ResponseWriter, r *http.Request, d *utils.Deps)
)

func (c *Controller) ToHandler(action Action) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		d := utils.NewDeps(c.entClient, r.Context())
		action(w, r, d)
	}
}
