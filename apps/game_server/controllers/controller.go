package controllers

import (
	"encoding/json"
	"fmt"
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

func (c *Controller) renderJson(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		fmt.Println("failed to write response:", err)
	}
}

func (c *Controller) authenticate(deps *utils.Deps, w http.ResponseWriter, r *http.Request) error {
	token := r.Header.Get("Authorization")[len("Bearer "):]
	cs, err := models.FindClientSessionByToken(deps, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	if err := cs.LoadUser(deps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}
	return true
}
