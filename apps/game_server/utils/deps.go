package utils

import (
	"context"

	"github.com/sijiaoh/go-godot-template/game_server/ent"
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

func (d *Deps) WithEntTx(fn func(txDeps *Deps) error) error {
	tx, err := d.EntClient.Tx(d.Ctx)
	if err != nil {
		return err
	}

	txDeps := *d
	if err := fn(&txDeps); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return tx.Commit()
}
