package repositories

import (
	"context"
	"os"

	"github.com/sijiaoh/go-godot-template/game_server/ent"

	_ "github.com/mattn/go-sqlite3"
)

func NewEntClient() *ent.Client {
	dataSourceName, ok := os.LookupEnv("DB_URL")
	if !ok {
		panic("DB_URL environment variable is not set")
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
