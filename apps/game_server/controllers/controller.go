package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/ent"
	"github.com/sijiaoh/go-godot-template/game_server/models"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type Controller struct {
	entClient *ent.Client
}

func NewController(entClient *ent.Client) *Controller {
	return &Controller{
		entClient: entClient,
	}
}

func (c *Controller) renderJSON(w http.ResponseWriter, response interface{}) {
	utils.RenderJSON(w, response)
}

func (c *Controller) authenticate(deps *utils.Deps, w http.ResponseWriter, r *http.Request, requireLogin bool) (*models.User, *models.ClientSession, error) {
	token := r.Header.Get("Authorization")[len("Bearer "):]
	cs, err := models.FindClientSessionByToken(deps, token, models.ClientSessionPreloadUser)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusUnauthorized)
		return nil, nil, err
	}

	if requireLogin && cs == nil {
		utils.RenderJSONError(w, "Unauthorized", http.StatusUnauthorized)
		return nil, nil, &UnauthorizedError{}
	}

	if cs == nil {
		return nil, nil, nil
	}

	return cs.User, cs, nil
}
