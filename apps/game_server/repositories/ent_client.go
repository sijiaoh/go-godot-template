package repositories

import (
	"context"

	"github.com/sijiaoh/go-godot-template/game_server/ent"

	_ "github.com/mattn/go-sqlite3"
)

func NewEntClient(memory bool) *ent.Client {
	dataSourceName := "file:tmp/ent.db?cache=shared&_fk=1"
	if memory {
		dataSourceName = dataSourceName + "&mode=memory"
	} else {
		dataSourceName = dataSourceName + "&mode=rwc"
	}
	entClient, err := ent.Open("sqlite3", dataSourceName)
	if err != nil {
		panic(err)
	}
	// TODO: 切换到Migration工具
	if err := entClient.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
	return entClient
}
