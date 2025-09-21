package controllers

import "github.com/sijiaoh/go-godot-template/api_server/ent"

type Controller struct {
	entClient *ent.Client
}

func NewController(entClient *ent.Client) *Controller {
	return &Controller{
		entClient: entClient,
	}
}
