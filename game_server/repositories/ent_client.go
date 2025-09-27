package repositories

import (
	"context"

	"github.com/sijiaoh/go-godot-template/game_server/ent"

	_ "github.com/mattn/go-sqlite3"
)

func NewEntClient() *ent.Client {
	entClient, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	// TODO: 切换到Migration工具
	if err := entClient.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	return entClient
}
