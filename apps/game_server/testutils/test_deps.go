package testutils

import (
	"context"

	"github.com/sijiaoh/go-godot-template/game_server/repositories"
	"github.com/sijiaoh/go-godot-template/game_server/utils"
)

type TestDeps struct {
	utils.Deps
}

func NewTestDeps() *TestDeps {
	entClient := repositories.NewEntClient()
	utils.NewDeps(entClient, context.Background())
	return &TestDeps{
		Deps: *utils.NewDeps(entClient, context.Background()),
	}
}

func (td *TestDeps) Close() {
	td.EntClient.Close()
}
