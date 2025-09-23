package utils

import (
	"context"

	"github.com/sijiaoh/go-godot-template/api_server/ent"
)

type Deps struct {
	EntClient *ent.Client
	Ctx       context.Context
}

func NewDeps(entClient *ent.Client, ctx context.Context) *Deps {
	return &Deps{
		EntClient: entClient,
		Ctx:       ctx,
	}
}
