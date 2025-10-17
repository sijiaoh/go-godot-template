package repositories

import (
	"context"
	"os"
	"strings"

	"github.com/sijiaoh/go-godot-template/game_server/ent"

	_ "github.com/mattn/go-sqlite3"
)

func NewEntClient() *ent.Client {
	dbURL, ok := os.LookupEnv("DB_URL")
	if !ok {
		panic("DB_URL environment variable is not set")
	}
	dataSourceName := strings.Replace(dbURL, "sqlite://", "file:", 1)
	entClient, err := ent.Open("sqlite3", dataSourceName)
	if err != nil {
		panic(err)
	}

	isMemoryDB := strings.Contains(dataSourceName, "mode=memory")
	if isMemoryDB {
		if err := entClient.Schema.Create(context.Background()); err != nil {
			panic(err)
		}
	}

	return entClient
}
