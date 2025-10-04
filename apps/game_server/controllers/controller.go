package controllers

import (
	"net/http"

	"github.com/sijiaoh/go-godot-template/game_server/ent"
	"github.com/sijiaoh/go-godot-template/game_server/models"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type Controller struct {
	entClient *ent.Client

	currentUser          *models.User
	currentClientSession *models.ClientSession
}

func NewController(entClient *ent.Client) *Controller {
	return &Controller{
		entClient: entClient,
	}
}

func (c *Controller) renderJSON(w http.ResponseWriter, response interface{}) {
	utils.RenderJSON(w, response)
}

func (c *Controller) authenticate(deps *utils.Deps, w http.ResponseWriter, r *http.Request) error {
	token := r.Header.Get("Authorization")[len("Bearer "):]
	cs, err := models.FindClientSessionByToken(deps, token)
	if err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusUnauthorized)
		return err
	}

	if err := cs.LoadUser(deps); err != nil {
		utils.RenderJSONError(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	c.currentUser = cs.User
	c.currentClientSession = cs

	return nil
}

func (c *Controller) isLoggedIn() bool {
	return c.currentUser != nil
}

func (c *Controller) requireLogin(w http.ResponseWriter) bool {
	if !c.isLoggedIn() {
		utils.RenderJSONError(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}
